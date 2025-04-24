package auth

import (
	"encoding/json"
	"fmt"
	"github.com/phuc-create/animals-io/internal/models"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	var user models.UserCredential
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err := http.StatusBadRequest
		http.Error(w, "Something went wrong", err)
		return
	}
	u := users[user.Username]
	if err := Authorize(r, user); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//	Clear the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		MaxAge:   -1,
		HttpOnly: true,
	})

	// Clear the CSRF Token
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		MaxAge:   -1,
		HttpOnly: false,
	})

	// Clear the token from database
	u.SessionToken = ""
	u.CSRFToken = ""
	users[user.Username] = u

	fmt.Fprintf(w, "Logged out successfully!")
}
