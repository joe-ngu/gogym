package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joe-ngu/gogym/store"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func CreateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
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
		log.Println("calling jwt middleware")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			PermissionDenied()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := getID(r)
		if err != nil {
			PermissionDenied()
			return
		}

		user, err := s.GetUserByID(userID)
		if err != nil {
			PermissionDenied()
			return
		}

		claims, err := validateJWT(tokenString)
		if user.UserName != claims.Username {
			PermissionDenied()
			return
		}
		if err != nil {
			log.Println("Invalid token")
			PermissionDenied()
			return
		}
		h(w, r)
	}
}

func getID(r *http.Request) (uuid.UUID, error) {
	idStr := r.PathValue("id")
	if idStr == "" {
		return uuid.Nil, fmt.Errorf("id not found")
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}
