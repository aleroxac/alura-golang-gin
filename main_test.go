package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aleroxac/alura-golang-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestHealthcheckRoute(t *testing.T) {
	route := SetupTestRoutes()
	route.GET("/api/v1/healthz", controllers.Healthcheck)

	request, _ := http.NewRequest("GET", "/api/v1/healthz", nil)
	response := httptest.NewRecorder()
	route.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Should be the same")

	expected := `{"status":"OK"}`
	returned, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, expected, string(returned))
}
