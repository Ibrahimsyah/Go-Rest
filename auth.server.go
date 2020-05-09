package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthServer struct {
	Users UserStore
}

type AuthResponse struct {
	Data struct {
		Key string `json:"accessToken"`
	} `json:"data"`
	Error string `json:"error"`
}

func getSecretKey() []byte {
	return []byte("Hayooiniadalahsecretkeybuatjwtbos")
}
func (a *AuthServer) loginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	var response AuthResponse
	if res, exists := a.Users.GetUserByUsername(user.Username); exists {
		if res.Password == user.Password {
			expired := time.Now().Add(time.Hour * 24).Unix()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"User":      res.Username,
				"expiredAt": expired,
			})
			if tokenString, err := token.SignedString(getSecretKey()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				response.Error = err.Error()
			} else {
				response.Data.Key = tokenString
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
			response.Error = "Wrong Password"
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		response.Error = "User Not Found"

	}
	json.NewEncoder(w).Encode(response)
}
