package server

import (
	"net/http"
)

func initRouter() {
	http.HandleFunc("/api/count", Count)
	http.HandleFunc("/api/version", GetVersion)
	http.HandleFunc("/api/login", Login)
}
