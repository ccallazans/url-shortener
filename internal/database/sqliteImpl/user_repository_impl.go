package sqliteImpl

import (
	"context"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/domain/repository"
	"gorm.io/gorm"
)

type sqliteUserRepository struct {
	db *gorm.DB
}

func NewSqliteUserRepository(db *gorm.DB) repository.UserRepositoryInterface {
	return &sqliteUserRepository{
		db: db,
	}
}

func (r *sqliteUserRepository) Save(ctx context.Context, user *models.UserEntity) error {

	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *sqliteUserRepository) FindAll(ctx context.Context) ([]*models.UserEntity, error) {

	var users []*models.UserEntity
	result := r.db.Preload("Role").Preload("Urls").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *sqliteUserRepository) FindById(ctx context.Context, id int) (*models.UserEntity, error) {

	var user models.UserEntity
	result := r.db.Preload("Role").Preload("Urls").First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *sqliteUserRepository) FindByUsername(ctx context.Context, username string) (*models.UserEntity, error) {

	var user models.UserEntity
	result := r.db.Preload("Role").Preload("Urls").First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *sqliteUserRepository) Update(ctx context.Context, user *models.UserEntity) error {
	return nil
}

func (r *sqliteUserRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}
