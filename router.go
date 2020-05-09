package main

import (
	"net/http"
)

var noteServer = &NoteServer{NewMemoryNoteStore()}
var userServer = &UserServer{NewMemoryUserStore()}

func mainRouter(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/note", applyMiddleware(noteHandler, middleware1, middleware2))
	http.HandleFunc("/user", userHandler)
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
