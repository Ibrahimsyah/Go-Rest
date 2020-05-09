package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserStore interface {
	GetUserByUsername(username string) (User, bool)
	InsertUser(user User)
}

type UserServer struct {
	store UserStore
}

func (u *UserServer) getUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	res, _ := u.store.GetUserByUsername(user.Username)
	json.NewEncoder(w).Encode(res)
}

func (u *UserServer) insertUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	u.store.InsertUser(user)
}
