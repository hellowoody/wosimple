package server

import (
	"log"
	"net/http"
)

func HttpStart() {
	initRouter()
	// err := http.ListenAndServeTLS(":9060", "/etc/letsencrypt/archive/woodyhello.com/fullchain1.pem", "/etc/letsencrypt/archive/woodyhello.com/privkey1.pem", nil)
	// err := http.ListenAndServeTLS(":9060", "/etc/letsencrypt/live/woodyhello.com/fullchain.pem", "/etc/letsencrypt/live/woodyhello.com/privkey.pem", nil)
	err := http.ListenAndServe(":9060", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
