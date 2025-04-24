package auth

import (
	"encoding/json"
	"fmt"
	"github.com/phuc-create/animals-io/internal/models"
	"github.com/phuc-create/animals-io/pkg/utils"
	"net/http"
	"net/url"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}
	var user models.UserCredential
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

	sessionToken := utils.GenerateToken(32)
	csrfToken, err := url.QueryUnescape(utils.GenerateToken(32)) // Cross Site Request Forgery token
	if err != nil {
		err := http.StatusBadRequest
		http.Error(w, "Something went wrong", err)
		return
	}
	// Set session to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour), // 24 hours
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	// Store token in the database
	u.SessionToken = sessionToken
	// Store csrf token in the database
	u.CSRFToken = csrfToken
	users[user.Username] = u

	_, err = fmt.Fprintln(w, "Login successfully!")
	if err != nil {
		return
	}
}
