package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edwinjordan/golang_test_dot.git/app/usecase/usecase_customer"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/mysql"
	"github.com/edwinjordan/golang_test_dot.git/pkg/redis"
	"github.com/edwinjordan/golang_test_dot.git/repository/customer_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func setupTestCustDB() *sql.DB {
	// Implement the logic to setup and return a test database connection
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_test_dot")
	if err != nil {
		panic(err)
	}
	return db
}

func setupRouterCust() http.Handler {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	db_redis := redis.NewRedisClient()
	router := mux.NewRouter()
	customerRepository := customer_repository.New(db, db_redis)
	customerController := usecase_customer.NewUseCase(customerRepository, validate)
	router.HandleFunc("/api/customer", customerController.FindAll).Methods("GET")
	router.HandleFunc("/api/customer", customerController.Create).Methods("POST")
	router.HandleFunc("/api/customer/{customerId}", customerController.Update).Methods("PUT")
	router.HandleFunc("/api/customer/{customerId}", customerController.Delete).Methods("DELETE")

	return router

}

func TestGetCustomer(t *testing.T) {
	db := setupTestCustDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterCust()

	// Insert test data
	db.Exec("INSERT INTO tb_customer (customer_id, customer_name, customer_gender, customer_phonenumber,customer_email, customer_password) VALUES (1, 'Jordan','L','085617251423','jordan@gmail.com','Aero1996'), (2, 'Laksono','L','085617251423','laksono@gmail.com','Aero1996')")

	req, _ := http.NewRequest("GET", "/api/customer", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body, _ := io.ReadAll(rr.Body)
	var customers []entity.Customer
	json.Unmarshal(body, &customers)

	//assert.Equal(t, 2, len(categories))
	if len(customers) > 0 {
		assert.Equal(t, "Jordan", customers[0].CustomerName)
	}
	if len(customers) > 1 {
		assert.Equal(t, "Laksono", customers[1].CustomerName)
	}
}

func TestCreateCustomer(t *testing.T) {
	db := setupTestCustDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterCust()

	payload := `{"customer_name": "Jono" , "customer_gender" : "L","customer_phonenumber":"086786543222", "customer_email":"jono@gmail.com","customer_password":"Aero1996"}`
	req, _ := http.NewRequest("POST", "/api/customer", strings.NewReader(payload))
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
	assert.Equal(t, "Jono", responseBody["data"].(map[string]interface{})["customer_name"])
}

func TestCreateCustomerFailed(t *testing.T) {
	db := setupTestCustDB()
	defer db.Close()
	//	truncateCategory(db)
	router := setupRouterCust()

	requestBody := strings.NewReader(`{"customer_name": "" , "customer_gender" : "L","customer_phonenumber":"086786543255", "customer_email":"joni@gmail.com","customer_password":"Aero1996"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3131/api/customer", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	//assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Data tidak ditemukan", responseBody["status"])
}

func TestUpdateCustomer(t *testing.T) {
	db := setupTestCustDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterCust()

	// Insert test data
	db.Exec("INSERT INTO tb_customer (customer_id, customer_name, customer_gender, customer_phonenumber,customer_email, customer_password) VALUES (1, 'Jordan','L','085617251423','jordan@gmail.com','$2a$10$15myar1kDjnpQJRs7dTopubBo25IRzXZwkN/G/f.Q0w1iyrS3TmUK')")

	payload := strings.NewReader(`{"customer_name" : "Gadget", "customer_old_password" : "Aero1996", "customer_new_password" : "Aero1996", "customer_gender" : "L"}`)
	req := httptest.NewRequest(http.MethodPut, "http://localhost:3131/api/customer/1", payload)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	response := rr.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Berhasil mengubah data", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["customer_name"])
}

func TestUpdateCustomerFailed(t *testing.T) {
	db := setupTestCustDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterCust()

	requestBody := strings.NewReader(`{"customer_name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3131/api/customer/1", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	//assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Data tidak ditemukan", responseBody["status"])
}

func TestDeleteCustomerSuccess(t *testing.T) {
	db := setupTestCustDB()
	//truncateCategory(db)
	defer db.Close()
	router := setupRouterCust()

	db.Exec("INSERT INTO tb_customer (customer_id, customer_name, customer_gender, customer_phonenumber,customer_email, customer_password) VALUES (1, 'Jordan','L','085617251423','jordan@gmail.com','$2a$10$15myar1kDjnpQJRs7dTopubBo25IRzXZwkN/G/f.Q0w1iyrS3TmUK')")

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3131/api/customer/1", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Berhasil menghapus data", responseBody["status"])
}

func TestDeleteCustomerFailed(t *testing.T) {
	db := setupTestCustDB()
	//truncateCategory(db)
	defer db.Close()
	router := setupRouterCust()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3131/api/customer/404", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	//assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Data tidak ditemukan", responseBody["status"])
}
