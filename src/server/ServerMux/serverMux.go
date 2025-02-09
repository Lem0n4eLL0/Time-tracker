package servermux

import (
	"net/http"
	handler "timeTrackerApp/src/server/Handler"

	"github.com/gorilla/mux"
)

var r *mux.Router

func CreateServerMux() *mux.Router {
	r = mux.NewRouter()

	// Маршруты для статических страниц
	r.HandleFunc("/", handler.ServeIndex).Methods("GET")
	r.HandleFunc("/profile", handler.ServeProfile).Methods("GET")
	r.HandleFunc("/projects/{id:[0-9]+}", handler.ServeProject).Methods("GET")

	// API маршруты
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/user/login", handler.LoginHandler).Methods("POST")
	api.HandleFunc("/user/registration", handler.RegisterHandler).Methods("POST")
	api.HandleFunc("/user/projects", handler.GetUserProjects).Methods("GET")
	api.HandleFunc("/user/projects", handler.CreateUserProject).Methods("POST")
	api.HandleFunc("/user/projects/{id:[0-9]+}", handler.UpdateUserProject).Methods("PUT")
	api.HandleFunc("/user/projects/{id:[0-9]+}", handler.DeleteUserProject).Methods("DELETE")
	api.HandleFunc("/user/tasks/categories", handler.GetTaskCategories).Methods("GET")

	api.HandleFunc("/projects/{id:[0-9]+}/tasks", handler.GetProjectTasks).Methods("GET")
	api.HandleFunc("/projects/{id:[0-9]+}/tasks", handler.CreateProjectTask).Methods("POST")
	api.HandleFunc("/projects/{id:[0-9]+}/tasks/{taskId:[0-9]+}", handler.UpdateProjectTask).Methods("PUT")
	api.HandleFunc("/projects/{id:[0-9]+}/tasks/{taskId:[0-9]+}", handler.DeleteProjectTask).Methods("DELETE")
	api.HandleFunc("/projects/report/pdf", handler.GetReportPDF).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../../dist/")))
	return r
}
