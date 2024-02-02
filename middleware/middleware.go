package middleware

import (
	"fmt"
	"go_tutorials/database"
	"go_tutorials/models"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RemoveBearerToken(authHeader string) string {
	if strings.HasPrefix(authHeader, "Bearer ") {
		return authHeader[len("Bearer "):] // Slice off the "Bearer " prefix.
	}
	return authHeader // Return the original string if prefix is not found.
}

func Authenticated(c *gin.Context) {

	authToken := RemoveBearerToken(c.GetHeader("Authorization"))

	fmt.Println(authToken)

	if authToken == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_TOKEN")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp := claims["exp"].(float64)

		if float64(time.Now().Unix()) > exp {
			c.AbortWithStatus(http.StatusUnauthorized)

		}

		var user models.User

		database.Database.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
		//continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
