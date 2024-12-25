package usecase_customer_address

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
	CustomerAddressRepository repository.CustomerAddressRepository
	Validate                  *validator.Validate
}

func NewUseCase(addressRepo repository.CustomerAddressRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:                  validate,
		CustomerAddressRepository: addressRepo,
	}
}

// Create implements UseCase.
func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.CustomerAddress{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	if dataRequest.AddressCustomerId != "" {
		dataRequest.AddressName = html.EscapeString(dataRequest.AddressName)
		dataRequest.AddressCustomerId = html.EscapeString(dataRequest.AddressCustomerId)
		dataRequest.AddressText = html.EscapeString(dataRequest.AddressText)
		dataRequest.AddressPostalCode = html.EscapeString(dataRequest.AddressPostalCode)

		dataResponse := controller.CustomerAddressRepository.Create(r.Context(), dataRequest)
		// webResponse := handler.WebResponse{
		// 	Error:   false,
		// 	Message: config.LoadMessage().SuccessCreateData,
		// 	Data:    dataResponse,
		// }
		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessCreateData,
			"data":   dataResponse,
		}
		helpers.WriteToResponseBody(w, webResponse)
	} else {
		webResponse := map[string]interface{}{
			"code":   400,
			"status": config.LoadMessage().GetDataByIdNotFound,
		}

		helpers.WriteToResponseBody(w, webResponse)
	}
}

// FindAll implements UseCase.
func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	dataResponse := controller.CustomerAddressRepository.FindSpesificData(r.Context(), entity.CustomerAddress{
		AddressCustomerId: vars.Get("customer_id"),
	})

	webResponse := map[string]interface{}{
		"code":   200,
		"status": config.LoadMessage().SuccessGetData,
		"data":   dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
