package repository

import (
	"context"
	"myapi/internal/app/domain"
)

type IRole interface {
	FindById(ctx context.Context, id int) (*domain.Role, error)
}
