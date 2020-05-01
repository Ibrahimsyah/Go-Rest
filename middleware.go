package main

import (
	"encoding/json"
	"net/http"
	"strings"
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

func middleware1(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "" {
			if token := strings.SplitN(auth, " ", 2); token[1] != "" {
				h.ServeHTTP(w, r)
			}
		} else {
			error := Error{"Gaoleh Masuk Bos", "Endi token 1 mu?"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(error)
			return
		}
	})
}

func middleware2(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Auth"); auth != "" {
			if token := strings.SplitN(auth, " ", 2); token[1] != "" {
				h.ServeHTTP(w, r)
			}
		} else {
			error := Error{"Gaoleh Masuk Bos", "Endi token 2 mu?"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(error)
			return
		}
	})
}


