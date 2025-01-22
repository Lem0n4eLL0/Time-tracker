package main

import (
	"fmt"
	"log"
	"net/http"
	logging "timeTrackerApp/src/server/Logging"
	servermux "timeTrackerApp/src/server/ServerMux"
)

var serverPort string = "8080"

func initServer() {
	mux := servermux.CreateServerMux()
	http.Handle("/", mux)
	fmt.Printf("Сервер запущен на порту %s", serverPort)
	if err := http.ListenAndServe(":"+serverPort, logging.Logging(mux)); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func main() {
	initServer()
}
