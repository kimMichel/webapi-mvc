package main

import "github.com/kimMichel/webapi-mvc/server"

func main() {
	server := server.NewServer()

	server.Run()
}
