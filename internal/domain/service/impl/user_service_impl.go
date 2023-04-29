package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ccallazans/url-shortener/internal/domain/mappers"
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

	userExists, _ := s.userRepository.FindByUsername(ctx, user.Username)
	if userExists != nil {
		return errors.New("username already exists")
	}

	return s.userRepository.Save(ctx, mappers.NewUserMapper().UserToUserEntity(user))
}

func (s *userService) FindAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	usersEntity, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, userEntity := range usersEntity {
		users = append(users, mappers.NewUserMapper().UserEntityToUser(userEntity))
	}

	return users, nil
}

func (s *userService) FindById(ctx context.Context, id string) (*models.User, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Erro ao converter a string para inteiro: %v", err)
		return nil, err
	}

	userEntity, err := s.userRepository.FindById(ctx, idInt)
	if err != nil {
		return nil, err
	}

	return mappers.NewUserMapper().UserEntityToUser(userEntity), nil
}

func (s *userService) Update(ctx context.Context, user *models.User) error {
	return s.userRepository.Update(ctx, mappers.NewUserMapper().UserToUserEntity(user))
}

func (s *userService) DeleteById(ctx context.Context, id int) error {
	return s.userRepository.DeleteById(ctx, id)
}
