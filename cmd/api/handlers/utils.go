package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"
	"url-shortener/models"

	"golang.org/x/crypto/bcrypt"
)

// Parse JSON
func WriteJSON(w http.ResponseWriter, status int, wrap string, data interface{}) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func ErrorJSON(w http.ResponseWriter, statusCode int, err error) {

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	WriteJSON(w, statusCode, "error", theError)
}

// Verify user credentials
func verifyParsedFields(user models.User) error {

	err := verifyValidEmail(user.Email)
	if err != nil {
		return err
	}

	err = verifyValidPassword(user.Password)
	if err != nil {
		return err
	}

	return nil
}

func verifyValidEmail(email string) error {
	// Verify valid email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	return nil
}

func verifyValidPassword(password string) error {
	// Verify valid password
	if password == "" || len(password) < 8 {
		return errors.New("invalid password")
	}
	return nil
}

// Password Hash
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}