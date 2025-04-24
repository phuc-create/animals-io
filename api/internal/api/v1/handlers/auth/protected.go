package auth

import (
	"encoding/json"
	"fmt"
	"github.com/phuc-create/animals-io/internal/models"
	"net/http"
)

func Protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", err)
		return
	}

	var user models.UserCredential
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
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
