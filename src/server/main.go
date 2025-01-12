package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"time"
	eh "timeTrackerApp/src/server/ErrorHandler"
	register "timeTrackerApp/src/server/Handler/RgisterHandler"
	s "timeTrackerApp/src/server/Structures"
	"timeTrackerApp/src/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn
var serverPort string = "8080"
var jwtSecretKey = []byte("key") // ключ переделать

type User struct {
	UserID   int    `db:"user_id" json:"id"`
	Name     string `db:"username" json:"username"`
	Email    any    `db:"email" json:"email"`
	Password string `db:"password_hash" json:"password"`
	Created  any    `db:"created_at" json:"created_at"`
	Updated  any    `db:"updated_at" json:"updated_at"`
	Role     int    `db:"role_id" json:"role_id"`
}

type ResponseMessage struct {
	Message string `json:message"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenDate struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password_hash"`
	Role     string `json:"role_id"`
}

// type APIError struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }

func Logging(next http.Handler) http.Handler { // Логирование
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Received request: Method=%s, URL=%s", req.Method, req.RequestURI)
		next.ServeHTTP(w, req)
		log.Printf("Processed request: Method=%s, URL=%s", req.Method, req.RequestURI)
	})
}

// func responseWithError(w http.ResponseWriter, code int, message string) { // Обработка ошибок
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	json.NewEncoder(w).Encode(APIError{Code: code, Message: message})
// 	fmt.Println(message)
// }

// errHand.ResponseWithError(w, code, message)
func checkConn(w http.ResponseWriter) { // Ошибка если база данных не подключена
	if conn == nil {
		eh.ResponseWithError(w, http.StatusInternalServerError, "Подключение к базе данных не инициализировано")
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) { // Получить список userov
	if r.Method != http.MethodGet {
		eh.ResponseWithError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}
	checkConn(w)

	rows, err := conn.Query(context.Background(),
		"SELECT user_id, username, email, password_hash, created_at, updated_at, roles.name FROM users INNER JOIN roles ON roles.role_id=users.role_id ORDER BY user_id ASC")
	if err != nil {
		eh.ResponseWithError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.User])
	if err != nil {
		eh.ResponseWithError(w, http.StatusMethodNotAllowed, "Метод не ")
		return
	}

	for i := range users {
		users[i].ResetNonAdminFild()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

type MyClaims struct { //fbg
	jwt.RegisteredClaims
	Username string   `json:"sub"`
	Exp      int64    `json:"exp"`
	Admin    bool     `json:"admin"`
	Date     []string `json:"date"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) { // Аутентификация пользователя
	if r.Method != http.MethodPost {
		eh.ResponseWithError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}
	checkConn(w)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при чтении данных")
		return
	}
	defer r.Body.Close()

	var loginRequest LoginRequest
	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при парсинге JSON")
		return
	}

	row := conn.QueryRow(context.Background(), "SELECT password_hash FROM users WHERE username=$1", loginRequest.Username)

	var password string
	err = row.Scan(&password)
	if err != nil {
		eh.ResponseWithError(w, http.StatusInternalServerError, "Ошибка при выполнении запроса")
	}

	if password != utils.Sha512Hashing(loginRequest.Password) {
		eh.ResponseWithError(w, http.StatusUnauthorized, "Неверный пароль")
		return
	}

	// payload := jwt.MapClaims{
	// 	"sub": loginRequest.Username,
	// 	"exp": time.Now().Add(time.Hour).Unix(),
	// }
	//.

	// Создание токена с пользовательскими утверждениями
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		Username:         loginRequest.Username,
		Admin:            false,
		Exp:              time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Секретный ключ для подписи

	// Генерация токена в строковом формате
	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		log.Fatalf("Произошла ошибка: %v", err)
	}
	//.

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", t)
	w.WriteHeader(http.StatusOK)

	jsonResp, err := json.Marshal(ResponseMessage{"asd"})
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при генерации ответа")
		return
	}
	w.Write(jsonResp)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) { //Регестрация пользователя
	if r.Method != http.MethodPost {
		eh.ResponseWithError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}
	checkConn(w)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при чтении данных")
		return
	}
	defer r.Body.Close()

	var registerRequest RegisterRequest
	err = json.Unmarshal(body, &registerRequest)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при парсинге JSON")
		return
	}

	var userId int
	row := conn.QueryRow(context.Background(), "SELECT user_id FROM users WHERE username=$1", registerRequest.Username)
	err = row.Scan(&userId)
	if err == nil {
		eh.ResponseWithError(w, http.StatusInternalServerError, "Такой пользователь уже существует")
		return
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3)", registerRequest.Username, utils.Sha512Hashing(registerRequest.Password), registerRequest.Email)
	if err != nil {
		fmt.Println(err)
		eh.ResponseWithError(w, http.StatusInternalServerError, "Ошибка выполнения запроса")
		return
	}
	fmt.Println("Пользователь успешно создан")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователь успешно создан"))
}

func verifyToken(token string) (bool, error) {
	// Нам нужно определить функцию, которую пакет jwt будет использовать для разбора tokenString
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		// Проверяем, что используется ожидаемый метод подписи
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Неожиданный метод подписи: %v", t.Header["alg"])
		}
		// Возвращаем секретный ключ для jwt токена, в формате []byte, совпадающий с ключом, использованным для подписи ранее
		return jwtSecretKey, nil
	}
	// Разбор токена
	claims := &MyClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, keyFunc)
	if err != nil {
		return false, err
	}
	if !parsedToken.Valid {
		return false, err
	}
	return true, nil
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		eh.ResponseWithError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}
	checkConn(w)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при чтении данных")
		return
	}
	defer r.Body.Close()

	ok, err := verifyToken(r.Header.Get("Authorization"))
	fmt.Println(ok)
	if !ok {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при чтении данных "+err.Error())
		return
	}
	fmt.Println(body)
	w.Write([]byte("Тест успешен"))
}

func main() {
	var err error
	conn, err = pgx.Connect(context.Background(), "postgres://postgres:zerr0@[::1]:8000/USERS?sslmode=disable")
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer conn.Close(context.Background())

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("../../dist/")))
	mux.HandleFunc("/users/", getUsers)
	mux.HandleFunc("/login/", LoginHandler)
	//mux.HandleFunc("/registration/", RegisterHandler)
	mux.HandleFunc("/registration/", register.RegisterHandler)
	mux.HandleFunc("/profile/", ProfileHandler)

	log.Printf("Сервер запущен на порту %s", serverPort)
	if err := http.ListenAndServe(":"+serverPort, Logging(mux)); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
