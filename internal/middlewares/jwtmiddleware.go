package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	errorForbidden = errors.New("permission denied")
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errorForbidden.Error()})
		return
		
		// c.Next()
	}
}