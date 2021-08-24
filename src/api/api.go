package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiMetricsResponse struct {
	Metrics map[string]int32 `json:"metrics"`
}

type ApiHealthResponse struct {
	Status string `json:"status"`
}

var apiMetricsResponse ApiMetricsResponse

func getMetrics(c *gin.Context) {
	response := ApiMetricsResponse{
		Metrics: map[string]int32{"co2_emissions": 100},
	}
	c.IndentedJSON(http.StatusOK, response)
}

func getHealth(c *gin.Context) {
	response := ApiHealthResponse{
		Status: "ok",
	}
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.GET("/metrics", getMetrics)
	router.GET("/health", getHealth)
	router.Run("localhost:10000")
}
