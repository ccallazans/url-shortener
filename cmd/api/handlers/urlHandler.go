package handlers

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"

	// "github.com/ccallazans/url-shortener/models"

	"github.com/ccallazans/url-shortener/cmd/api/utils"
	"github.com/ccallazans/url-shortener/models"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

func (h *BaseHandler) CreateUrlHandler(w http.ResponseWriter, r *http.Request) {

	claims := r.Context().Value("user").(jwt.MapClaims)
	identification := claims["email"].(string)

	var input struct {
		Url string `json:"url" validate:"required,url"`
	}

	// Parse request json
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

	// Create new url
	newUrl := models.Url{
		Short:     generateShort(),
		Url:       input.Url,
		CreatedBy: identification,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	// if short exists
	if h.urlRepo.ValueExists(newUrl.Short, "short") {
		utils.ErrorJSON(w, http.StatusConflict, errors.New(`error generating short`))
		return
	}

	// Add into database
	err = h.urlRepo.CreateUrl(newUrl)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
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

func (h *BaseHandler) GetUrlByShortHandler(w http.ResponseWriter, r *http.Request) {
	// Get hash from url
	short := chi.URLParam(r, "short")

	// if dont exist
	if !h.urlRepo.ValueExists(short, "short") {
		utils.ErrorJSON(w, http.StatusNotFound, errors.New(`short do not exist`))
		return
	}

	// Query data
	newUrl, err := h.urlRepo.GetUrlByShort(short)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Redirect to url
	http.Redirect(w, r, newUrl.Url, http.StatusPermanentRedirect)
}

func (h *BaseHandler) GetAllUrlsHandler(w http.ResponseWriter, r *http.Request) {
	// Query all rows
	allUrls, err := h.urlRepo.GetAllUrls()
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

func (h *BaseHandler) UpdateUrlByShortHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request into variable
	var input struct {
		Short string `json:"short" validate:"required,len=5" db:"short"`
		Url   string `json:"url" validate:"required,url" db:"url"`
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

	// if short dont exist
	if !h.urlRepo.ValueExists(input.Short, "short") {
		utils.ErrorJSON(w, http.StatusNotFound, errors.New("short do not exist"))
		return
	}

	// Update by short
	err = h.urlRepo.UpdateUrlByShort(input.Short, input.Url)
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

func generateShort() string {
	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	hash := make([]byte, 5)
	for i := range hash {
		hash[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(hash)
}
