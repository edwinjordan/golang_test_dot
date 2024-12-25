package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edwinjordan/golang_test_dot.git/app/usecase/usecase_category"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/mysql"
	"github.com/edwinjordan/golang_test_dot.git/pkg/redis"
	"github.com/edwinjordan/golang_test_dot.git/repository/category_repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// func truncateCategory(db *sql.DB) {
// 	db.Exec("TRUNCATE category")
// }

func setupTestDB() *sql.DB {
	// Implement the logic to setup and return a test database connection
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_test_dot")
	if err != nil {
		panic(err)
	}
	return db
}

func setupRouter() http.Handler {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	db_redis := redis.NewRedisClient()
	router := mux.NewRouter()
	categoryRepository := category_repository.New(db, db_redis)
	categoryController := usecase_category.NewUseCase(categoryRepository, validate)
	router.HandleFunc("/api/category", categoryController.FindAll).Methods("GET")
	router.HandleFunc("/api/category", categoryController.Create).Methods("POST")
	router.HandleFunc("/api/category/{categoryId}", categoryController.Update).Methods("PUT")
	router.HandleFunc("/api/category/{categoryId}", categoryController.Delete).Methods("DELETE")

	return router

}

func TestGetCategories(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouter()

	// Insert test data
	db.Exec("INSERT INTO ms_category (category_id, category_name) VALUES (1, 'Category 1'), (2, 'Category 2')")

	req, _ := http.NewRequest("GET", "/api/category", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body, _ := io.ReadAll(rr.Body)
	var categories []entity.Category
	json.Unmarshal(body, &categories)

	//assert.Equal(t, 2, len(categories))
	if len(categories) > 0 {
		assert.Equal(t, "Category 1", categories[0].CategoryName)
	}
	if len(categories) > 1 {
		assert.Equal(t, "Category 2", categories[1].CategoryName)
	}
}

func TestCreateCategory(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouter()

	payload := `{"category_name": "Gadget"}`
	req, _ := http.NewRequest("POST", "/api/category", strings.NewReader(payload))
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
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["category_name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//	truncateCategory(db)
	router := setupRouter()

	requestBody := strings.NewReader(`{"category_name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3131/api/category", requestBody)
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

func TestUpdateCategory(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouter()

	// Insert test data
	db.Exec("INSERT INTO ms_category (category_id, category_name) VALUES ('52fe3bcda0cc77b20b7bf45a6be58878', 'Category 1')")

	payload := strings.NewReader(`{"category_name" : "Gadget"}`)
	req := httptest.NewRequest(http.MethodPut, "http://localhost:3131/api/category/52fe3bcda0cc77b20b7bf45a6be58878", payload)
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
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["category_name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	//truncateCategory(db)
	router := setupRouter()

	requestBody := strings.NewReader(`{"category_name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3131/api/category/52fe3bcda0cc77b20b7bf45a6be58878", requestBody)
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

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	//truncateCategory(db)
	defer db.Close()
	router := setupRouter()

	db.Exec("INSERT INTO ms_category (category_id, category_name) VALUES ('51fe3bcda0cc77b20b7bf45a6be58878', 'Category Example')")

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3131/api/category/51fe3bcda0cc77b20b7bf45a6be58878", nil)
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

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	//truncateCategory(db)
	defer db.Close()
	router := setupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3131/api/category/404", nil)
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
