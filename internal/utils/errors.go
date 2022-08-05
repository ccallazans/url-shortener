package utils

import "errors"

var (
	ErrWrongRequestBody = errors.New("wrong request body")
	ErrInvalidRequirements = errors.New("invalid requirements")
	ErrEmailDoNotExist  = errors.New("email no not exist")
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrInvalidPassword  = errors.New("wrong password")
	ErrGenerateJWT = errors.New("error generate jwt")
	ErrAddToDatabase = errors.New("error inserting into database")
	ErrValueDoNotExist = errors.New("value do not exist")
)
