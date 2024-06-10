package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lkphuong/room-management/internal/utils"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if !strings.HasPrefix(tokenString, "Bearer ") {
			utils.SadResp(errors.New("invalid token format"), http.StatusUnauthorized, c)
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		mySigningKey := []byte(os.Getenv("SECRET_KEY"))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Println("check here 123")
			utils.SadResp(err, http.StatusUnauthorized, c)
			c.Abort()
		}

		if !token.Valid {
			fmt.Println(err)
		}

		c.Next()
	}
}
