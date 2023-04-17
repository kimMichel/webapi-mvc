package main

import (
	"github.com/kimMichel/webapi-mvc/database"
	"github.com/kimMichel/webapi-mvc/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
