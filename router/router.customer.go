package router

import (
	"github.com/edwinjordan/golang_test_dot.git/app/usecase/usecase_customer"
	"github.com/edwinjordan/golang_test_dot.git/repository/customer_repository"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CustomerRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router, redisClient *redis.Client) {
	customerRepository := customer_repository.New(db, redisClient)
	customerController := usecase_customer.NewUseCase(customerRepository, validate)
	router.HandleFunc("/api/customer", customerController.FindAll).Methods("GET")
	router.HandleFunc("/api/customer/{customerId}", customerController.FindById).Methods("GET")
	router.HandleFunc("/api/customer", customerController.Create).Methods("POST")
	router.HandleFunc("/api/customer/{customerId}", customerController.Update).Methods("PUT")
	router.HandleFunc("/api/customer/{customerId}", customerController.Delete).Methods("DELETE")
}
