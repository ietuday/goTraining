package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretKey"

var (
	ErrInvalidAuthHeader = errors.New("invalid authorization header")
	ErrInvalidToken      = errors.New("invalid or expired token")
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

// ExtractBearerToken converts "Bearer <token>" -> "<token>"
func ExtractBearerToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", ErrInvalidAuthHeader
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" || strings.TrimSpace(parts[1]) == "" {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimSpace(parts[1]), nil
}

// VerifyToken returns claims so you can read userId/email
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		// Ensure HS256/HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretKey), nil
	})

	if err != nil || parsedToken == nil || !parsedToken.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
