package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edwinjordan/golang_test_dot.git/app/usecase/usecase_order"
	"github.com/edwinjordan/golang_test_dot.git/pkg/mysql"
	"github.com/edwinjordan/golang_test_dot.git/repository/order_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestDBOrder() *sql.DB {
	// Implement the logic to setup and return a test database connection
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_test_dot")
	if err != nil {
		panic(err)
	}
	return db
}

func setupRouterOrder() http.Handler {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	router := mux.NewRouter()
	orderRepository := order_repository.NewOrder(db)
	orderDetailRepository := order_repository.NewOrderDetail(db)

	orderController := usecase_order.NewUseCase(orderRepository, orderDetailRepository, validate)
	router.HandleFunc("/api/order", orderController.Create).Methods("POST")

	return router
}

func TestCreateCustomerOrder(t *testing.T) {
	db := setupTestDBOrder()
	defer db.Close()
	router := setupRouterOrder()
	//payload := `{"category_name": "Gadget"}`

	payload := `{
	  "parent": 
	      {
	          "order_customer_id": "9533d4b4c417ebab4f09e54ab9c96857",
	          "order_total_item": 3,
	          "order_subtotal": 200000,
	          "order_discount": 0,
	          "order_total": 600000,
	          "order_notes": "ok"
			 
	      },
	   "detail" : [
										{
												"product_nama" : "Baju",
												"product_qty" : 1,
												"product_price" : 400000

										},
										{
												"product_nama" : "Celana",
												"product_qty" : 2,
												"product_price" : 400000

										}
								]

		}
	`
	req, _ := http.NewRequest("POST", "/api/order", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	response := rr.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Berhasil menambah data", responseBody["status"])
	//assert.Equal(t, "9533d4b4c417ebab4f09e54ab9c96857", responseBody["data"].(map[string]interface{})["address_customer_id"])
}
