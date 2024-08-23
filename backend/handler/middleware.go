package handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/joe-ngu/gogym/store"
)

type contextKey string

const UserIDKey = contextKey("userID")

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID   uuid.UUID
	Username string
	jwt.RegisteredClaims
}

func CreateJWT(userID uuid.UUID, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func validateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func JWTAuthMiddleware(h http.HandlerFunc, s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			PermissionDenied(w)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := validateJWT(tokenString)
		if err != nil {
			log.Println("invalid token")
			PermissionDenied(w)
			return
		}

		user, err := s.GetUserByID(claims.UserID)
		if err != nil {
			PermissionDenied(w)
			return
		}

		if user.UserName != claims.Username {
			PermissionDenied(w)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, user.ID)

		h(w, r.WithContext(ctx))
	}
}
