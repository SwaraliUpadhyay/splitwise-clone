package routes

import (
	"splitwise-clone/handlers"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.GET("/health", handlers.HealthCheck)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", handlers.Signup)
		auth.POST("/login", handlers.Login)
	}
}
