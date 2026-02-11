package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(authHeader)
		if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			tokenString = strings.TrimSpace(tokenString[7:])
		}

		fmt.Printf("DEBUG: Token String yang diterima: [%s]\n", tokenString) // Tambahkan ini

		token, err := util.ValidateJWT(tokenString)
		if err != nil {
			fmt.Println("DEBUG: Error Parse JWT:", err.Error()) // INI PENTING: Lihat error aslinya apa
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token: " + err.Error(), // Sementara tampilkan errornya
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			c.Abort()
			return
		}

		userIDFloat := claims["user_id"].(float64)
		userID := uint(userIDFloat)

		// simpan ke context
		c.Set("user_id", userID)

		c.Next()
	}
}
