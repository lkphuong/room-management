package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JwtPayload struct {
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	StoreIDs []string `json:"store_ids"`
}

func GetInfoUser(c *gin.Context) JwtPayload {
	user, exists := c.Get("user")
	if !exists {
		return JwtPayload{}
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		fmt.Println("user is not of type jwt.MapClaims")
		return JwtPayload{}
	}

	storeIDs := make([]string, len(claims["store_ids"].([]interface{})))
	for i, v := range claims["store_ids"].([]interface{}) {
		storeIDs[i] = fmt.Sprintf("%v", v)
	}

	userPayload := JwtPayload{
		Code:     claims["code"].(string),
		Name:     claims["name"].(string),
		StoreIDs: storeIDs,
	}

	return userPayload
}

func ConvertSliceToString(slice []string) string {
	quotedIDs := make([]string, len(slice))
	for i, id := range slice {
		quotedIDs[i] = fmt.Sprintf("'%s'", id)
	}
	return strings.Join(quotedIDs, ",")
}
