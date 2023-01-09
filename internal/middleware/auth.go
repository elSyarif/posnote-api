package middleware

import (
	"errors"
	"fmt"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

func Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if !strings.Contains(authorizationHeader, "Bearer") {
			c.Error(errors.New("invalid token"))
			return
		}

		token := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		data, err := helper.GetJWTData(token, os.Getenv("ACCESS_TOKEN_KEY"))
		if err != nil {
			helper.HTTPResponseError(c, http.StatusUnauthorized, "fail", "Invavalid token", nil)
			c.Abort()
			return
		}

		if time.Now().Unix() > data.ExpiresAt {
			c.Error(errors.New("token expires"))
			c.Abort()
			return
		}
		fmt.Println(data.EmpId)
		c.Next()
	}
}
