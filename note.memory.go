package main

type MemoryNoteStore struct {
	store Notes
}

type Notes []Note

func NewMemoryNoteStore() *MemoryNoteStore {
	return &MemoryNoteStore{}
}

func (s *MemoryNoteStore) InsertNote(note Note) {
	s.store = append(s.store, note)
}

func (s *MemoryNoteStore) GetAllNotes() Notes {
	return s.store
}
