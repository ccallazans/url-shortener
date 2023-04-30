package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/repository"
	"github.com/ccallazans/url-shortener/internal/domain/service"
	"github.com/ccallazans/url-shortener/internal/utils"
)

type urlService struct {
	urlRepository repository.UrlRepositoryInterface
}

func NewUrlService(urlRepository repository.UrlRepositoryInterface) service.UrlServiceInterface {
	return &urlService{
		urlRepository: urlRepository,
	}
}

func (s *urlService) Save(ctx context.Context, url *models.Url) (*models.Url, error) {

	if !url.HasHash() {
		newHash := utils.GenerateHash()
		url.Hash = newHash
	}

	hashExists, _ := s.urlRepository.FindByHash(ctx, url.Hash)
	if hashExists != nil {
		return nil, errors.New(utils.HASH_ALREADY_EXISTS)
	}

	err := s.urlRepository.Save(ctx, url)
	if err != nil {
		return nil, errors.New(utils.ENTITY_SAVE_ERROR)
	}

	return url, nil
}

func (s *urlService) FindAll(ctx context.Context) ([]*models.Url, error) {

	urls, err := s.urlRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (s *urlService) FindById(ctx context.Context, id string) (*models.Url, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	url, err := s.urlRepository.FindById(ctx, idInt)
	if err != nil {
		return nil, errors.New(utils.REQUIRE_INTEGER)
	}

	return url, nil
}

func (s *urlService) FindByHash(ctx context.Context, hash string) (*models.Url, error) {

	url, err := s.urlRepository.FindByHash(ctx, hash)
	if err != nil {
		return nil, errors.New(utils.HASH_NOT_FOUND)
	}

	return url, nil
}

func (s *urlService) Update(ctx context.Context, url *models.Url) error {
	return s.urlRepository.Update(ctx, url)
}

func (s *urlService) DeleteById(ctx context.Context, id int) error {
	return s.urlRepository.DeleteById(ctx, id)
}
