package internal

import (
	"errors"
	"github.com/phuc-create/animals-io/internal/models"
	"net/http"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request, u models.UserCredential) error {
	// Get the session token from cookie
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != "user.SessionToken" { // UPDATE LATER
		return AuthError
	}
	// Get the CSRF Token from the headers
	csrfToken := r.Header.Get("X-Csrf-Token")
	if csrfToken != "user.CSRFToken" || csrfToken == "" { // UPDATE LATER
		return AuthError
	}
	return nil
}
