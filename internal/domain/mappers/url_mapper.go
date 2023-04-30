package mappers

import "github.com/ccallazans/url-shortener/internal/domain/models"

type UrlMapperInterface interface {
	UrlRequestToUrl(urlRequest *models.UrlRequest) *models.Url
	UrlToUrlResponse(url *models.Url) *models.UrlResponse
	UrlsToUrlReponses(urls []models.Url) []models.UrlResponse
}

type UrlMapper struct{}

func NewUrlMapper() UrlMapperInterface {
	return &UrlMapper{}
}

func (mapper *UrlMapper) UrlRequestToUrl(urlRequest *models.UrlRequest) *models.Url {
	return &models.Url{
		Url:    urlRequest.Url,
		Hash:   urlRequest.Hash,
		UserID: urlRequest.UserID,
	}
}

func (mapper *UrlMapper) UrlToUrlResponse(url *models.Url) *models.UrlResponse {
	return &models.UrlResponse{
		Url:  url.Url,
		Hash: url.Hash,
	}
}

func (mapper *UrlMapper) UrlsToUrlReponses(urls []models.Url) []models.UrlResponse {
	var urlResponses []models.UrlResponse

	for _, url := range urls {
		urlResponse := mapper.UrlToUrlResponse(&url)
		urlResponses = append(urlResponses, *urlResponse)
	}

	return urlResponses
}
