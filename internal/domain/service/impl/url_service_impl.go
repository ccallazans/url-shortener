package service

import (
	"context"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/repository"
	"github.com/ccallazans/url-shortener/internal/domain/service"
)

type urlService struct {
	urlRepository repository.UrlRepositoryInterface
}

func NewUrlService(urlRepository repository.UrlRepositoryInterface) service.UrlServiceInterface {
	return &urlService{
		urlRepository: urlRepository,
	}
}

func (s *urlService) Save(ctx context.Context, url *models.Url) error {
	return s.urlRepository.Save(ctx, url)
}

func (s *urlService) FindById(ctx context.Context, id int) (*models.Url, error) {
	return s.urlRepository.FindById(ctx, id)
}

func (s *urlService) Update(ctx context.Context, url *models.Url) error {
	return s.urlRepository.Update(ctx, url)
}

func (s *urlService) DeleteById(ctx context.Context, id int) error {
	return s.urlRepository.DeleteById(ctx, id) 
}
