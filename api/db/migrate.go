package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	dbSourceName := "postgres://localhost:5432/animals_io?sslmode=disable"
	m, err := migrate.New("file://./db/migrations", dbSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
