package utils

import (
	"math/rand"
	"net/http"
	"os"
)

func ReadSqlFile(path string) (string, error) {

	sql_query, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(sql_query), nil
}

func GenerateHash() string {
	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	hash := make([]byte, 5)
	for i := range hash {
		hash[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(hash)
}

type StatusMessage struct {
	Status    int
	ErrorType string
	Message   string
}

func MatchError(err error) StatusMessage {
	var info StatusMessage

	switch err.Error() {
	case USERNAME_ALREADY_EXISTS:
		info = StatusMessage{Status: http.StatusConflict, ErrorType: "Conflict", Message: "User already exists."}
	case PASSWORD_HASH_ERROR:
		info = StatusMessage{Status: http.StatusInternalServerError, ErrorType: "Internal Server Error", Message: "An error occurred while generating token."}
	case HASH_ALREADY_EXISTS:
		info = StatusMessage{Status: http.StatusConflict, ErrorType: "Conflict", Message: "Hash already exists."}
	case ENTITY_SAVE_ERROR:
		info = StatusMessage{Status: http.StatusInternalServerError, ErrorType: "Internal Server Error", Message: "An error occurred while attempting to save the entity to the database."}
	case ATOI_ERROR:
		info = StatusMessage{Status: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The parameter must be of type integer."}
	case USER_NOT_FOUND:
		info = StatusMessage{Status: http.StatusNotFound, ErrorType: "Not Found", Message: "User not found."}
	case DATA_NOT_FOUND:
		info = StatusMessage{Status: http.StatusNotFound, ErrorType: "Not Found", Message: "No data found for the requested resource."}
	case AUTHORIZATION_HEADER_EMPTY:
		info = StatusMessage{Status: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "Authorization header is empty."}
	case AUTHORIZATION_HEADER_FORMAT_ERROR:
		info = StatusMessage{Status: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "Invalid authorization header format. Expected 'Bearer <token>'."}
	case INVALID_SIGNING_METHOD:
		info = StatusMessage{Status: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "Invalid signing method used in the authorization header."}
	case TOKEN_INVALID:
		info = StatusMessage{Status: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "The access token provided is invalid or has expired."}
	case BAD_REQUEST:
		info = StatusMessage{Status: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The request body is invalid."}

	default:
		info = StatusMessage{Status: http.StatusInternalServerError, ErrorType: "Internal Server Error"}
	}

	return info
}
