package BearerToken

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v4"
	"log/slog"
	"net/http"
	"strings"
)

func check_jwt_token(log *slog.Logger, next http.Handler, secretKey []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			log.Error("Anauthorized acess attepmt !!! from:", r.RemoteAddr)
			render.JSON(w, r, "Anauthorized")
			return
		}

		if !(strings.Contains(header, "Bearer")) {
			log.Error("Invalid token format in ", middleware.GetReqID(r.Context()))
			render.JSON(w, r, "Invalid token format")
			return
		}

		tokenstr := strings.Replace(header, "Bearer", "", 1)

		token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("некорректный signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if token.Valid {
			next.ServeHTTP(w, r)
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			log.Error("Invalid token format in ", middleware.GetReqID(r.Context()))
			render.JSON(w, r, "Invalid format")
			return
		}
	})
}
