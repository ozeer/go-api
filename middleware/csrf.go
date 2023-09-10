package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func CsrfProtect() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrf.Protect([]byte("251e79cd5d1a994c51fd316f7040f13d"))
		// c.Header("X-CSRF-Token", csrf.Token(c.Request))
		c.Next()
	}
}
