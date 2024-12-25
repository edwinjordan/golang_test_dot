package order_repository

import (
	"context"

	"github.com/edwinjordan/golang_test_dot.git/app/repository"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerOrderDetailRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderDetail(db *gorm.DB) repository.CustomerOrderDetailRepository {
	return &CustomerOrderDetailRepositoryImpl{
		DB: db,
	}
}

// Create implements repository.CustomerOrderDetailRepository.
func (repo *CustomerOrderDetailRepositoryImpl) Create(ctx context.Context, order entity.CustomerOrderDetail) entity.CustomerOrderDetail {
	orderData := &CustomerOrderDetail{}
	orderData = orderData.FromEntity(&order)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&orderData).Error
	helpers.PanicIfError(err)

	return *orderData.ToEntity()
}
