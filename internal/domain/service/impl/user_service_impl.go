package service

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/ccallazans/url-shortener/internal/domain/mappers"
	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/models/roles"
	"github.com/ccallazans/url-shortener/internal/domain/repository"
	"github.com/ccallazans/url-shortener/internal/domain/service"
	"github.com/ccallazans/url-shortener/internal/utils"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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
		return errors.New(utils.USERNAME_ALREADY_EXISTS)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(utils.PASSWORD_HASH_ERROR)
	}

	userEntity := models.UserEntity{
		Username: user.Username,
		Password: string(hashedPassword),
		Role:     models.Role{Role: roles.USER},
	}

	err = s.userRepository.Save(ctx, &userEntity)
	if err != nil {
		return errors.New(utils.ENTITY_SAVE_ERROR)
	}

	return nil
}

func (s *userService) FindAll(ctx context.Context) ([]*models.User, error) {

	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return nil, errors.New(utils.DATA_NOT_FOUND)
	}

	return mappers.NewUserMapper().UserEntitiesToUser(users), nil
}

func (s *userService) FindById(ctx context.Context, id string) (*models.User, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New(utils.ATOI_ERROR)
	}

	user, err := s.userRepository.FindById(ctx, idInt)
	if err != nil {
		return nil, errors.New(utils.USER_NOT_FOUND)
	}

	return mappers.NewUserMapper().UserEntityToUser(user), nil
}

func (s *userService) FindByUsername(ctx context.Context, username string) (*models.User, error) {

	user, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.New(utils.USER_NOT_FOUND)
	}

	return mappers.NewUserMapper().UserEntityToUser(user), nil
}

func (s *userService) Update(ctx context.Context, user *models.User) error {
	return nil
}

func (s *userService) DeleteById(ctx context.Context, id int) error {
	return nil
}

func (s *userService) Auth(ctx context.Context, user *models.User) (string, error) {

	validUser, err := s.userRepository.FindByUsername(ctx, user.Username)
	if err != nil {
		return "", errors.New(utils.USER_NOT_FOUND)
	}

	err = bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New(utils.INVALID_PASSWORD)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.UserClaims{
		User: *mappers.NewUserMapper().UserEntityToUser(validUser),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_SECRET_KEY")))
	if err != nil {
		return "", errors.New(utils.TOKEN_GENERATE_ERROR)
	}

	return tokenString, nil
}
