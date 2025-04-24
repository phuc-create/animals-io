package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DatabaseConnect() (*sql.DB, error) {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading env file...")
	}

	dbHost := os.Getenv("POSTGRES_DB_HOST")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	dbUsername := os.Getenv("POSTGRES_DB_USERNAME")
	dbPassword := os.Getenv("POSTGRES_DB_PASSWORD")
	dbPort := os.Getenv("POSTGRES_DB_PORT")

	conStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	fmt.Printf("Database connected susscessfully")
	return db, nil
}
