package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ccallazans/url-shortener/cmd/api/auth"
	"github.com/ccallazans/url-shortener/cmd/api/utils"
	"github.com/ccallazans/url-shortener/models"
	"github.com/google/uuid"
)

func (h *BaseHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request json
	var input struct {
		Email    string `json:"email" validate:"required,email" db:"email"`
		Password string `json:"password" validate:"required,min=8,max=30" db:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	err = validate.Struct(input)
	if err != nil {
		utils.ErrorJSON(w, http.StatusConflict, err)
		return
	}

	// Verify if email already exists
	log.Println(h.userRepo.ValueExists(input.Email, "email"))
	if h.userRepo.ValueExists(input.Email, "email") {
		utils.ErrorJSON(w, http.StatusConflict, errors.New(`email already registred`))
		return
	}

	newUser := models.User{
		UUID:      uuid.New().String(),
		Email:     input.Email,
		Password:  input.Password,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	// Verify if uuid already exists
	if h.userRepo.ValueExists(newUser.UUID, "uuid") {
		utils.ErrorJSON(w, http.StatusInternalServerError, errors.New(`error creating user`))
		return
	}

	// Create Hashed Password
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	newUser.Password = *hashedPassword

	err = h.userRepo.CreateUser(newUser)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, "response", "user created successful")
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *BaseHandler) AuthUserHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Email    string `json:"email" validate:"required,email" db:"email"`
		Password string `json:"password" validate:"required" db:"password"`
	}

	// Parse request json
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Validate input
	err = validate.Struct(input)
	if err != nil {
		utils.ErrorJSON(w, http.StatusConflict, err)
		return
	}

	// Verify and get user
	validUser, err := h.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		utils.ErrorJSON(w, http.StatusNotFound, errors.New("email do not exist"))
		return
	}

	// Check user password
	err = utils.CheckPasswordHash(input.Password, validUser.Password)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnauthorized, errors.New("invalid password"))
		return
	}

	// Generate user JWT
	validToken, err := auth.GenerateJWT(validUser)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, "token", validToken)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
}
