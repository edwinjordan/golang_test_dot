package router

import (
	"github.com/edwinjordan/golang_test_dot.git/app/usecase/usecase_category"
	"github.com/edwinjordan/golang_test_dot.git/repository/category_repository"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CategoryRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router, redisClient *redis.Client) {
	categoryRepository := category_repository.New(db, redisClient)
	categoryController := usecase_category.NewUseCase(categoryRepository, validate)
	router.HandleFunc("/api/category", categoryController.FindAll).Methods("GET")
	router.HandleFunc("/api/category/{categoryId}", categoryController.FindById).Methods("GET")
	router.HandleFunc("/api/category", categoryController.Create).Methods("POST")
	router.HandleFunc("/api/category/{categoryId}", categoryController.Update).Methods("PUT")
	router.HandleFunc("/api/category/{categoryId}", categoryController.Delete).Methods("DELETE")
	router.HandleFunc("/api/category/{categoryId}", categoryController.Patch).Methods("PATCH")
}
