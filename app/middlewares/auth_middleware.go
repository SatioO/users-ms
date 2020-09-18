package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"github.com/satioO/users/app/domain"

	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware ...
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := TokenValid(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// ExtractToken ...
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

// VerifyToken ...
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenStr := ExtractToken(r)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected string method %v", token.Header["Alg"])
		}
		return []byte(viper.GetString("security.accesssecret")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid ...
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

// ExtractMetadata ...
func ExtractMetadata(r *http.Request) (*domain.AccessDetails, error) {
	token, err := VerifyToken(r)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)

		if !ok {
			return nil, err
		}

		userID := claims["user_id"].(string)

		return &domain.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}

	return nil, err
}
