package repository

import (
	"context"

	"github.com/edwinjordan/golang_test_dot/entity"
)

type CategoryRepository interface {
	FindById(ctx context.Context, categoryId string) (entity.Category, error)
	FindAll(ctx context.Context) []entity.Category
}
