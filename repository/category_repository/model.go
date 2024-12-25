package category_repository

import (
	"time"

	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"gorm.io/gorm"
)

type Category struct {
	CategoryId       string    `json:"category_id"`
	CategoryName     string    `json:"category_name"`
	CategoryDeleteAt time.Time `json:"category_delete_at"`
}

func (Category) TableName() string {
	return "ms_category"
}

func (model *Category) BeforeCreate(tx *gorm.DB) (err error) {
	model.CategoryId = helpers.GenUUID()
	return
}

func (Category) FromEntity(e *entity.Category) *Category {
	return &Category{
		CategoryId:       e.CategoryId,
		CategoryName:     e.CategoryName,
		CategoryDeleteAt: e.CategoryDeleteAt,
	}
}

func (model *Category) ToEntity() *entity.CategoryResponse {
	modelData := &entity.CategoryResponse{
		CategoryId:       model.CategoryId,
		CategoryName:     model.CategoryName,
		CategoryDeleteAt: model.CategoryDeleteAt,
	}
	return modelData
}
