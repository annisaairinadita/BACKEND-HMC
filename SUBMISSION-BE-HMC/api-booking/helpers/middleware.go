package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeMiddleware(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		claims, exists := ctx.Get("claims")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		role, ok := claims.(jwt.MapClaims)["role"].(string)
		if !ok || role != requiredRole {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}