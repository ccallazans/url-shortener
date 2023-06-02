package usecase

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"myapi/internal/app/domain"
	"myapi/internal/app/domain/repository"
	"myapi/internal/app/interfaces/auth"
	"myapi/utils"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const ()

type UserUsecase struct {
	userRepo repository.IUser
}

func NewUserUsecase(userRepo repository.IUser) UserUsecase {
	return UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Save(ctx context.Context, user *domain.User) error {

	userExists, _ := u.userRepo.FindByUsername(ctx, user.Username)
	if userExists != nil {
		return errors.New(utils.USERNAME_ALREADY_EXISTS)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(utils.PASSWORD_HASH_ERROR)
	}

	userEntity := domain.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Role:     domain.Role{Role: domain.USER},
	}

	err = u.userRepo.Save(ctx, &userEntity)
	if err != nil {
		return errors.New(utils.ENTITY_SAVE_ERROR)
	}

	return nil
}

func (u *UserUsecase) FindAll(ctx context.Context) ([]*domain.User, error) {

	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, errors.New(utils.DATA_NOT_FOUND)
	}

	return users, nil
}

func (u *UserUsecase) FindById(ctx context.Context, id string) (*domain.User, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New(utils.ATOI_ERROR)
	}

	user, err := u.userRepo.FindById(ctx, idInt)
	if err != nil {
		return nil, errors.New(utils.USER_NOT_FOUND)
	}

	return user, nil
}

func (s *UserUsecase) FindByUsername(ctx context.Context, username string) (*domain.User, error) {

	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.New(utils.USER_NOT_FOUND)
	}

	return user, nil
}

func (s *UserUsecase) Update(ctx context.Context, user *domain.User) error {
	return nil
}

func (s *UserUsecase) DeleteById(ctx context.Context, id int) error {
	return nil
}

func (u *UserUsecase) Auth(ctx context.Context, user *domain.User) (string, error) {

	validUser, err := u.userRepo.FindByUsername(ctx, user.Username)
	if err != nil {
		return "", errors.New(utils.USER_NOT_FOUND)
	}

	err = bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New(utils.INVALID_PASSWORD)
	}

	token, err := GenerateJWT(
		&auth.UserAuth{
			UUID:     validUser.UUID,
			Username: validUser.Username,
			Role:     validUser.Role.Role,
		},
	)
	if err != nil {
		return "", errors.New(utils.TOKEN_GENERATE_ERROR)
	}

	return *token, nil
}

func GenerateJWT(user *auth.UserAuth) (*string, error) {

	claims := &auth.JWTClaim{
		User: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "localhost",
			Subject:   user.UUID.String(),
			Audience:  jwt.ClaimStrings{"localhost"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
