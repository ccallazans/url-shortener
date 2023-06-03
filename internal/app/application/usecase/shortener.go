package usecase

import (
	"context"
	"errors"
	"myapi/internal/app/domain"
	"myapi/internal/app/domain/repository"
	"myapi/internal/app/shared"
)

type ShortenerUsecase struct {
	shortenerRepo repository.IShortener
}

func NewShortenerUsecase(shortenerRepo repository.IShortener) ShortenerUsecase {
	return ShortenerUsecase{
		shortenerRepo: shortenerRepo,
	}
}

func (u *ShortenerUsecase) Save(ctx context.Context, shortener *domain.Shortener) error {

	hashExists, _ := u.shortenerRepo.FindByHash(ctx, shortener.Hash)
	if hashExists != nil {
		return errors.New(shared.HASH_ALREADY_EXISTS)
	}

	shortenerEntity := domain.Shortener{
		Url:  shortener.Url,
		Hash: shortener.Hash,
		User: shortener.User,
	}

	err := u.shortenerRepo.Save(ctx, &shortenerEntity)
	if err != nil {
		return errors.New(shared.ENTITY_SAVE_ERROR)
	}

	return nil
}

func (u *ShortenerUsecase) FindAll(ctx context.Context) ([]*domain.Shortener, error) {

	urls, err := u.shortenerRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (u *ShortenerUsecase) FindByHash(ctx context.Context, hash string) (*domain.Shortener, error) {

	url, err := u.shortenerRepo.FindByHash(ctx, hash)
	if err != nil {
		return nil, errors.New(shared.HASH_NOT_FOUND)
	}

	return url, nil
}
