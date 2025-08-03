package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-analyser-backend/handlers"
	"url-analyser-backend/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddURL(t *testing.T) {
	// Create a mock database
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Mock the database behavior
	mock.ExpectExec("INSERT INTO urls").WithArgs("https://example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Create a Gin router and attach the handler
	router := gin.Default()
	router.POST("/urls", handlers.AddURL(db))

	// Create a test request
	body := `{"url": "https://example.com"}`
	req, _ := http.NewRequest("POST", "/urls", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "URL submitted", response["message"])
	assert.Equal(t, float64(1), response["id"]) // JSON unmarshals numbers as float64
}

func TestListURLs(t *testing.T) {
	// Create a mock database
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Mock the database behavior
	rows := sqlmock.NewRows([]string{"id", "url", "status", "created_at", "updated_at"}).
		AddRow(1, "https://example.com", "queued", "2025-07-15T12:00:00Z", "2025-07-15T12:00:00Z")
	mock.ExpectQuery("SELECT id, url, status, created_at, updated_at FROM urls").WillReturnRows(rows)

	// Create a Gin router and attach the handler
	router := gin.Default()
	router.GET("/urls", handlers.ListURLs(db))

	// Create a test request
	req, _ := http.NewRequest("GET", "/urls", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.URL
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Len(t, response, 1)
	assert.Equal(t, "https://example.com", response[0].URL)
	assert.Equal(t, "queued", response[0].Status)
}
