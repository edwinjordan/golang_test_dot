package router

import (
	"github.com/edwinjordan/golang_test_dot.git/app/usecase/usecase_order"
	"github.com/edwinjordan/golang_test_dot.git/repository/order_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func OrderRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	orderRepository := order_repository.NewOrder(db)
	orderDetailRepository := order_repository.NewOrderDetail(db)

	orderController := usecase_order.NewUseCase(orderRepository, orderDetailRepository, validate)
	router.HandleFunc("/api/order", orderController.Create).Methods("POST")

}
