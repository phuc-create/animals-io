package auth

import (
	"encoding/json"
	"fmt"
	"github.com/phuc-create/animals-io/internal/models"
	"github.com/phuc-create/animals-io/pkg/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
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
