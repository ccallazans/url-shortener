package utils

import "net/http"

// // Validation
const USERNAME_SHORT_ERROR = "short username error"
const USERNAME_LONG_ERROR = "error: long username"
const USERNAME_INVALID_CHAR_ERROR = "error: invalid characters"

const PASSWORD_SHORT_ERROR = "error: short password"
const PASSWORD_LONG_ERROR = "error: long password"

const URL_EMPTY_ERROR = "empty url error"
const URL_INVALID_ERROR = "invalid url error"

// // Handler
const INTERNAL_SERVER_ERROR = "an internal server error occurred."
const FORBIDDEN_ERROR = "forbidden error"

const CONTEXT_PARSE_ERROR = "failed to parse context"

const USER_CREATE_ERROR = "error: failed on create user"

// // Service
const USERNAME_ALREADY_EXISTS = "username already exists"
const HASH_ALREADY_EXISTS = "hash already exists"
const HASH_NOT_FOUND = "hash not found"
const USER_NOT_FOUND = "user not found"
const DATA_NOT_FOUND = "data not found"

const PASSWORD_HASH_ERROR = "failed to hash password"
const ENTITY_SAVE_ERROR = "failed to save entity to database"
const ATOI_ERROR = "error converting value"
const TOKEN_GENERATE_ERROR = "failed to generate token"

const INVALID_PASSWORD = "invalid password"

const REQUIRE_INTEGER = "require integer" //remove

// // MIDDLEWARE
const AUTHORIZATION_HEADER_EMPTY = "empty authorization header"
const AUTHORIZATION_HEADER_FORMAT_ERROR = "invalid Authorization header format"
const INVALID_SIGNING_METHOD = "invalid signing method"
const TOKEN_INVALID = "invalid token"
const BAD_REQUEST = "bad request"

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
	case INVALID_PASSWORD:
		info = StatusMessage{Status: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "The provided password is invalid."}
	case TOKEN_GENERATE_ERROR:
		info = StatusMessage{Status: http.StatusInternalServerError, ErrorType: "Internal Server Error", Message: "Failed to generate token. Please try again later."}
	case USERNAME_SHORT_ERROR:
		info = StatusMessage{Status: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The provided username is too short. Required at least 5 chars."}
	case USERNAME_LONG_ERROR:
		info = StatusMessage{Status: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The provided username is too long. Username length exceeds maximum of 15 characters"}
	case USERNAME_INVALID_CHAR_ERROR:
		info = StatusMessage{Status: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The username contains invalid characters. Only letters and numbers are allowed."}
	case PASSWORD_SHORT_ERROR:
		info = StatusMessage{Status: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The provided password is too short. Required at least 8 chars."}
	case PASSWORD_LONG_ERROR:
		info = StatusMessage{Status: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The provided password is too long. Password length exceeds maximum of 50 characters"}
	case URL_EMPTY_ERROR:
		info = StatusMessage{Status: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The URL field cannot be empty."}
	case URL_INVALID_ERROR:
		info = StatusMessage{Status: http.StatusBadRequest, ErrorType: "Bad Request", Message: "Invalid URL format. Please provide a valid URL."}
	case FORBIDDEN_ERROR:
		info = StatusMessage{Status: http.StatusForbidden, ErrorType: "Forbidden", Message: "You are not authorized to access this resource."}
	case HASH_NOT_FOUND:
		info = StatusMessage{Status: http.StatusNotFound, ErrorType: "Not Found", Message: "Resource not found"}

	default:
		info = StatusMessage{Status: http.StatusInternalServerError, ErrorType: "Internal Server Error", Message: err.Error()}
	}

	return info
}
