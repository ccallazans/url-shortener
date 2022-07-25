package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"url-shortener/models"

	"github.com/google/uuid"
)

func (h *BaseHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request json
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Verify if email and password are valid
	err = verifyParsedFields(newUser)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Verify if email already exists
	if h.userRepo.EmailExists(newUser.Email) {
		errorJSON(w, http.StatusBadRequest, errors.New(`email already registred`))
		return
	}

	// Verify if uuid already exists
	newUser.UUID = uuid.New().String()
	if h.userRepo.UUIDExists(newUser.UUID) {
		errorJSON(w, http.StatusBadRequest, errors.New(`error creating user`))
		return
	}

	// Create Hashed Password
	newUser.Password, err = hashPassword(newUser.Password)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Add created_at
	newUser.CreatedAt = time.Now()

	err = h.userRepo.CreateUser(newUser)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	err = writeJSON(w, http.StatusAccepted, "response", "user created successful")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}
}

func (h *BaseHandler) AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request json
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Verify if email and password are valid
	err = verifyParsedFields(user)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	validUser, err := h.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, errors.New("email do not exist"))
		return
	}

	if !checkPasswordHash(user.Password, validUser.Password) {
		errorJSON(w, http.StatusBadRequest, errors.New("invalid password"))
		return
	}

	err = writeJSON(w, http.StatusAccepted, "response", "authenticated")
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}
}
