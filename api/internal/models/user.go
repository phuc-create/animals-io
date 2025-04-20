package models

type UserCredential struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
