package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	Secret   []byte
	Duration time.Duration
}

func (m JWTMaker) GenerateToken(userID int32) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(int64(userID), 10),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(m.Duration)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.Secret)
}
