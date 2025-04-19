package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/phuc-create/animals-io/pkg/utils"
	"net/http"
)

type Authentication struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}
type UserCredential struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// Key is the username
var users = map[string]Authentication{}

func AuthenticationHandlers() chi.Router {
	r := chi.NewRouter()

	r.Post("/register", register)
	r.Post("/login", login)
	r.Post("/logout", logout)
	r.Post("/protected", protected)

	return r
}

func protected(w http.ResponseWriter, r *http.Request) {

}

func logout(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}
	var user UserCredential
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err := http.StatusBadRequest
		http.Error(w, "Something went wrong", err)
		return
	}
	u, ok := users[user.Username]
	if !ok || !utils.CheckPasswordHash(user.Password, u.HashedPassword) {
		err := http.StatusUnauthorized
		http.Error(w, "Invalid username and password", err)
		return
	}

	_, err = fmt.Fprintln(w, "Login successfully!")
	if err != nil {
		return
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}
	var user UserCredential
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(user.Username) < 8 || len(user.Password) < 8 {
		http.Error(w, "Invalid username/password", http.StatusNotAcceptable)
		return
	}
	if _, existed := users[user.Username]; existed {
		http.Error(w, "User already existed", http.StatusConflict)
		return
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusConflict)
		return
	}

	users[user.Username] = Authentication{HashedPassword: hashedPassword}
	fmt.Println("Registered user account successfully!")
	fmt.Println(users)
}
