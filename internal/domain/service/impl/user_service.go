package service

import (
	"context"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/repository"
	"github.com/ccallazans/url-shortener/internal/domain/service"
)

type userService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) service.UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Save(ctx context.Context, user *models.User) error {
	return s.userRepository.Save(ctx, user)
}

func (s *userService) FindById(ctx context.Context, id int) (*models.User, error) {
	return s.userRepository.FindById(ctx, id)
}

func (s *userService) Update(ctx context.Context, user *models.User) error {
	return s.userRepository.Update(ctx, user)
}

func (s *userService) DeleteById(ctx context.Context, id int) error {
	return s.userRepository.DeleteById(ctx, id)
}
