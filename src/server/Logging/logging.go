package logging

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Received request: Method=%s, URL=%s", req.Method, req.RequestURI)
		next.ServeHTTP(w, req)
		log.Printf("Processed request: Method=%s, URL=%s", req.Method, req.RequestURI)
	})
}
