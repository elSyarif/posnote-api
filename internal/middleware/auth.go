package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			helper.HTTPResponseError(c, http.StatusUnauthorized, "fail", "Unauthenticated", nil)
			c.Abort()
			return
		}

		if !strings.Contains(authorization, "Bearer") {
			helper.HTTPResponseError(c, http.StatusUnauthorized, "fail", "invalid token", nil)
			c.Abort()
			return
		}

		token := strings.Replace(authorization, "Bearer ", "", -1)
		_, err := helper.GetJWTData(token, os.Getenv("ACCESS_TOKEN_KEY"))
		if err != nil {
			helper.HTTPResponseError(c, http.StatusUnauthorized, "fail", err.Error(), nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
