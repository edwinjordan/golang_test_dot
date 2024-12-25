package order_repository

import (
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerOrderDetail struct {
	OrderDetailId       string  `gorm:"column:order_detail_id"`
	OrderDetailParentId string  `gorm:"column:order_detail_parent_id"`
	OrderDetailProduct  string  `gorm:"column:order_detail_product"`
	OrderDetailQty      int     `gorm:"column:order_detail_qty"`
	OrderDetailPrice    float64 `gorm:"column:order_detail_price"`
	OrderDetailSubtotal float64 `gorm:"column:order_detail_subtotal"`
}

func (CustomerOrderDetail) TableName() string {
	return "tb_order_detail"
}

func (model *CustomerOrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	model.OrderDetailId = helpers.GenUUID()
	return
}

func (CustomerOrderDetail) FromEntity(e *entity.CustomerOrderDetail) *CustomerOrderDetail {
	return &CustomerOrderDetail{
		OrderDetailId:       e.OrderDetailId,
		OrderDetailParentId: e.OrderDetailParentId,
		OrderDetailProduct:  e.OrderDetailProduct,
		OrderDetailQty:      e.OrderDetailQty,
		OrderDetailPrice:    e.OrderDetailPrice,
		OrderDetailSubtotal: e.OrderDetailSubtotal,
	}
}

func (model *CustomerOrderDetail) ToEntity() *entity.CustomerOrderDetail {
	modelData := &entity.CustomerOrderDetail{
		OrderDetailId:       model.OrderDetailId,
		OrderDetailParentId: model.OrderDetailParentId,
		OrderDetailProduct:  model.OrderDetailProduct,
		OrderDetailQty:      model.OrderDetailQty,
		OrderDetailPrice:    model.OrderDetailPrice,
		OrderDetailSubtotal: model.OrderDetailSubtotal,
	}

	return modelData
}
