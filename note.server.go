package main

import (
	"encoding/json"
	"net/http"
)

type Note struct {
	Title string `json:"title"`
	Note  string `json:"note"`
}

type NoteStore interface {
	InsertNote(note Note)
	GetAllNotes() Notes
}

type NoteServer struct {
	store NoteStore
}

func (n *NoteServer) addNote(w http.ResponseWriter, r *http.Request) {
	var t Note
	json.NewDecoder(r.Body).Decode(&t)
	n.store.InsertNote(t)
}

func (n *NoteServer) getAllNotes(w http.ResponseWriter) {
	Notes := n.store.GetAllNotes()
	json.NewEncoder(w).Encode(Notes)
}
