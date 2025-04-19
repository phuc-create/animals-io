package main

import (
	"fmt"
	v1 "github.com/phuc-create/animals-io/internal/api/v1"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Animals social app")
	r := v1.NewRouter()

	log.Fatal(http.ListenAndServe(":8888", r))
}
