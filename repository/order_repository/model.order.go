package order_repository

import (
	"time"

	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"github.com/edwinjordan/golang_test_dot.git/repository/customer_repository"
	"gorm.io/gorm"
)

type CustomerOrder struct {
	OrderId         string                        `gorm:"column:order_id"`
	OrderCustomerId string                        `gorm:"column:order_customer_id"`
	OrderInvNumber  string                        `gorm:"column:order_inv_number"`
	OrderTotalItem  int                           `gorm:"column:order_total_item"`
	OrderSubtotal   float64                       `gorm:"column:order_subtotal"`
	OrderDiscount   float64                       `gorm:"column:order_discount"`
	OrderTotal      float64                       `gorm:"column:order_total"`
	OrderNotes      string                        `gorm:"column:order_notes"`
	OrderStatus     int                           `gorm:"column:order_status"`
	OrderCreateAt   time.Time                     `gorm:"column:order_create_at"`
	OrderDetail     *[]CustomerOrderDetail        `gorm:"foreignKey:OrderDetailParentId;references:OrderId"`
	Customer        *customer_repository.Customer `gorm:"foreignKey:OrderCustomerId;references:CustomerId"`
}

func (CustomerOrder) TableName() string {
	return "tb_order"
}

func (model *CustomerOrder) BeforeCreate(tx *gorm.DB) (err error) {
	model.OrderId = helpers.GenUUID()
	model.OrderCreateAt = helpers.CreateDateTime()
	model.OrderStatus = 1
	return
}

func (CustomerOrder) FromEntity(e *entity.CustomerOrder) *CustomerOrder {
	return &CustomerOrder{
		OrderId:         e.OrderId,
		OrderCustomerId: e.OrderCustomerId,
		OrderInvNumber:  e.OrderInvNumber,
		OrderTotalItem:  e.OrderTotalItem,
		OrderSubtotal:   e.OrderSubtotal,
		OrderDiscount:   e.OrderDiscount,
		OrderTotal:      e.OrderTotal,
		OrderNotes:      e.OrderNotes,
		OrderStatus:     e.OrderStatus,
		OrderCreateAt:   e.OrderCreateAt,
	}
}

func (model *CustomerOrder) ToEntity() *entity.CustomerOrder {
	modelData := &entity.CustomerOrder{
		OrderId:         model.OrderId,
		OrderCustomerId: model.OrderCustomerId,
		OrderInvNumber:  model.OrderInvNumber,
		OrderTotalItem:  model.OrderTotalItem,
		OrderSubtotal:   model.OrderSubtotal,
		OrderDiscount:   model.OrderDiscount,
		OrderTotal:      model.OrderTotal,
		OrderNotes:      model.OrderNotes,
		OrderStatus:     model.OrderStatus,
		OrderCreateAt:   model.OrderCreateAt,
	}

	if model.OrderDetail != nil {
		var tempMenu []entity.CustomerOrderDetail
		for _, v := range *model.OrderDetail {
			tempMenu = append(tempMenu, *v.ToEntity())
		}
		modelData.OrderDetail = &tempMenu
	}

	if model.Customer != nil {
		modelData.Customer = model.Customer.ToEntity()
	}
	// if model.Address != nil {
	// 	modelData.Address = model.Address.ToEntity()
	// }

	return modelData
}
