package usecase

import (
	"context"
	"errors"

	"myapi/internal/app/application/usecase/auth"
	"myapi/internal/app/application/usecase/factory"
	"myapi/internal/app/domain"

	"myapi/internal/app/domain/repository"
	"myapi/internal/app/shared"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo repository.IUser
}

func NewUserUsecase(userRepo repository.IUser) UserUsecase {
	return UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Save(ctx context.Context, user domain.User) error {

	_, err := u.userRepo.FindByUsername(ctx, user.Username)
	if err == nil {
		return errors.New(shared.USERNAME_ALREADY_EXISTS)
	}

	userEntity, err := factory.NewUserFactory(user.Username, user.Password)
	if err != nil {
		return err
	}

	err = u.userRepo.Save(ctx, userEntity)
	if err != nil {
		return errors.New(shared.ENTITY_SAVE_ERROR)
	}

	return nil
}

func (u *UserUsecase) FindAll(ctx context.Context) ([]domain.User, error) {

	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, errors.New(shared.DATA_NOT_FOUND)
	}

	return users, nil
}

func (u *UserUsecase) FindByUUID(ctx context.Context, uuid string) (domain.User, error) {

	user, err := u.userRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return domain.User{}, errors.New(shared.USER_NOT_FOUND)
	}

	return user, nil
}

func (s *UserUsecase) FindByUsername(ctx context.Context, username string) (domain.User, error) {

	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return domain.User{}, errors.New(shared.USER_NOT_FOUND)
	}

	return user, nil
}

func (s *UserUsecase) Update(ctx context.Context, user domain.User) error {
	return nil
}

func (s *UserUsecase) DeleteById(ctx context.Context, id int) error {
	return nil
}

func (u *UserUsecase) Auth(ctx context.Context, user domain.User) (string, error) {

	validUser, err := u.userRepo.FindByUsername(ctx, user.Username)
	if err != nil {
		return "", errors.New(shared.USER_NOT_FOUND)
	}

	err = bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New(shared.INVALID_PASSWORD)
	}

	token, err := GenerateJWT(
		auth.UserAuth{
			UUID:     validUser.UUID,
			Username: validUser.Username,
			Role:     validUser.Role,
		},
	)
	if err != nil {
		return "", errors.New(shared.TOKEN_GENERATE_ERROR)
	}

	return token, nil
}

func GenerateJWT(user auth.UserAuth) (string, error) {

	claims := &auth.JWTClaim{
		User: user,
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
	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
