package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/phuc-create/animals-io/internal/models"
	"github.com/phuc-create/animals-io/pkg/utils"
	"net/http"
	"net/url"
	"time"
)

type Authentication struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// Users Key is the username
var users = map[string]Authentication{}

func AuthenticationHandlers() chi.Router {
	r := chi.NewRouter()

	r.Post("/register", register)
	r.Post("/login", login)
	r.Post("/logout", logout)
	r.Post("/protected", protected)

	return r
}

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request, u models.UserCredential) error {
	user := users[u.Username]
	// Get the session token from cookie
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		fmt.Println("ERROR READING SESSION TOKEN: ", err)
		return AuthError
	}

	// Get the CSRF Token from the headers
	csrfToken, err := url.QueryUnescape(r.Header.Get("X-Csrf-Token"))
	if err != nil {
		return AuthError
	}
	if csrfToken != user.CSRFToken || csrfToken == "" {
		return AuthError
	}
	fmt.Println("PASS CSRF TOKEN CHECK")
	return nil
}

func protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", err)
		return
	}

	var user models.UserCredential
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		er := http.StatusBadRequest
		http.Error(w, "Something went wrong: "+err.Error(), er)
		return
	}

	if er := Authorize(r, user); er != nil {
		err := http.StatusUnauthorized
		http.Error(w, "Unauthorized", err)
		return
	}

	_, err = fmt.Fprintf(w, "CSRF validation successfully! Welcome %v", user.Username)
	if err != nil {
		fmt.Fprintf(w, "Fprintf: %v\n", err)
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
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

func login(w http.ResponseWriter, r *http.Request) {
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

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}
	var user models.UserCredential
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
