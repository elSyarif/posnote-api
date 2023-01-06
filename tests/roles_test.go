package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/elSyarif/posnote-api.git/internal/api/roles"
	"github.com/elSyarif/posnote-api.git/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var (
	URL = "http://localhost:5000/"
)

func setupTestDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/posnote_test?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sqlx.DB) *gin.Engine {
	app := gin.New()

	app.Use(gin.Recovery())
	app.Use(middleware.ErrorHandler)

	app.Group("v1")
	v1 := app.Group("v1")
	roles.NewRolesRoute(v1, db)

	return app
}

func truncate(db *sqlx.DB) {
	db.Exec("delete from roles")
}

func TestCreateRolesSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncate(db)

	requestBody := strings.NewReader(`{"name": "admin", "description" : "administrator"}`)
	request := httptest.NewRequest(http.MethodPost, URL+"v1/roles", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, response.StatusCode)
	assert.Equal(t, "success", responseBody["status"])
	assert.Nil(t, nil, responseBody["data"])
}
