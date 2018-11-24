package main

import "github.com/wetterj/gin-sqlx-crud/internal/server"

func main() {
	server, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	server.Gin.Run()
}
