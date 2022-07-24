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

func (h *BaseHandler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/404", http.StatusMovedPermanently) 
	errorJSON(w, http.StatusNotFound, errors.New("page not found"))
}

func (h *BaseHandler) InsertUrlHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request json
	var newUrl models.UrlShortRequest
	err := json.NewDecoder(r.Body).Decode(&newUrl)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Create new url
	newUrl.Hash = generateHash()
	newUrl.CreatedAt = time.Now()

	// Verify if url already exists
	err = h.urlRepo.VerifyExists("original_url", newUrl.OriginalUrl)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Verify if hash already exists
	err = h.urlRepo.VerifyExists("hash", newUrl.Hash)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Add into database
	err = h.urlRepo.AddUrl(newUrl)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Send response
	err = writeJSON(w, http.StatusOK, "response", newUrl)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *BaseHandler) GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	// Get hash from url
	hash := chi.URLParam(r, "hash")

	// Verify if hash exists
	err := h.urlRepo.VerifyExists("hash", hash)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Query data
	url, err := h.urlRepo.GetUrl(hash)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Redirect to url
	http.Redirect(w, r, *url, http.StatusMovedPermanently)
	// http.Redirect(w, r, fmt.Sprintf("http://%s", *url), 301)
}

func (h *BaseHandler) GetAllUrlsHandler(w http.ResponseWriter, r *http.Request) {
	// Query all rows
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

func (h *BaseHandler) UpdateHashHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request into variable
	var newUrl models.UrlShortRequest
	err := json.NewDecoder(r.Body).Decode(&newUrl)
	if err != nil {
		errorJSON(w, http.StatusBadRequest, err)
		return
	}

	// Verify if request is empty
	if newUrl.Hash == "" || newUrl.OriginalUrl == "" {
		errorJSON(w, http.StatusBadRequest, errors.New("missing 'hash' or 'original_url' parameters"))
		return
	}

	//Verify if url exists
	err = h.urlRepo.VerifyExists("original_url", newUrl.OriginalUrl)
	if err != nil {
		errorJSON(w, http.StatusNotFound, err)
		return
	}

	//Verify if hash exists
	err = h.urlRepo.VerifyExists("hash", newUrl.Hash)
	if err != nil {
		errorJSON(w, http.StatusNotFound, err)
		return
	}

	// Update hash
	err = h.urlRepo.UpdateHash(newUrl)
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