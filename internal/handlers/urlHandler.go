package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ccallazans/url-shortener/internal/models"
	"github.com/ccallazans/url-shortener/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func (h *DBRepo) CreateUrlHandler(w http.ResponseWriter, r *http.Request) {

	claims := r.Context().Value("user")//.(*models.User)
	log.Println(claims)
	identification := "asdasd"

	var input struct {
		Url string `json:"url" validate:"required,url"`
	}

	// Parse request json
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

	// Create new url
	newUrl := models.Url{
		Short:     utils.GenerateShort(),
		Url:       input.Url,
		CreatedBy: identification,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	// if short exists
	if h.DB.UrlValueExists(newUrl.Short, "short") {
		utils.ErrorJSON(w, http.StatusConflict, errors.New(`error generating short`))
		return
	}

	// Add into database
	err = h.DB.CreateUrl(newUrl)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, utils.ErrAddToDatabase)
		return
	}

	sendUrl := models.UrlResponse{
		Short:     newUrl.Short,
		Url:       newUrl.Url,
		CreatedBy: newUrl.CreatedBy,
		UpdatedAt: newUrl.UpdatedAt,
		CreatedAt: newUrl.CreatedAt,
	}

	// Send response
	err = utils.WriteJSON(w, http.StatusOK, "response", sendUrl)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *DBRepo) GetUrlByShortHandler(w http.ResponseWriter, r *http.Request) {
	// Get hash from url
	short := chi.URLParam(r, "short")

	// if dont exist
	if !h.DB.UrlValueExists(short, "short") {
		utils.ErrorJSON(w, http.StatusNotFound, errors.New(`short do not exist`))
		return
	}

	// Query data
	newUrl, err := h.DB.GetUrlByShort(short)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, utils.ErrValueDoNotExist)
		return
	}

	// Redirect to url
	http.Redirect(w, r, newUrl.Url, http.StatusPermanentRedirect)
}

func (h *DBRepo) GetAllUrlsHandler(w http.ResponseWriter, r *http.Request) {
	// Query all rows
	allUrls, err := h.DB.GetAllUrls()
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, "response", allUrls)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *DBRepo) UpdateUrlByShortHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request into variable
	var input struct {
		Short string `json:"short" validate:"required,len=5" db:"short"`
		Url   string `json:"url" validate:"required,url" db:"url"`
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

	// if short dont exist
	if !h.DB.UrlValueExists(input.Short, "short") {
		utils.ErrorJSON(w, http.StatusNotFound, utils.ErrValueDoNotExist)
		return
	}

	// Update by short
	err = h.DB.UpdateUrlByShort(input.Short, input.Url)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, "response", input)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
}
