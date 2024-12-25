package usecase_category

import (
	"net/http"

	"github.com/edwinjordan/golang_test_dot.git/app/repository"
	"github.com/edwinjordan/golang_test_dot.git/config"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/handler"
	"github.com/edwinjordan/golang_test_dot.git/pkg/exceptions"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang.org/x/net/html"
)

type UseCaseImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewUseCase(categoryRepo repository.CategoryRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:           validate,
		CategoryRepository: categoryRepo,
	}
}

func (controller *UseCaseImpl) Patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["categoryId"]
	dataRequest := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	_, err := controller.CategoryRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	if dataRequest["category_name"].(string) != "" {
		dataCategory := entity.Category{
			CategoryName: html.EscapeString(dataRequest["category_name"].(string)),
		}

		dataResponse := controller.CategoryRepository.Update(r.Context(), []string{"category_name"}, dataCategory, id)
		w.WriteHeader(http.StatusOK)
		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessUpdateData,
			"data":   dataResponse,
		}
		helpers.WriteToResponseBody(w, webResponse)
	} else {
		w.WriteHeader(http.StatusOK)
		webResponse := map[string]interface{}{
			"code":   400,
			"status": config.LoadMessage().GetDataByIdNotFound,
			//"data":   null,
		}

		helpers.WriteToResponseBody(w, webResponse)
	}
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["categoryId"]
	_, err := controller.CategoryRepository.FindById(r.Context(), id)
	// if err != nil {
	// 	panic(exceptions.NewNotFoundError(err.Error()))
	// }
	// webResponse := handler.WebResponse{
	// 	Error:   false,
	// 	Message: config.LoadMessage().SuccessDeleteData,
	// 	"code":   200
	// }
	if err != nil {
		webResponse := map[string]interface{}{
			"code":   404,
			"status": config.LoadMessage().GetDataByIdNotFound,
		}
		helpers.WriteToResponseBody(w, webResponse)
	} else {
		controller.CategoryRepository.Delete(r.Context(), id)

		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessDeleteData,
		}
		helpers.WriteToResponseBody(w, webResponse)
	}

}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["categoryId"]
	dataRequest := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	_, err := controller.CategoryRepository.FindById(r.Context(), id)
	if err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	if dataRequest["category_name"].(string) != "" {
		dataCategory := entity.Category{
			CategoryName: html.EscapeString(dataRequest["category_name"].(string)),
		}

		dataResponse := controller.CategoryRepository.Update(r.Context(), []string{"category_name"}, dataCategory, id)
		w.WriteHeader(http.StatusOK)
		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessUpdateData,
			"data":   dataResponse,
		}
		helpers.WriteToResponseBody(w, webResponse)
	} else {
		w.WriteHeader(http.StatusOK)
		webResponse := map[string]interface{}{
			"code":   400,
			"status": config.LoadMessage().GetDataByIdNotFound,
			//"data":   null,
		}

		helpers.WriteToResponseBody(w, webResponse)
	}

}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.Category{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	if dataRequest.CategoryName != "" {

		dataRequest.CategoryName = html.EscapeString(dataRequest.CategoryName)
		category := controller.CategoryRepository.Create(r.Context(), dataRequest)

		// dataResponse := map[string]interface{}{
		// 	"category": category,
		// }

		w.WriteHeader(http.StatusOK)
		webResponse := map[string]interface{}{
			"code":   200,
			"status": config.LoadMessage().SuccessCreateData,
			"data":   category,
		}

		helpers.WriteToResponseBody(w, webResponse)
	} else {
		w.WriteHeader(http.StatusOK)
		webResponse := map[string]interface{}{
			"code":   400,
			"status": config.LoadMessage().GetDataByIdNotFound,
		}

		helpers.WriteToResponseBody(w, webResponse)
	}

}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["categoryId"]
	dataResponse, err := controller.CategoryRepository.FindById(r.Context(), id)
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

func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {

	dataResponse := controller.CategoryRepository.FindAll(r.Context())

	webResponse := map[string]interface{}{
		"code":   200,
		"status": config.LoadMessage().SuccessGetData,
		"data":   dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
