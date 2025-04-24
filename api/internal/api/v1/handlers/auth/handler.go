package auth

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/phuc-create/animals-io/internal/models"
	"net/http"
	"net/url"
)

type Authentication struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// Users Key is the username
var users = map[string]Authentication{}

func Handlers() chi.Router {
	r := chi.NewRouter()

	r.Post("/register", Register)
	r.Post("/login", Login)
	r.Post("/logout", Logout)
	r.Post("/protected", Protected)

	return r
}

var AuthorizeError = errors.New("unauthorized")

func Authorize(r *http.Request, u models.UserCredential) error {
	user := users[u.Username]
	// Get the session token from cookie
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		fmt.Println("ERROR READING SESSION TOKEN: ", err)
		return AuthorizeError
	}

	// Get the CSRF Token from the headers
	csrfToken, err := url.QueryUnescape(r.Header.Get("X-Csrf-Token"))
	if err != nil {
		return AuthorizeError
	}
	if csrfToken != user.CSRFToken || csrfToken == "" {
		return AuthorizeError
	}
	fmt.Println("PASS CSRF TOKEN CHECK")
	return nil
}
