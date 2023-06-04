package factory

import (
	"errors"
	"myapi/internal/app/domain"
	"myapi/internal/app/shared"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	MAX_USERNAME_LENGHT = 15
	MIN_USERNAME_LENGHT = 5
	MAX_PASSWORD_LENGHT = 50
	MIN_PASSWORD_LENGHT = 8
)

func NewUserFactory(username string, password string) (domain.User, error) {

	verify := userVerify{}
	err := verify.execute(username, password)
	if err != nil {
		return domain.User{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, errors.New(shared.PASSWORD_HASH_ERROR)
	}

	return domain.User{
		UUID:       uuid.New(),
		Username:   username,
		Password:   string(hashedPassword),
		Role:       domain.USER_ROLE,
		Shorteners: []domain.Shortener{},
	}, nil
}

type UserValidator interface {
	execute(username string, password string)
	setNext(UserValidator)
}

//

type userVerify struct {
	next UserValidator
}

func (r *userVerify) setNext(next UserValidator) {
	r.next = next
}

func (v *userVerify) execute(username string, password string) error {
	if len(username) < MIN_USERNAME_LENGHT {
		return errors.New(shared.USERNAME_SHORT_ERROR)
	}

	if len(username) > MAX_USERNAME_LENGHT {
		return errors.New(shared.USERNAME_LONG_ERROR)
	}

	if len(password) < MIN_PASSWORD_LENGHT {
		return errors.New(shared.PASSWORD_SHORT_ERROR)
	}

	if len(password) > MAX_PASSWORD_LENGHT {
		return errors.New(shared.PASSWORD_LONG_ERROR)
	}

	return nil
}
