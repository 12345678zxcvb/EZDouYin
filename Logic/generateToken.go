package Logic

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateToken(username, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})
	tokenString, err := token.SignedString([]byte("zhang"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func AuthMiddleware(secretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Query("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse claims"})
			c.Abort()
			return
		}

		//fmt.Println("User:", claims["username"])

		c.Next()
	}
}
func AuthMiddleware2(secretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.PostForm("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse claims"})
			c.Abort()
			return
		}

		fmt.Println("User:", claims["username"])

		c.Next()
	}
}
func ExtractUsernameFromToken(tokenString string, secretKey []byte) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("无效的 token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("解析声明失败")
	}

	username := claims["username"].(string)
	return username, nil
}
