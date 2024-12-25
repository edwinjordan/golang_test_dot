package usecase_order

import (
	"html"
	"net/http"

	"github.com/edwinjordan/golang_test_dot.git/app/repository"
	"github.com/edwinjordan/golang_test_dot.git/config"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type UseCaseImpl struct {
	CustomerOrderRepository       repository.CustomerOrderRepository
	CustomerOrderDetailRepository repository.CustomerOrderDetailRepository
	Validate                      *validator.Validate
}

func NewUseCase(
	orderRepo repository.CustomerOrderRepository,
	orderDetailRepo repository.CustomerOrderDetailRepository,
	validate *validator.Validate,
) UseCase {
	return &UseCaseImpl{
		CustomerOrderRepository:       orderRepo,
		CustomerOrderDetailRepository: orderDetailRepo,
		Validate:                      validate,
	}

}

// Create implements UseCase.
func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	dataParent := dataRequest["parent"].(map[string]interface{})

	dataResponse := controller.CustomerOrderRepository.Create(r.Context(), entity.CustomerOrder{
		OrderInvNumber:  controller.CustomerOrderRepository.GenInvoice(r.Context()),
		OrderCustomerId: dataParent["order_customer_id"].(string),
		OrderTotalItem:  int(dataParent["order_total_item"].(float64)),
		OrderSubtotal:   dataParent["order_subtotal"].(float64),
		OrderDiscount:   dataParent["order_discount"].(float64),
		OrderTotal:      dataParent["order_total"].(float64),
		OrderNotes:      html.EscapeString(dataParent["order_notes"].(string)),
	})

	for _, v := range dataRequest["detail"].([]interface{}) {
		dt := v.(map[string]interface{})
		controller.CustomerOrderDetailRepository.Create(r.Context(), entity.CustomerOrderDetail{
			OrderDetailParentId: dataResponse.OrderId,
			OrderDetailProduct:  dt["product_nama"].(string),
			OrderDetailQty:      int(dt["product_qty"].(float64)),
			OrderDetailPrice:    dt["product_price"].(float64),
			OrderDetailSubtotal: (dt["product_qty"].(float64) * dt["product_price"].(float64)),
		})
	}
	webResponse := map[string]interface{}{
		"code":   200,
		"status": config.LoadMessage().SuccessCreateData,
		"data":   dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
