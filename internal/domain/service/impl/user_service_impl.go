package service

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/ccallazans/url-shortener/internal/domain/mappers"
	"github.com/ccallazans/url-shortener/internal/domain/models"
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
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

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
		return nil, err
	}

	userEntity, err := s.userRepository.FindById(ctx, idInt)
	if err != nil {
		return nil, errors.New(utils.REQUIRE_INTEGER)
	}

	return mappers.NewUserMapper().UserEntityToUser(userEntity), nil
}

func (s *userService) Update(ctx context.Context, user *models.User) error {
	return s.userRepository.Update(ctx, mappers.NewUserMapper().UserToUserEntity(user))
}

func (s *userService) DeleteById(ctx context.Context, id int) error {
	return s.userRepository.DeleteById(ctx, id)
}

func (s *userService) Auth(ctx context.Context, user *models.User) (string, error) {

	validUser, err := s.userRepository.FindByUsername(ctx, user.Username)
	if err != nil {
		return "", errors.New(utils.USERNAME_ALREADY_EXISTS)
	}

	err = bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New(utils.WRONG_PASSWORD)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.UserClaims{
		User: *mappers.NewUserMapper().UserEntityToUser(validUser),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_SECRET_KEY"))) // change this to your own secret key
	if err != nil {
		return "", errors.New(utils.TOKEN_GENERATE_ERROR)
	}

	return tokenString, nil
}
