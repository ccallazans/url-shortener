package handlers

import (
	"encoding/json"
	"fmt"
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
	var newUrl models.UrlShort
	err := json.NewDecoder(r.Body).Decode(&newUrl)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Create new url
	newUrl.Hash = generateHash()
	newUrl.CreatedAt = time.Now()

	// Add into database
	err = h.urlRepo.AddUrl(newUrl)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Send response
	writeJSON(w, http.StatusOK, "response", newUrl)
}

func (h *BaseHandler) GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")

	original_url, err := h.urlRepo.GetUrl(hash)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("http://%s", *original_url), 301)
}

func (h *BaseHandler) GetAllUrlsHandler(w http.ResponseWriter, r *http.Request) {
	urls, err := h.urlRepo.GetAllUrls()
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = writeJSON(w, http.StatusOK, "urls", urls)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}
}
