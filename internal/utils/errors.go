package utils

//// Validation
const USERNAME_SHORT_ERROR = "short username error"
const USERNAME_LONG_ERROR = "error: long username"
const USERNAME_INVALID_CHAR_ERROR = "error: invalid characters"

const PASSWORD_SHORT_ERROR = "error: short password"
const PASSWORD_LONG_ERROR = "error: long password"

const URL_EMPTY_ERROR = "empty url error"
const URL_INVALID_ERROR = "invalid url error"

//// Handler
const INTERNAL_SERVER_ERROR = "An internal server error occurred."
const FORBIDDEN_ERROR = "You do not have permission to access this resource."

const CONTEXT_PARSE_ERROR = "failed to parse context"

const USER_CREATE_ERROR = "error: failed on create user"

//// Service
const USERNAME_ALREADY_EXISTS = "username already exists"
const HASH_ALREADY_EXISTS = "hash already exists"
const USER_NOT_FOUND = "user not found"
const DATA_NOT_FOUND = "data not found"

const PASSWORD_HASH_ERROR = "failed to hash password"
const ENTITY_SAVE_ERROR = "failed to save entity to database"
const ATOI_ERROR = "error converting value"
const TOKEN_GENERATE_ERROR = "failed to generate token"

const INVALID_PASSWORD = "invalid password"

const REQUIRE_INTEGER = "require integer" //remove

//// MIDDLEWARE
const AUTHORIZATION_HEADER_EMPTY = "empty authorization header"
const AUTHORIZATION_HEADER_FORMAT_ERROR = "invalid Authorization header format"
const INVALID_SIGNING_METHOD = "invalid signing method"
const TOKEN_INVALID = "invalid token"
const BAD_REQUEST = "bad request"