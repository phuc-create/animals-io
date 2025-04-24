package db

import (
	"errors"
	"fmt"
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

	version, dirty, err := m.Version()
	if err != nil {
		//	Special case: no any migration has been run yet
		if err == migrate.ErrNilVersion {
			log.Println("No migration applied yet.")
		} else {
			log.Fatalf("Failed to get migration version: %v", err)
		}
	}
	fmt.Printf("Current version: %d, dirty: %v\n", version, dirty)
	if dirty {
		forceVersion := int(version - 1)

		// Force migration version
		err := m.Force(forceVersion)
		if err != nil {
			log.Fatalf("Failed to force migration version: %v", err)
		}
		fmt.Printf("Forced database to clean version %d\n", forceVersion)
	} else {
		fmt.Println("Database is clean")
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Migration applied successfully!")
}
