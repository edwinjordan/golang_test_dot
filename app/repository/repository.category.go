package repository

import (
	"context"

	"github.com/edwinjordan/golang_test_dot.git/entity"
)

type CategoryRepository interface {
	Create(ctx context.Context, category entity.Category) entity.CategoryResponse
	FindById(ctx context.Context, categoryId string) (entity.CategoryResponse, error)
	FindAll(ctx context.Context) []entity.CategoryResponse
	Update(ctx context.Context, selectField interface{}, category entity.Category, categoryId string) entity.CategoryResponse
	Patch(ctx context.Context, selectField interface{}, category entity.Category, categoryId string) entity.CategoryResponse
	Delete(ctx context.Context, categoryId string)
}
