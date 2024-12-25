package repository

import (
	"context"

	"github.com/edwinjordan/golang_test_dot.git/entity"
)

type CustomerOrderRepository interface {
	Create(ctx context.Context, order entity.CustomerOrder) entity.CustomerOrder
	GenInvoice(ctx context.Context) string
}

// Create implements CustomerOrderDetailRepository.
// func (c CustomerOrderRepository) Create(ctx context.Context, order entity.CustomerOrderDetail) entity.CustomerOrderDetail {
// 	panic("unimplemented")
// }

// Create implements CustomerOrderDetailRepository.
// func (c CustomerOrderRepository) Create(ctx context.Context, order entity.CustomerOrderDetail) entity.CustomerOrderDetail {
// 	panic("unimplemented")
// }
