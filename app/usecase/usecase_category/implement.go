package usecase_category

import (
	"net/http"

	"github.com/edwinjordan/golang_test_dot/app/repository"
	"github.com/edwinjordan/golang_test_dot/config"
	"github.com/edwinjordan/golang_test_dot/handler"
	"github.com/edwinjordan/golang_test_dot/pkg/exceptions"
	"github.com/edwinjordan/golang_test_dot/pkg/helpers"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
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
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
