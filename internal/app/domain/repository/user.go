package repository

import (
	"context"
	"myapi/internal/app/domain"
)

type IUser interface {
	Save(ctx context.Context, user domain.User) error
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByUUID(ctx context.Context, uuid string) (domain.User, error)
	FindByUsername(ctx context.Context, username string) (domain.User, error)
	DeleteById(ctx context.Context, id int) error
}
