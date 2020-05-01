package main

import (
	"fmt"
	"net/http"
)

var noteServer = &NoteServer{NewNoteStore()}

func mainRouter(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/note", applyMiddleware(noteHandler, middleware1, middleware2))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func noteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		noteServer.getAllNotes(w)
	case http.MethodPost:
		noteServer.addNote(w, r)
	}
}
