package helper

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"training-final/model"

	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	jwtMethod = jwt.SigningMethodHS256
	jwtKey    = []byte("Key123")
)

type claims struct {
	jwt.StandardClaims
	Id int `json:"Id"`
}

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/login") || strings.Contains(r.URL.Path, "/register") {
			next.ServeHTTP(w, r)
			return
		}

		auth := r.Header.Get("Authorization")
		if !strings.Contains(auth, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tokenStr := strings.Replace(auth, "Bearer ", "", -1)

		claim, status, _ := ValidateJWT(tokenStr)
		if claim == nil {
			w.Write([]byte(fmt.Sprintf("token invalid, status : %d", status)))
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claim)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == "user") && (password == "pass")
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

func GenerateJWT(user model.User) (string, error) {
	claims := claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(20 * time.Minute).Unix(),
		},
		Id: user.Id,
	}
	token := jwt.NewWithClaims(jwtMethod, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenStr string) (jwt.MapClaims, int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwtMethod {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return jwtKey, nil
	})

	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, http.StatusBadRequest, err
	}

	return claims, http.StatusOK, nil
}
