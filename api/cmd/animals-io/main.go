package main

import (
	"fmt"
	dbPkg "github.com/phuc-create/animals-io/db"
	v1 "github.com/phuc-create/animals-io/internal/api/v1"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Animals social app")
	_, err := dbPkg.DatabaseConnect()
	if err != nil {
		log.Fatal(err)
	}

	r := v1.NewRouter()

	log.Fatal(http.ListenAndServe(":8888", r))
}
