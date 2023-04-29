package mappers

import "github.com/ccallazans/url-shortener/internal/domain/models"

type UserMapperInterface interface {
	UserEntityToUser(userEntity *models.UserEntity) *models.User

	UserToUserResponse(user *models.User) *models.UserResponse
	UserToUserEntity(user *models.User) *models.UserEntity
	UsersToUserResponses(users []*models.User) []*models.UserResponse

	UserRequestToUser(userRequest *models.UserRequest) *models.User
	UserRequestToUserResponse(userRequest *models.UserRequest) *models.UserResponse
}

type UserMapper struct{}

func NewUserMapper() UserMapperInterface {
	return &UserMapper{}
}

func (mapper *UserMapper) UserEntityToUser(userEntity *models.UserEntity) *models.User {
	return &models.User{
		ID:       userEntity.ID,
		Username: userEntity.Username,
		Password: userEntity.Password,
	}
}

func (mapper *UserMapper) UserToUserResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{
		Username: user.Username,
		Urls:     user.Urls,
	}
}

func (mapper *UserMapper) UserToUserEntity(user *models.User) *models.UserEntity {
	return &models.UserEntity{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}
}

func (mapper *UserMapper) UserRequestToUser(userRequest *models.UserRequest) *models.User {
	return &models.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}
}

func (mapper *UserMapper) UsersToUserResponses(users []*models.User) []*models.UserResponse {
	var userResponses []*models.UserResponse

	for _, user := range users {
		userResponse := mapper.UserToUserResponse(user)
		userResponses = append(userResponses, userResponse)
	}

	return userResponses
}

func (mapper *UserMapper) UserRequestToUserResponse(userRequest *models.UserRequest) *models.UserResponse {
	return &models.UserResponse{
		Username: userRequest.Username,
	}
}