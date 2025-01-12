package register

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	db "timeTrackerApp/src/server/Database"
	eh "timeTrackerApp/src/server/ErrorHandler"
	s "timeTrackerApp/src/server/Structures"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) { //Регестрация пользователя
	if r.Method != http.MethodPost {
		eh.ResponseWithError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}
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
	_, err = db.GetUserByName(registerRequest.Username)
	if err == nil {
		eh.ResponseWithError(w, http.StatusInternalServerError, "Такой пользователь уже существует")
		return
	}
	err = db.CreateUser(&s.User{Email: registerRequest.Email, Name: registerRequest.Username, Password: registerRequest.Password})
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
