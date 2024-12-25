package router

import (
	"github.com/edwinjordan/golang_test_dot.git/app/usecase/usecase_customer_address"
	"github.com/edwinjordan/golang_test_dot.git/repository/customer_address_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CustomerAddressRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	addressRepository := customer_address_repository.New(db)
	addressController := usecase_customer_address.NewUseCase(addressRepository, validate)

	router.HandleFunc("/api/address", addressController.FindAll).Methods("GET")
	router.HandleFunc("/api/address", addressController.Create).Methods("POST")

}
