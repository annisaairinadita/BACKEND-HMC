package helpers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateTokenJWT(userID int, role string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := []byte(os.Getenv("JWT_KEY")) 

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,	
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
	}

	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokenJWT.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateTokenJWT(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	tokenJWT, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := tokenJWT.Claims.(jwt.MapClaims); ok && tokenJWT.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := godotenv.Load()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading .env file"})
			ctx.Abort()
			return
		}

		secretKey := []byte(os.Getenv("JWT_KEY")) 

		headerToken := ctx.GetHeader("Authorization")
		if headerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		if !strings.HasPrefix(headerToken, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})

			return
		}

		tokenString := strings.TrimPrefix(headerToken, "Bearer ")

		claims, err := ValidateTokenJWT(tokenString, secretKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}

func GetUserClaims(c *gin.Context) (int, string, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return 0, "", errors.New("Unauthorized")
	}

	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("Unauthorized")
	}

	userIDFloat, ok := jwtClaims["user_id"].(float64)
	if !ok {
		return 0, "", errors.New("Unauthorized")
	}

	role, ok := jwtClaims["role"].(string)
	if !ok {
		return 0, "", errors.New("Unauthorized")
	}

	userID := int(userIDFloat)
	return userID, role, nil
}