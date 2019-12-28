package server

import (
	"net/http"
	db "wosimple/db"
)

func Count(w http.ResponseWriter, r *http.Request) {
	res, _ := db.ExitsKey("Count")
	if res == 0 {
		db.SetRedisString("Count", 1)
	} else {
		db.Incr("Count")
	}
	count, _ := db.GetRedisString("Count")
	var body string = "current count is " + count
	w.Write([]byte(body))
}
