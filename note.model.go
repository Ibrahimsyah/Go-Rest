package main

//Note object type
type Note struct {
	Title string `json:"title"`
	Note  string `json:"note"`
}

type MemoryNoteStore struct {
	store Notes
}

type Notes []Note

func NewNoteStore() *MemoryNoteStore {
	return &MemoryNoteStore{}
}

func (s *MemoryNoteStore) InsertNote(note Note) {
	s.store = append(s.store, note)
}

func (s *MemoryNoteStore) GetAllNotes() Notes {
	return s.store
}
