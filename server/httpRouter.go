package server

import (
	"net/http"
)

func initRouter() {
	http.HandleFunc("/api/count", Count)
}
