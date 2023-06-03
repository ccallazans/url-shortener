package shared

import "net/http"

type ResponseError struct {
	StatusCode int    `json:"code"`
	ErrorType  string `json:"type"`
	Message    string `json:"message"`
}

func HandleError(err error) ResponseError {
	switch err.Error() {
	case USERNAME_ALREADY_EXISTS:
		return ResponseError{StatusCode: http.StatusConflict, ErrorType: "Conflict", Message: "User already exists."}
	case PASSWORD_HASH_ERROR:
		return ResponseError{StatusCode: http.StatusInternalServerError, ErrorType: "Internal Server Error", Message: "An error occurred while generating token."}
	case HASH_ALREADY_EXISTS:
		return ResponseError{StatusCode: http.StatusConflict, ErrorType: "Conflict", Message: "Hash already exists."}
	case ENTITY_SAVE_ERROR:
		return ResponseError{StatusCode: http.StatusInternalServerError, ErrorType: "Internal Server Error", Message: "An error occurred while attempting to save the entity to the database."}
	case ATOI_ERROR:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The parameter must be of type integer."}
	case USER_NOT_FOUND:
		return ResponseError{StatusCode: http.StatusNotFound, ErrorType: "Not Found", Message: "User not found."}
	case DATA_NOT_FOUND:
		return ResponseError{StatusCode: http.StatusNotFound, ErrorType: "Not Found", Message: "No data found for the requested resource."}
	case AUTHORIZATION_HEADER_EMPTY:
		return ResponseError{StatusCode: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "Authorization header is empty."}
	case AUTHORIZATION_HEADER_FORMAT_ERROR:
		return ResponseError{StatusCode: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "Invalid authorization header format. Expected 'Bearer <token>'."}
	case INVALID_SIGNING_METHOD:
		return ResponseError{StatusCode: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "Invalid signing method used in the authorization header."}
	case TOKEN_INVALID:
		return ResponseError{StatusCode: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "The access token provided is invalid or has expired."}
	case BAD_REQUEST:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The request body is invalid."}
	case INVALID_PASSWORD:
		return ResponseError{StatusCode: http.StatusUnauthorized, ErrorType: "Unauthorized", Message: "The provided password is invalid."}
	case TOKEN_GENERATE_ERROR:
		return ResponseError{StatusCode: http.StatusInternalServerError, ErrorType: "Internal Server Error", Message: "Failed to generate token. Please try again later."}
	case USERNAME_SHORT_ERROR:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The provided username is too short. Required at least 5 chars."}
	case USERNAME_LONG_ERROR:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The provided username is too long. Username length exceeds maximum of 15 characters"}
	case USERNAME_INVALID_CHAR_ERROR:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The username contains invalid characters. Only letters and numbers are allowed."}
	case PASSWORD_SHORT_ERROR:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The provided password is too short. Required at least 8 chars."}
	case PASSWORD_LONG_ERROR:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The provided password is too long. Password length exceeds maximum of 50 characters"}
	case URL_EMPTY_ERROR:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "The URL field cannot be empty."}
	case CONTEXT_PARSE_ERROR:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "Invalid input provided."}
	case URL_INVALID_ERROR:
		return ResponseError{StatusCode: http.StatusBadRequest, ErrorType: "Bad Request", Message: "Invalid URL format. Please provide a valid URL."}
	case FORBIDDEN_ERROR:
		return ResponseError{StatusCode: http.StatusForbidden, ErrorType: "Forbidden", Message: "You are not authorized to access this resource."}
	case HASH_NOT_FOUND:
		return ResponseError{StatusCode: http.StatusNotFound, ErrorType: "Not Found", Message: "Resource not found"}

	default:
		return ResponseError{StatusCode: http.StatusInternalServerError, ErrorType: "Internal Server Error", Message: err.Error()}
	}
}

const (
	// Validation
	USERNAME_SHORT_ERROR        = "short username error"
	USERNAME_LONG_ERROR         = "error: long username"
	USERNAME_INVALID_CHAR_ERROR = "error: invalid characters"
	PASSWORD_SHORT_ERROR        = "error: short password"
	PASSWORD_LONG_ERROR         = "error: long password"
	URL_EMPTY_ERROR             = "empty url error"
	URL_INVALID_ERROR           = "invalid url error"

	// Handler
	INTERNAL_SERVER_ERROR = "an internal server error occurred."
	FORBIDDEN_ERROR       = "forbidden error"
	CONTEXT_PARSE_ERROR   = "failed to parse context"

	// Usecase
	USERNAME_ALREADY_EXISTS = "username already exists"
	HASH_ALREADY_EXISTS     = "hash already exists"
	HASH_NOT_FOUND          = "hash not found"
	USER_NOT_FOUND          = "user not found"
	DATA_NOT_FOUND          = "data not found"
	PASSWORD_HASH_ERROR     = "failed to hash password"
	ENTITY_SAVE_ERROR       = "failed to save entity to database"
	ATOI_ERROR              = "error converting value"
	TOKEN_GENERATE_ERROR    = "failed to generate token"
	INVALID_PASSWORD        = "invalid password"

	// Middleware
	AUTHORIZATION_HEADER_EMPTY        = "empty authorization header"
	AUTHORIZATION_HEADER_FORMAT_ERROR = "invalid Authorization header format"
	INVALID_SIGNING_METHOD            = "invalid signing method"
	TOKEN_INVALID                     = "invalid token"
	BAD_REQUEST                       = "bad request"
)
