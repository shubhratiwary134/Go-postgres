package routes

import (
	"fortune-teller-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/fortunes", handlers.GetRandomFortune)

	return router
}
