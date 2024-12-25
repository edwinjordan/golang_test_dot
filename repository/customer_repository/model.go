package customer_repository

import (
	"time"

	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"github.com/edwinjordan/golang_test_dot.git/repository/customer_address_repository"
	"gorm.io/gorm"
)

type Customer struct {
	CustomerId          string                                         `gorm:"column:customer_id"`
	CustomerCode        string                                         `gorm:"column:customer_code"`
	CustomerName        string                                         `gorm:"column:customer_name"`
	CustomerGender      string                                         `gorm:"column:customer_gender"`
	CustomerPhonenumber string                                         `gorm:"column:customer_phonenumber"`
	CustomerEmail       string                                         `gorm:"column:customer_email"`
	CustomerPassword    string                                         `gorm:"column:customer_password"`
	CustomerStatus      int                                            `gorm:"column:customer_status"`
	CustomerCreateAt    time.Time                                      `gorm:"column:customer_create_at"`
	CustomerUpdateAt    time.Time                                      `gorm:"column:customer_update_at"`
	Address             *[]customer_address_repository.CustomerAddress `gorm:"foreignKey:AddressCustomerId;references:CustomerId"`
}

func (Customer) TableName() string {
	return "tb_customer"
}

func (model *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	model.CustomerId = helpers.GenUUID()
	model.CustomerCreateAt = helpers.CreateDateTime()
	model.CustomerUpdateAt = helpers.CreateDateTime()
	//model.CustomerLastStatus = 0
	model.CustomerStatus = 0
	return
}

func (Customer) FromEntity(e *entity.Customer) *Customer {
	return &Customer{
		CustomerId:          e.CustomerId,
		CustomerCode:        e.CustomerCode,
		CustomerName:        e.CustomerName,
		CustomerGender:      e.CustomerGender,
		CustomerPhonenumber: e.CustomerPhonenumber,
		CustomerEmail:       e.CustomerEmail,
		CustomerPassword:    e.CustomerPassword,
		CustomerStatus:      e.CustomerStatus,
		CustomerCreateAt:    e.CustomerCreateAt,
		CustomerUpdateAt:    e.CustomerUpdateAt,
	}
}

func (model *Customer) ToEntity() *entity.CustomerResponse {
	modelData := &entity.CustomerResponse{
		CustomerId:          model.CustomerId,
		CustomerCode:        model.CustomerCode,
		CustomerName:        model.CustomerName,
		CustomerGender:      model.CustomerGender,
		CustomerPhonenumber: model.CustomerPhonenumber,
		CustomerEmail:       model.CustomerEmail,
		CustomerPassword:    model.CustomerPassword,
		CustomerStatus:      model.CustomerStatus,
		CustomerCreateAt:    model.CustomerCreateAt,
		CustomerUpdateAt:    model.CustomerUpdateAt,
	}

	if model.Address != nil {
		var tempMenu []entity.CustomerAddress
		for _, v := range *model.Address {
			tempMenu = append(tempMenu, *v.ToEntity())
		}
		modelData.Address = &tempMenu
	}
	return modelData
}
