package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthServer struct {
	Users UserStore
}

func (a *AuthServer) loginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	if res, exists := a.Users.GetUserByUsername(user.Username); exists {
		fmt.Println(res)
		// TODO Generate JWT Token
	}
}
