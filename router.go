package main

import (
	"fmt"
	"net/http"
)

var noteServer = &NoteServer{NewMemoryNoteStore()}
var userServer = &UserServer{NewMemoryUserStore()}
var authServer = &AuthServer{userServer.store}

func mainRouter(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/note", applyMiddleware(noteHandler, validateToken))
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/login", loginHandler)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userServer.getUser(w, r)
	case http.MethodPost:
		userServer.insertUser(w, r)
	}
}

func noteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		noteServer.getAllNotes(w)
	case http.MethodPost:
		noteServer.addNote(w, r)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		authServer.loginUser(w, r)
	} else {
		errHandler(w, r)
	}
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not Found")
}
