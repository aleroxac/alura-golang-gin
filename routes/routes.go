package routes

import (
	"github.com/aleroxac/alura-golang-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/api/v1/healthz", controllers.Healthcheck)
	router.GET("/api/v1/skills", controllers.List)
	router.GET("/api/v1/skills/:name", controllers.GetByName)
	router.POST("/api/v1/skills", controllers.Create)
	router.PATCH("/api/v1/skills/:name", controllers.Update)
	router.DELETE("/api/v1/skills/:name", controllers.Delete)
	router.Run(":8000")
}
