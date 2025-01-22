package handler

import (
	"html/template"
	"net/http"
	"strconv"
	database "timeTrackerApp/src/server/Database"
	eh "timeTrackerApp/src/server/ErrorHandler"
	token "timeTrackerApp/src/server/Token"

	"github.com/gorilla/mux"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("G:\\kursach_PP\\dist\\index.html")
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	p := token.Payload{}
	c, err := r.Cookie("session_token")
	if err != nil {
		tmpl.Execute(w, p)
		return
	}

	payload, err := token.GetTokenMaker().VerifyToken(c.Value)
	if err != nil {
		tmpl.Execute(w, p)
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, payload)
}

func ServeProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		eh.ResponseWithError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	c, err := r.Cookie("session_token")
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	payload, err := token.GetTokenMaker().VerifyToken(c.Value)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	tmpl, err := template.ParseFiles("G:\\kursach_PP\\dist\\profile.html")
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, payload)
}

func ServeProject(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	payload, err := token.GetTokenMaker().VerifyToken(c.Value)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	projectID, _ := strconv.Atoi(mux.Vars(r)["id"])

	tasks, err := database.GetProject(payload.UserID, projectID)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if tasks == nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Данный проект не найден")
		return
	}

	tmpl, err := template.ParseFiles("G:\\kursach_PP\\dist\\project.html")
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, payload)
}
