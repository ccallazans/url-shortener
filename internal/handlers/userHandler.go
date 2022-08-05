package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ccallazans/url-shortener/internal/auth"
	"github.com/ccallazans/url-shortener/internal/models"
	"github.com/ccallazans/url-shortener/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (h *DBRepo) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request json
	var input struct {
		Email    string `json:"email" validate:"required,email" db:"email"`
		Password string `json:"password" validate:"required,min=8,max=30" db:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, utils.ErrWrongRequestBody)
		return
	}

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		utils.ErrorJSON(w, http.StatusConflict, utils.ErrInvalidRequirements)
		return
	}

	// Verify if email already exists
	if h.DB.UserValueExists(input.Email, "email") {
		utils.ErrorJSON(w, http.StatusConflict, utils.ErrEmailAlreadyExists)
		return
	}

	newUser := models.User{
		UUID:      uuid.NewString(),
		Email:     input.Email,
		Password:  input.Password,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	// Verify if uuid already exists
	if h.DB.UserValueExists(newUser.UUID, "uuid") {
		utils.ErrorJSON(w, http.StatusInternalServerError, utils.ErrAddToDatabase)
		return
	}

	// Create Hashed Password
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	newUser.Password = *hashedPassword

	err = h.DB.CreateUser(newUser)
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

func (h *DBRepo) AuthUserHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Email    string `json:"email" validate:"required,email" db:"email"`
		Password string `json:"password" validate:"required" db:"password"`
	}

	// Parse request json
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, utils.ErrWrongRequestBody)
		return
	}

	// Validate input
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		utils.ErrorJSON(w, http.StatusConflict, utils.ErrInvalidRequirements)
		return
	}

	// Verify and get user
	validUser, err := h.DB.GetUserByEmail(input.Email)
	if err != nil {
		utils.ErrorJSON(w, http.StatusNotFound, utils.ErrEmailDoNotExist)
		return
	}

	// Check user password
	err = utils.CheckPasswordHash(input.Password, validUser.Password)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnauthorized, utils.ErrInvalidPassword)
		return
	}

	// Generate user JWT
	validToken, err := auth.GenerateJWT(validUser)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, utils.ErrGenerateJWT)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, "token", validToken)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
}
