package main

import (
	"github.com/go-2024/internal"
	"github.com/go-2024/routes"
)

func main() {
	server, err := internal.NewServer()
	if err != nil {
		panic("")
	}

	routes.PostRoutes(server)

	server.S.Run(":8080")

}
