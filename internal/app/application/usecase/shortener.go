package usecase

import (
	"context"
	"errors"
	"math/rand"
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

func (u *ShortenerUsecase) Save(ctx context.Context, shortener *domain.Shortener) (*domain.Shortener, error) {

	shortener.Hash = generateHash()

	hashExists, _ := u.shortenerRepo.FindByHash(ctx, shortener.Hash)
	if hashExists != nil {
		return nil, errors.New(shared.HASH_ALREADY_EXISTS)
	}

	shortenerEntity := domain.Shortener{
		Url:  shortener.Url,
		Hash: shortener.Hash,
		User: shortener.User,
	}

	err := u.shortenerRepo.Save(ctx, &shortenerEntity)
	if err != nil {
		return nil, errors.New(shared.ENTITY_SAVE_ERROR)
	}

	return &shortenerEntity, nil
}

func (u *ShortenerUsecase) FindAll(ctx context.Context) ([]*domain.Shortener, error) {

	urls, err := u.shortenerRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (u *ShortenerUsecase) FindByHash(ctx context.Context, val string) (*domain.Shortener, error) {

	url, err := u.shortenerRepo.FindByHash(ctx, val)
	if err != nil {
		return nil, errors.New(shared.HASH_NOT_FOUND)
	}

	return url, nil
}

func generateHash() string {

	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	hash := make([]byte, 5)
	for i := range hash {
		hash[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(hash)
}
