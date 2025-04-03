package services

import (
	"errors"

	"github.com/mas-diq/htmx-basic-crud/internals/domain"
	"github.com/mas-diq/htmx-basic-crud/internals/repositories"
)

// ErrNoteNotFound is returned when a note is not found
var ErrNoteNotFound = errors.New("note not found")

// NoteService defines the interface for note business logic
type NoteService interface {
	GetAllNotes() ([]*domain.Note, error)
	GetNoteByID(id int64) (*domain.Note, error)
	CreateNote(title, content string) (*domain.Note, error)
	UpdateNote(id int64, title, content string) (*domain.Note, error)
	DeleteNote(id int64) error
}

type noteService struct {
	repo repositories.NoteRepository
}

// NewNoteService creates a new note service
func NewNoteService(repo repositories.NoteRepository) NoteService {
	return &noteService{repo}
}

// GetAllNotes returns all notes
func (s *noteService) GetAllNotes() ([]*domain.Note, error) {
	return s.repo.FindAll()
}

// GetNoteByID returns a note by ID
func (s *noteService) GetNoteByID(id int64) (*domain.Note, error) {
	note, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if note == nil {
		return nil, ErrNoteNotFound
	}
	return note, nil
}

// CreateNote creates a new note
func (s *noteService) CreateNote(title, content string) (*domain.Note, error) {
	note := domain.NewNote(title, content)
	id, err := s.repo.Create(note)
	if err != nil {
		return nil, err
	}
	note.ID = id
	return note, nil
}

// UpdateNote updates an existing note
func (s *noteService) UpdateNote(id int64, title, content string) (*domain.Note, error) {
	note, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if note == nil {
		return nil, ErrNoteNotFound
	}

	note.Title = title
	note.Content = content

	if err := s.repo.Update(note); err != nil {
		return nil, err
	}
	return note, nil
}

// DeleteNote deletes a note
func (s *noteService) DeleteNote(id int64) error {
	note, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if note == nil {
		return ErrNoteNotFound
	}
	return s.repo.Delete(id)
}
