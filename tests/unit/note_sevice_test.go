package unit

import (
	"testing"
	"time"

	"github.com/mas-diq/htmx-basic-crud/internals/domain"
	"github.com/mas-diq/htmx-basic-crud/internals/services"
)

// Mock repository implementation for testing
type mockNoteRepository struct {
	notes  map[int64]*domain.Note
	nextID int64
}

func newMockRepository() *mockNoteRepository {
	return &mockNoteRepository{
		notes:  make(map[int64]*domain.Note),
		nextID: 1,
	}
}

func (m *mockNoteRepository) FindAll() ([]*domain.Note, error) {
	var notes []*domain.Note
	for _, note := range m.notes {
		notes = append(notes, note)
	}
	return notes, nil
}

func (m *mockNoteRepository) FindByID(id int64) (*domain.Note, error) {
	note, exists := m.notes[id]
	if !exists {
		return nil, nil
	}
	return note, nil
}

func (m *mockNoteRepository) Create(note *domain.Note) (int64, error) {
	id := m.nextID
	m.nextID++
	note.ID = id
	m.notes[id] = note
	return id, nil
}

func (m *mockNoteRepository) Update(note *domain.Note) error {
	if _, exists := m.notes[note.ID]; !exists {
		return nil
	}
	m.notes[note.ID] = note
	return nil
}

func (m *mockNoteRepository) Delete(id int64) error {
	delete(m.notes, id)
	return nil
}

func TestCreateNote(t *testing.T) {
	repo := newMockRepository()
	service := services.NewNoteService(repo)

	title := "Test Note"
	content := "This is a test note"

	note, err := service.CreateNote(title, content)
	if err != nil {
		t.Fatalf("Error creating note: %v", err)
	}

	if note.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", note.ID)
	}

	if note.Title != title {
		t.Errorf("Expected title to be %q, got %q", title, note.Title)
	}

	if note.Content != content {
		t.Errorf("Expected content to be %q, got %q", content, note.Content)
	}

	// Check that CreatedAt and UpdatedAt are set
	if note.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}

	if note.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should not be zero")
	}
}

func TestGetNoteByID(t *testing.T) {
	repo := newMockRepository()
	service := services.NewNoteService(repo)

	// Create a note first
	now := time.Now()
	savedNote := &domain.Note{
		ID:        1,
		Title:     "Test Note",
		Content:   "This is a test note",
		CreatedAt: now,
		UpdatedAt: now,
	}
	repo.notes[1] = savedNote

	// Get the note
	note, err := service.GetNoteByID(1)
	if err != nil {
		t.Fatalf("Error getting note: %v", err)
	}

	if note.ID != savedNote.ID {
		t.Errorf("Expected ID to be %d, got %d", savedNote.ID, note.ID)
	}

	if note.Title != savedNote.Title {
		t.Errorf("Expected title to be %q, got %q", savedNote.Title, note.Title)
	}

	// Test non-existent note
	_, err = service.GetNoteByID(999)
	if err != services.ErrNoteNotFound {
		t.Errorf("Expected ErrNoteNotFound, got %v", err)
	}
}

func TestUpdateNote(t *testing.T) {
	repo := newMockRepository()
	service := services.NewNoteService(repo)

	// Create a note first
	now := time.Now()
	savedNote := &domain.Note{
		ID:        1,
		Title:     "Test Note",
		Content:   "This is a test note",
		CreatedAt: now,
		UpdatedAt: now,
	}
	repo.notes[1] = savedNote

	// Update the note
	updatedTitle := "Updated Title"
	updatedContent := "Updated content"

	note, err := service.UpdateNote(1, updatedTitle, updatedContent)
	if err != nil {
		t.Fatalf("Error updating note: %v", err)
	}

	if note.Title != updatedTitle {
		t.Errorf("Expected title to be %q, got %q", updatedTitle, note.Title)
	}

	if note.Content != updatedContent {
		t.Errorf("Expected content to be %q, got %q", updatedContent, note.Content)
	}

	// Test non-existent note
	_, err = service.UpdateNote(999, "Title", "Content")
	if err != services.ErrNoteNotFound {
		t.Errorf("Expected ErrNoteNotFound, got %v", err)
	}
}

func TestDeleteNote(t *testing.T) {
	repo := newMockRepository()
	service := services.NewNoteService(repo)

	// Create a note first
	now := time.Now()
	savedNote := &domain.Note{
		ID:        1,
		Title:     "Test Note",
		Content:   "This is a test note",
		CreatedAt: now,
		UpdatedAt: now,
	}
	repo.notes[1] = savedNote

	// Delete the note
	err := service.DeleteNote(1)
	if err != nil {
		t.Fatalf("Error deleting note: %v", err)
	}

	// Verify the note is gone
	_, err = service.GetNoteByID(1)
	if err != services.ErrNoteNotFound {
		t.Errorf("Expected ErrNoteNotFound, got %v", err)
	}

	// Test deleting non-existent note
	err = service.DeleteNote(999)
	if err != services.ErrNoteNotFound {
		t.Errorf("Expected ErrNoteNotFound, got %v", err)
	}
}

func TestGetAllNotes(t *testing.T) {
	repo := newMockRepository()
	service := services.NewNoteService(repo)

	// Create some notes
	now := time.Now()
	repo.notes[1] = &domain.Note{ID: 1, Title: "Note 1", Content: "Content 1", CreatedAt: now, UpdatedAt: now}
	repo.notes[2] = &domain.Note{ID: 2, Title: "Note 2", Content: "Content 2", CreatedAt: now, UpdatedAt: now}
	repo.notes[3] = &domain.Note{ID: 3, Title: "Note 3", Content: "Content 3", CreatedAt: now, UpdatedAt: now}

	// Get all notes
	notes, err := service.GetAllNotes()
	if err != nil {
		t.Fatalf("Error getting all notes: %v", err)
	}

	if len(notes) != 3 {
		t.Errorf("Expected 3 notes, got %d", len(notes))
	}
}
