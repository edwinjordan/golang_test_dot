package repository

import (
	"context"

	"github.com/edwinjordan/golang_test_dot.git/entity"
)

type CustomerOrderDetailRepository interface {
	Create(ctx context.Context, order entity.CustomerOrderDetail) entity.CustomerOrderDetail
}
