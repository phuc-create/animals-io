package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", CreateUser)
	return r
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}
