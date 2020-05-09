package main

import "fmt"

type MemoryUserStore struct {
	store Users
}

type Users map[string]User

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{Users{}}
}

func (u *MemoryUserStore) InsertUser(user User) {
	u.store[user.Username] = user
	fmt.Println(u.store)
}

func (u *MemoryUserStore) GetUserByUsername(username string) User {
	return u.store[username]
}
