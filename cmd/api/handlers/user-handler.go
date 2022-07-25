package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"url-shortener/cmd/api/auth"
	"url-shortener/models"

	"github.com/google/uuid"
)

func (h *BaseHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request json
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Verify if email and password are valid
	err = verifyParsedFields(newUser)
	if err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Verify if email already exists
	if h.userRepo.EmailExists(newUser.Email) {
		ErrorJSON(w, http.StatusConflict, errors.New(`email already registred`))
		return
	}

	// Verify if uuid already exists
	newUser.UUID = uuid.New().String()
	if h.userRepo.UUIDExists(newUser.UUID) {
		ErrorJSON(w, http.StatusInternalServerError, errors.New(`error creating user`))
		return
	}

	// Create Hashed Password
	newUser.Password, err = hashPassword(newUser.Password)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Add created_at
	newUser.CreatedAt = time.Now()

	err = h.userRepo.CreateUser(newUser)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = WriteJSON(w, http.StatusCreated, "response", "user created successful")
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *BaseHandler) AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request json
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Verify if email and password are valid
	err = verifyParsedFields(user)
	if err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	validUser, err := h.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		ErrorJSON(w, http.StatusNotFound, errors.New("email do not exist"))
		return
	}

	if !checkPasswordHash(user.Password, validUser.Password) {
		ErrorJSON(w, http.StatusUnauthorized, errors.New("invalid password"))
		return
	}


	validToken, err := auth.GenerateJwt(user.UUID)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}


	err = WriteJSON(w, http.StatusAccepted, "jwt-token", validToken)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
}
