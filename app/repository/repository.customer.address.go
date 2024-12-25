package repository

import (
	"context"

	"github.com/edwinjordan/golang_test_dot.git/entity"
)

type CustomerAddressRepository interface {
	Create(ctx context.Context, dataEn entity.CustomerAddress) entity.CustomerAddress
	FindSpesificData(ctx context.Context, where entity.CustomerAddress) []entity.CustomerAddress
	FindAll(ctx context.Context) []entity.CustomerAddress
}
