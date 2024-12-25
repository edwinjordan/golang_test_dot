package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edwinjordan/golang_test_dot.git/app/usecase/usecase_customer_address"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/mysql"
	"github.com/edwinjordan/golang_test_dot.git/repository/customer_address_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestDBCustAdd() *sql.DB {
	// Implement the logic to setup and return a test database connection
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_test_dot")
	if err != nil {
		panic(err)
	}
	return db
}

func setupRouterCustAdd() http.Handler {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	router := mux.NewRouter()
	addressRepository := customer_address_repository.New(db)
	addressController := usecase_customer_address.NewUseCase(addressRepository, validate)
	router.HandleFunc("/api/address", addressController.FindAll).Methods("GET")
	router.HandleFunc("/api/address", addressController.Create).Methods("POST")

	return router
}

func TestGetCustomerAddress(t *testing.T) {
	db := setupTestDBCustAdd()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterCustAdd()

	// Insert test data
	db.Exec("INSERT INTO ms_category (address_id, address_customer_id, address_text,address_name, address_postal_code) VALUES (1, '9533d4b4c417ebab4f09e54ab9c96857','Jl Raya Bogor','Rumah Sendiri', '64181')")

	req, _ := http.NewRequest("GET", "/api/address", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body, _ := io.ReadAll(rr.Body)
	var categories []entity.CustomerAddress
	json.Unmarshal(body, &categories)

	//assert.Equal(t, 2, len(categories))
	if len(categories) > 0 {
		assert.Equal(t, "9533d4b4c417ebab4f09e54ab9c96857", categories[0].AddressCustomerId)
	}
}

func TestCreateCustomerAddress(t *testing.T) {
	db := setupTestDBCustAdd()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouterCustAdd()

	payload := `{"address_customer_id": "9533d4b4c417ebab4f09e54ab9c96857" , "address_text" : "Jl Raya Kediri","address_name":"Rumah sendiri", "customer_email":"jono@gmail.com","address_postal_code":"64181"}`
	req, _ := http.NewRequest("POST", "/api/address", strings.NewReader(payload))
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
	assert.Equal(t, "9533d4b4c417ebab4f09e54ab9c96857", responseBody["data"].(map[string]interface{})["address_customer_id"])
}

func TestCreateCustomerAddressFailed(t *testing.T) {
	db := setupTestDBCustAdd()
	defer db.Close()
	//	truncateCategory(db)
	router := setupRouterCustAdd()

	requestBody := strings.NewReader(`{"address_customer_id": "" , "address_text" : "Jl Raya Kediri","address_name":"Rumah sendiri", "customer_email":"jono@gmail.com","address_postal_code":"64181"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3131/api/address", requestBody)
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
