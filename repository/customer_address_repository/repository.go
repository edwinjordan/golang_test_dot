package customer_address_repository

import (
	"context"

	"github.com/edwinjordan/golang_test_dot.git/app/repository"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerAddressRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.CustomerAddressRepository {
	return &CustomerAddressRepositoryImpl{
		DB: db,
	}
}

// Create implements repository.CustomerAddressRepository.
func (repo *CustomerAddressRepositoryImpl) Create(ctx context.Context, dataEn entity.CustomerAddress) entity.CustomerAddress {
	data := &CustomerAddress{}
	data = data.FromEntity(&dataEn)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&data).Error
	helpers.PanicIfError(err)

	// Hapus cache jika ada
	// cacheKey := "customer_address"
	// repo.redisClient.Del(ctx, cacheKey)

	return *data.ToEntity()
}

// FindAll implements repository.CustomerAddressRepository.
func (repo *CustomerAddressRepositoryImpl) FindAll(ctx context.Context) []entity.CustomerAddress {

	var tempData []entity.CustomerAddress

	category := []CustomerAddress{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&category).Error
	helpers.PanicIfError(err)

	for _, v := range category {
		tempData = append(tempData, *v.ToEntity())
	}

	return tempData
}

func (repo *CustomerAddressRepositoryImpl) FindSpesificData(ctx context.Context, where entity.CustomerAddress) []entity.CustomerAddress {
	data := []CustomerAddress{}
	dataWhere := &CustomerAddress{}
	dataWhere = dataWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Order("address_create_at DESC").Where(dataWhere).Find(&data).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerAddress
	for _, v := range data {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
