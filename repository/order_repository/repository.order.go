package order_repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/edwinjordan/golang_test_dot.git/app/repository"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerOrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrder(db *gorm.DB) repository.CustomerOrderRepository {
	return &CustomerOrderRepositoryImpl{
		DB: db,
	}
}

// Create implements repository.CustomerOrderRepository.
func (repo *CustomerOrderRepositoryImpl) Create(ctx context.Context, order entity.CustomerOrder) entity.CustomerOrder {
	orderData := &CustomerOrder{}
	orderData = orderData.FromEntity(&order)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&orderData).Error
	helpers.PanicIfError(err)

	return *orderData.ToEntity()
}

// GenInvoice implements repository.CustomerOrderRepository.
func (repo *CustomerOrderRepositoryImpl) GenInvoice(ctx context.Context) string {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	invoice := map[string]interface{}{}

	month := fmt.Sprint(int(time.Now().Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	day := fmt.Sprint(time.Now().Day())
	if len(day) == 1 {
		day = "0" + day
	}
	year := fmt.Sprint(int(time.Now().Year()) % 1e2)

	date := day + month + year

	tx.WithContext(ctx).Table("tb_order").Select("IFNULL(order_inv_number,'') order_inv_number").Where("order_inv_number LIKE ?", "%"+date+"%").Order("order_inv_number DESC").Find(invoice)
	inv := ""
	if invoice["order_inv_number"] == nil {
		inv = "ORN-" + date + "-000"
	} else {
		inv = invoice["order_inv_number"].(string)
	}
	sort := inv[len(inv)-3:]
	newInv := inv[:len(inv)-3]
	str, _ := strconv.Atoi(sort)
	str += 1
	if str < 10 {
		sort = "00" + fmt.Sprint(str)
	} else if str < 100 {
		sort = "0" + fmt.Sprint(str)
	} else {
		sort = fmt.Sprint(str)
	}

	return newInv + sort
}
