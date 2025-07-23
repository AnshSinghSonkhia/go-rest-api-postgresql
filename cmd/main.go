package main

import (
	"log"

	"github.com/AnshSinghSonkhia/go-rest-api-postgresql/cmd/api"
)

func main() {
	// initialize db

	// start api server

	apiServer := api.NewAPIServer(":8080")
	if err := apiServer.Run(); err != nil {
		log.Fatal("Failed to start API server:", err)
	}
}

// 13:39
