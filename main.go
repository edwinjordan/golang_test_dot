package main

import (
	"net/http"

	"github.com/edwinjordan/golang_test_dot.git/config"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	//"github.com/edwinjordan/golang_test_dot/middleware"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"github.com/edwinjordan/golang_test_dot.git/pkg/mysql"
	"github.com/edwinjordan/golang_test_dot.git/pkg/redis"
	"github.com/edwinjordan/golang_test_dot.git/router"
	"github.com/rs/cors"
)

func main() {
	// fmt.Println("Hello world");
	validate := validator.New()
	db := mysql.DBConnectGorm()
	db_redis := redis.NewRedisClient()

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

	router.CategoryRouter(db, validate, route, db_redis)
	router.CustomerRouter(db, validate, route, db_redis)
	router.CustomerAddressRouter(db, validate, route)
	router.OrderRouter(db, validate, route)

	server := http.Server{
		Addr:    config.GetEnv("HOST_ADDR"),
		Handler: corsOpt.Handler(route),
	}
	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
