package auth

import (
	"fmt"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const expectedAud = "HeroAcademy"

// AuthMiddleware is just a middleware that check JWT
func AuthMiddleware(key []byte) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{

		Key:        key,
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,

		Authorizator: func(data interface{}, c *gin.Context) bool {
			claims := jwt.ExtractClaims(c)

			sub := claims["sub"]
			if sub != nil {
				fmt.Println("sub: " + sub.(string))
			}

			aud := claims["aud"]
			if aud != nil {
				fmt.Println("aud: " + aud.(string))
				return aud == expectedAud
			}
			return false
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
