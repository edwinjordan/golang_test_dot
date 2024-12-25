package usecase_customer

import (
	"fmt"
	"html"
	"net/http"
	"strconv"

	"github.com/edwinjordan/golang_test_dot.git/app/repository"
	"github.com/edwinjordan/golang_test_dot.git/config"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/handler"
	"github.com/edwinjordan/golang_test_dot.git/pkg/exceptions"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UseCaseImpl struct {
	CustomerRepository repository.CustomerRepository
	Validate           *validator.Validate
}

// Create implements UseCase.
func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.Customer{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	/* check if customer phonenumber exist */
	dataPhone := controller.CustomerRepository.FindSpesificData(r.Context(), entity.Customer{
		CustomerPhonenumber: dataRequest.CustomerPhonenumber,
	})

	if dataPhone != nil {
		panic(exceptions.NewConflictError("Nomor hp sudah digunakan, silahkan gunakan nomor hp lain atau masuk menggunakan nomor hp terdaftar"))
	}

	/* check if customer email exist */

	dataEmail := controller.CustomerRepository.FindSpesificData(r.Context(), entity.Customer{
		CustomerEmail: dataRequest.CustomerEmail,
	})

	if dataEmail != nil {
		panic(exceptions.NewConflictError("Email sudah digunakan, silahkan gunakan email lain"))
	}

	if dataRequest.CustomerName != "" {
		dataRequest.CustomerPassword = helpers.EncryptPassword(dataRequest.CustomerPassword)
		dataRequest.CustomerCode = controller.CustomerRepository.GenCustCode(r.Context())
		dataRequest.CustomerName = html.EscapeString(dataRequest.CustomerName)
		dataRequest.CustomerGender = html.EscapeString(dataRequest.CustomerGender)
		dataRequest.CustomerEmail = html.EscapeString(dataRequest.CustomerEmail)
		dataRequest.CustomerPhonenumber = html.EscapeString(dataRequest.CustomerPhonenumber)
		customer := controller.CustomerRepository.Create(r.Context(), dataRequest)

		// dataResponse := map[string]interface{}{
		// 	"customer": customer,
		// }

		// webResponse := handler.WebResponse{
		// 	Error:   false,
		// 	Message: config.LoadMessage().SuccessCreateData,
		// 	Data:    dataResponse,
		// }
		w.WriteHeader(http.StatusOK)
		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessCreateData,
			"data":   customer,
		}

		helpers.WriteToResponseBody(w, webResponse)
	} else {
		//w.WriteHeader(http.StatusBadRequest)
		webResponse := map[string]interface{}{
			"code":   400,
			"status": config.LoadMessage().GetDataByIdNotFound,
		}

		helpers.WriteToResponseBody(w, webResponse)
	}
}

// Delete implements UseCase.
func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customerId"]
	_, err := controller.CustomerRepository.FindById(r.Context(), id)
	// if err != nil {
	// 	panic(exceptions.NewNotFoundError(err.Error()))
	// }
	// controller.CustomerRepository.Delete(r.Context(), id)
	// webResponse := handler.WebResponse{
	// 	Error:   false,
	// 	Message: config.LoadMessage().SuccessDeleteData,
	// }
	// helpers.WriteToResponseBody(w, webResponse)
	if err != nil {
		webResponse := map[string]interface{}{
			"code":   404,
			"status": config.LoadMessage().GetDataByIdNotFound,
		}
		helpers.WriteToResponseBody(w, webResponse)
	} else {
		controller.CustomerRepository.Delete(r.Context(), id)

		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessDeleteData,
		}
		helpers.WriteToResponseBody(w, webResponse)
	}
}

// FindById implements UseCase.
func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customerId"]
	dataResponse, err := controller.CustomerRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

// Update implements UseCase.
func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customerId"]
	dataRequest := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	customer, err := controller.CustomerRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	if dataRequest["customer_name"].(string) != "" {
		password := customer.CustomerPassword

		/* check jika pengguna ingin mengubah kata sandinya */
		if dataRequest["customer_new_password"].(string) != "" {
			/* check kata sandi lamanya */
			checkPassword := bcrypt.CompareHashAndPassword([]byte(customer.CustomerPassword), []byte(dataRequest["customer_old_password"].(string)))
			if checkPassword != nil {
				panic(exceptions.NewBadRequestError("Tidak dapat mengubah kata sandi karena kata sandi lama tidak cocok"))
			}

			password = helpers.EncryptPassword(dataRequest["customer_new_password"].(string))
		}

		dataCustomer := entity.Customer{
			CustomerName:        html.EscapeString(dataRequest["customer_name"].(string)),
			CustomerGender:      dataRequest["customer_gender"].(string),
			CustomerPhonenumber: customer.CustomerPhonenumber,
			CustomerEmail:       customer.CustomerEmail,
			CustomerPassword:    password,
			CustomerUpdateAt:    helpers.CreateDateTime(),
		}
		dataResponse := controller.CustomerRepository.Update(r.Context(), []string{"customer_name", "customer_gender", "customer_phonenumber", "customer_email", "customer_password", "customer_update_at"}, dataCustomer, id)
		// webResponse := handler.WebResponse{
		// 	Error:   false,
		// 	Message: config.LoadMessage().SuccessUpdateData,
		// 	Data:    dataResponse,
		// }
		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessUpdateData,
			"data":   dataResponse,
		}
		helpers.WriteToResponseBody(w, webResponse)
	} else {
		webResponse := map[string]interface{}{
			"code":   400,
			"status": config.LoadMessage().GetDataByIdNotFound,
			//"data":   null,
		}

		helpers.WriteToResponseBody(w, webResponse)
	}
}

func NewUseCase(customerRepo repository.CustomerRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:           validate,
		CustomerRepository: customerRepo,
	}
}

func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	Qlimit := query.Get("limit")
	Qoffset := query.Get("offset")
	search := query.Get("search")

	if Qlimit == "" {
		Qlimit = "10"
	}

	if Qoffset == "" {
		Qoffset = "0"
	}

	limit, _ := strconv.Atoi(Qlimit)
	offset, _ := strconv.Atoi(Qoffset)

	nextOffset := limit + offset

	conf := map[string]interface{}{
		"limit":    limit,
		"offset":   offset,
		"search":   search,
		"customer": query.Get("customer"),
	}

	w.Header().Add("offset", fmt.Sprint(nextOffset))
	w.Header().Add("Access-Control-Expose-Headers", "offset")

	dataResponse := controller.CustomerRepository.FindAll(r.Context(), conf)
	// webResponse := handler.WebResponse{
	// 	Error:   false,
	// 	Message: config.LoadMessage().SuccessGetData,
	// 	Data:    dataResponse,
	// }

	webResponse := map[string]interface{}{
		"code":   200,
		"status": config.LoadMessage().SuccessGetData,
		"data":   dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
