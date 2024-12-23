package main

import (
	"net/http"

	"github.com/edwinjordan/golang_test_dot/config"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"github.com/edwinjordan/golang_test_dot/middleware"
	"github.com/edwinjordan/golang_test_dot/pkg/helpers"
	"github.com/edwinjordan/golang_test_dot/pkg/mysql"
	"github.com/edwinjordan/golang_test_dot/router"
	"github.com/rs/cors"
)

func main() {
	// fmt.Println("Hello world");
	validate := validator.New()
	db := mysql.DBConnectGorm()
	route := mux.NewRouter()

	corsOpt := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPut,
		},
		AllowedHeaders: []string{
			"*",
		},
	})

	router.CategoryRouter(db, validate, route)

	server := http.Server{
		Addr:    config.GetEnv("HOST_ADDR"),
		Handler: corsOpt.Handler(route),
	}
	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
