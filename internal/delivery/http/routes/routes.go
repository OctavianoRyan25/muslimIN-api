package routes

import (
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/handler"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authHandler *handler.UserHandler, doaHandler *handler.DoaHandler) {
	api := r.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running",
		})
	})

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// protected group
	protected := api.Group("/")
	protected.Use(middleware.JWTMiddleware())

	protected.POST("/generate-api-key", authHandler.GenerateAPIKey)

	protedtedApiKey := api.Group("/")
	protedtedApiKey.GET("/doa", doaHandler.GetAll)
	protedtedApiKey.GET("/doa/:id", doaHandler.GetById)
	protedtedApiKey.GET("/doa/random", doaHandler.GetRandom)

}
