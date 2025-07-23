package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AnshSinghSonkhia/go-rest-api-postgresql/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Build database connection string using environment variables
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBName,
	)

	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		dbURL,
	)
	if err != nil {
		log.Fatal("error initialize migration:", err)
	}

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Println("migrate up error:", err)
		} else {
			log.Println("migrate up success")
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Println("migrate down error:", err)
		} else {
			log.Println("migrate down success")
		}
	}
}
