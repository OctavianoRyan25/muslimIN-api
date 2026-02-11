package routes

import (
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authHandler *handler.UserHandler) {
	api := r.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running",
		})
	})

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

}
