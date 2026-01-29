package middlewares

import (
	"main/internal/exceptions"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const IDKey = "UserId"

func JWTMiddleware(secret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if strings.TrimSpace(auth) == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
			return
		}

		tokenStr := strings.TrimSpace(parts[1])

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, exceptions.ErrUnauthorized
			}
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, exceptions.ErrUnauthorized
			}
			return secret, nil
		})

		if err != nil || token == nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok || sub == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
			return
		}

		userID64, err := strconv.ParseInt(sub, 10, 32)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
			return
		}

		ctx.Set(IDKey, int32(userID64))
		ctx.Next()
	}
}
