package main

import (
	"fmt"
	server "wosimple/server"
)

func main() {
	fmt.Println("start wosimple")
	// db.InitDb()
	// defer db.SqlDB.Close()
	server.HttpStart()
}
