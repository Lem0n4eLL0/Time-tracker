package handler

import (
	"encoding/json"
	"io"

	"net/http"
	"time"
	db "timeTrackerApp/src/server/Database"
	eh "timeTrackerApp/src/server/ErrorHandler"

	token "timeTrackerApp/src/server/Token"
	"timeTrackerApp/src/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	var loginRequest LoginRequest
	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при парсинге JSON")
		return
	}

	user, err := db.GetUserByName(loginRequest.Username)
	if err != nil {
		eh.ResponseWithError(w, http.StatusInternalServerError, "Ошибка при выполнении запроса")
	}

	if user.Password != utils.Sha512Hashing(loginRequest.Password) {
		eh.ResponseWithError(w, http.StatusUnauthorized, "Неверный пароль")
		return
	}

	token, err := token.GetTokenMaker().CreateToken(user, time.Minute*5)

	if err != nil {
		eh.ResponseWithError(w, http.StatusUnauthorized, err.Error())
	}

	jsonResp, err := json.Marshal(ResponseMessage{Message: "Аутентификация прошла успешно"})
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при генерации ответа")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: token,
		Path:  "/",
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
