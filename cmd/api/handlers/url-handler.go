package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"url-shortener/models"

	"github.com/go-chi/chi/v5"
)

type BaseHandler struct {
	urlRepo models.UrlShortRepository
}

func NewBaseHandler(urlRepo models.UrlShortRepository) *BaseHandler {
	return &BaseHandler{
		urlRepo: urlRepo,
	}
}

func (h *BaseHandler) InsertUrlHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request json
	var newUrl models.UrlRequest
	err := json.NewDecoder(r.Body).Decode(&newUrl)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	if newUrl.Url == "" {
		errorJSON(w, http.StatusBadRequest, errors.New(`missing "url"`))
		return
	}

	if !IsUrl(newUrl.Url) {
		errorJSON(w, http.StatusBadRequest, errors.New(`invalid url`))
		return
	}

	// Create new url
	newUrl.Hash = generateHash()
	newUrl.CreatedAt = time.Now()

	// if hash exists
	if h.urlRepo.HashExists(newUrl.Hash) {
		errorJSON(w, http.StatusBadRequest, errors.New(`error generating hash`))
		return
	}

	// Add into database
	err = h.urlRepo.InsertUrlModel(newUrl)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}

	sendUrl := models.Url{
		Hash:      newUrl.Hash,
		Url:       newUrl.Url,
		CreatedAt: newUrl.CreatedAt,
	}

	// Send response
	err = writeJSON(w, http.StatusOK, "response", sendUrl)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *BaseHandler) GetByHashHandler(w http.ResponseWriter, r *http.Request) {
	// Get hash from url
	hash := chi.URLParam(r, "hash")

	// if dont exist
	if !h.urlRepo.HashExists(hash) {
		errorJSON(w, http.StatusBadRequest, errors.New(`hash do not exist`))
		return
	}

	// Query data
	newUrl, err := h.urlRepo.GetByHash(hash)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Redirect to url
	http.Redirect(w, r, newUrl.Url, http.StatusMovedPermanently)
}

func (h *BaseHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	// Query all rows
	allUrls, err := h.urlRepo.GetAllUrls()
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = writeJSON(w, http.StatusOK, "response", allUrls)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}
}

func (h *BaseHandler) UpdateByHashHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request into variable
	var newUrl models.UrlRequest
	err := json.NewDecoder(r.Body).Decode(&newUrl)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Verify if hash is empty
	if newUrl.Hash == "" {
		errorJSON(w, http.StatusBadRequest, errors.New("missing \"hash\""))
		return
	}

	// Verify if url is empty
	if newUrl.Url == "" {
		errorJSON(w, http.StatusBadRequest, errors.New("missing \"url\""))
		return
	}

	// if dont exist
	if !h.urlRepo.HashExists(newUrl.Hash) {
		errorJSON(w, http.StatusBadRequest, errors.New("hash do not exist"))
		return
	}

	// Update hash
	err = h.urlRepo.UpdateByHash(newUrl)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = writeJSON(w, http.StatusOK, "response", newUrl)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}
}
