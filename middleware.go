package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Error struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

func applyMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}
	current := h
	for i := len(m) - 1; i >= 0; i-- {
		current = m[i](current)
	}
	return current
}

func validateToken(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var er Error
		if auth := r.Header.Get("Authorization"); auth != "" {
			if authorizationKey := strings.SplitN(auth, " ", 2); authorizationKey[1] != "" {
				authToken := authorizationKey[1]
				token, _ := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return getSecretKey(), nil
				})
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					user := claims["user"]
					ctx := context.WithValue(r.Context(), "user", user)
					r.WithContext(ctx)
					h.ServeHTTP(w, r)
					return
				} else {
					er.Description = "Tokenmu gak valid"
				}
			}
		} else {
			er.Description = "Endi Tokenmu?"
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(er)
	})
}
