package validation

import (
	"errors"
	"regexp"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/utils"
)

const MIN_USERNAME_LEN = 5
const MAX_USERNAME_LEN = 15

const MIN_PASSWORD_LEN = 8
const MAX_PASSWORD_LEN = 50

type UserRequestValidation interface {
	Execute(userRequest *models.UserRequest) error
}

//

type UsernameValidation struct {
	Next UserRequestValidation
}

func (v *UsernameValidation) Execute(userRequest *models.UserRequest) error {
	if len(userRequest.Username) < MIN_USERNAME_LEN {
		return errors.New(utils.USERNAME_SHORT_ERROR)
	}

	if len(userRequest.Username) > MAX_USERNAME_LEN {
		return errors.New(utils.USERNAME_LONG_ERROR)
	}

	var validUsername = regexp.MustCompile("^[a-zA-Z0-9]+$")
	if !validUsername.MatchString(userRequest.Username) {
		return errors.New(utils.USERNAME_INVALID_CHAR_ERROR)
	}

	v.Next.Execute(userRequest)
	return nil
}

//

type PasswordValidation struct {
	Next UserRequestValidation
}

func (v *PasswordValidation) Execute(userRequest *models.UserRequest) error {
	if len(userRequest.Password) < MIN_PASSWORD_LEN {
		return errors.New(utils.PASSWORD_SHORT_ERROR)
	}

	if len(userRequest.Password) > MAX_PASSWORD_LEN {
		return errors.New(utils.PASSWORD_LONG_ERROR)
	}

	return nil
}
