package customer_address_repository

import (
	"time"

	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"

	//	"github.com/edwinjordan/golang_test_dot.git/helpers"
	"gorm.io/gorm"
)

type CustomerAddress struct {
	AddressId         string    `gorm:"column:address_id"`
	AddressCustomerId string    `gorm:"column:address_customer_id"`
	AddressText       string    `gorm:"column:address_text"`
	AddressName       string    `gorm:"column:address_name"`
	AddressPostalCode string    `gorm:"column:address_postal_code"`
	AddressCreateAt   time.Time `gorm:"column:address_create_at"`
	AddressUpdateAt   time.Time `gorm:"column:address_update_at"`
}

func (CustomerAddress) TableName() string {
	return "tb_customer_address"
}

func (model *CustomerAddress) BeforeCreate(tx *gorm.DB) (err error) {
	model.AddressId = helpers.GenUUID()
	model.AddressCreateAt = helpers.CreateDateTime()
	model.AddressUpdateAt = helpers.CreateDateTime()
	return
}

func (CustomerAddress) FromEntity(e *entity.CustomerAddress) *CustomerAddress {
	return &CustomerAddress{
		AddressId:         e.AddressId,
		AddressCustomerId: e.AddressCustomerId,
		AddressText:       e.AddressText,
		AddressName:       e.AddressName,
		AddressPostalCode: e.AddressPostalCode,
		AddressCreateAt:   e.AddressCreateAt,
		AddressUpdateAt:   e.AddressUpdateAt,
	}
}

func (model *CustomerAddress) ToEntity() *entity.CustomerAddress {
	modelData := &entity.CustomerAddress{
		AddressId:         model.AddressId,
		AddressCustomerId: model.AddressCustomerId,
		AddressText:       model.AddressText,
		AddressName:       model.AddressName,
		AddressPostalCode: model.AddressPostalCode,
		AddressCreateAt:   model.AddressCreateAt,
		AddressUpdateAt:   model.AddressUpdateAt,
	}

	return modelData
}
