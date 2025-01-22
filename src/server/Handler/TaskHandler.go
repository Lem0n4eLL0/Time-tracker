package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	database "timeTrackerApp/src/server/Database"
	eh "timeTrackerApp/src/server/ErrorHandler"
	s "timeTrackerApp/src/server/Structures"
	token "timeTrackerApp/src/server/Token"

	"github.com/gorilla/mux"
)

func GetProjectTasks(w http.ResponseWriter, r *http.Request) {
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

	tasks, err := database.GetTasks(payload.UserID, projectID)
	if err != nil {
		eh.ResponseWithError(w, http.StatusNotFound, err.Error())
		return
	}

	message, err := json.Marshal(struct{ Tasks []s.Task }{Tasks: tasks})
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func CreateProjectTask(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = token.GetTokenMaker().VerifyToken(c.Value)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	var task s.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = database.CreateTask(&task)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	p, err := database.GetLastCreateTask(task.ProjectID)
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

func DeleteProjectTask(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = token.GetTokenMaker().VerifyToken(c.Value)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	taskID, _ := strconv.Atoi(mux.Vars(r)["taskId"])

	err = database.DeleteTask(taskID)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateProjectTask(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = token.GetTokenMaker().VerifyToken(c.Value)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	var task s.Task

	err = json.Unmarshal(body, &task)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = database.UpdateTask(&task)
	if err != nil {
		eh.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetTaskCategories(w http.ResponseWriter, r *http.Request) {
	p, err := database.GetTaskCategories()
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
