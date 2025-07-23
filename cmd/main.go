package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/AnshSinghSonkhia/go-rest-api-postgresql/cmd/api"
	"github.com/AnshSinghSonkhia/go-rest-api-postgresql/config"
	"github.com/AnshSinghSonkhia/go-rest-api-postgresql/db"
)

func main() {
	// initialize db

	dbConn, err := db.NewPostgreSQL(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName,
	))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Printf("Connected to database: %s", config.Envs.DBName)

	if err := initDB(dbConn); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// start api server

	apiServer := api.NewAPIServer(":8080", dbConn)
	if err := apiServer.Run(); err != nil {
		log.Fatal("Failed to start API server:", err)
	}
}

func initDB(db *sql.DB) error {
	// Initialize database connections, migrations, etc.
	err := db.Ping()
	if err != nil {
		return err
	}

	return nil
}
