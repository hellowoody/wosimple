package server

import (
	"log"
	"net/http"
)

func HttpStart() {
	initRouter()
	err := http.ListenAndServe(":9060", nil)
	if err != nil {
		log.Fatal("http start fail: ", err)
	}
}
