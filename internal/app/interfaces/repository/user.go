package repository

import (
	"context"
	"myapi/internal/app/domain"
	"myapi/internal/app/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUser {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Save(ctx context.Context, user *domain.User) error {

	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*domain.User, error) {

	var users []*domain.User
	result := r.db.Preload("Role").Preload("Urls").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *userRepository) FindById(ctx context.Context, id int) (*domain.User, error) {

	var user *domain.User
	result := r.db.Preload("Role").Preload("Urls").First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {

	var user *domain.User
	result := r.db.Preload("Role").Preload("Urls").First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}
