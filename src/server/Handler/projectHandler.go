package handler

import (
	"encoding/json"
	"io"
	"strconv"

	"net/http"

	database "timeTrackerApp/src/server/Database"
	eh "timeTrackerApp/src/server/ErrorHandler"
	s "timeTrackerApp/src/server/Structures"
	token "timeTrackerApp/src/server/Token"

	"github.com/gorilla/mux"
)

// Получение всех проектов без Tasks
func GetUserProjects(w http.ResponseWriter, r *http.Request) {
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

	projects, err := database.GetProjects(payload.UserID)
	if err != nil {
		eh.ResponseWithError(w, http.StatusNotFound, err.Error())
		return
	}

	message, err := json.Marshal(struct{ Projects []s.Project }{Projects: projects})
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// Добавление проекта
func CreateUserProject(w http.ResponseWriter, r *http.Request) {
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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	var project s.Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при парсинге JSON")
		return
	}

	err = database.CreateProject(&s.Project{UserID: payload.UserID, ProjectName: project.ProjectName, Description: project.Description})
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	p, err := database.GetLastCreateProject(payload.UserID)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	message, err := json.Marshal(p)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// Удаление проекта
func DeleteUserProject(w http.ResponseWriter, r *http.Request) {
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

	err = database.DeleteProject(payload.UserID, projectID)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Изменение проекта
func UpdateUserProject(w http.ResponseWriter, r *http.Request) {
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

	var project s.Project
	err = json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, "Ошибка при парсинге JSON")
		return
	}

	err = database.UpdateProject(payload.UserID, &project)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// func ProjectHandlerGet(w http.ResponseWriter, r *http.Request) {
//   c, err := r.Cookie("session_token")
// 	if err != nil {
// 		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	payload, err := token.GetTokenMaker().VerifyToken(c.Value)
// 	if err != nil {
// 		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// }
