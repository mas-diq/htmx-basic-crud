package repositories

import (
	"database/sql"
	"time"

	"github.com/mas-diq/htmx-basic-crud/internals/domain"
)

// NoteRepository defines the interface for note database operations
type NoteRepository interface {
	FindAll() ([]*domain.Note, error)
	FindByID(id int64) (*domain.Note, error)
	Create(note *domain.Note) (int64, error)
	Update(note *domain.Note) error
	Delete(id int64) error
}

type noteRepository struct {
	db *sql.DB
}

// NewNoteRepository creates a new note repository
func NewNoteRepository(db *sql.DB) NoteRepository {
	return &noteRepository{db}
}

// FindAll returns all notes
func (r *noteRepository) FindAll() ([]*domain.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM notes ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*domain.Note
	for rows.Next() {
		note := &domain.Note{}
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

// FindByID returns a note by ID
func (r *noteRepository) FindByID(id int64) (*domain.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM notes WHERE id = ?`
	row := r.db.QueryRow(query, id)

	note := &domain.Note{}
	err := row.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return note, nil
}

// Create creates a new note
func (r *noteRepository) Create(note *domain.Note) (int64, error) {
	query := `INSERT INTO notes (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, note.Title, note.Content, note.CreatedAt, note.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update updates an existing note
func (r *noteRepository) Update(note *domain.Note) error {
	note.UpdatedAt = time.Now()
	query := `UPDATE notes SET title = ?, content = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, note.Title, note.Content, note.UpdatedAt, note.ID)
	return err
}

// Delete deletes a note
func (r *noteRepository) Delete(id int64) error {
	query := `DELETE FROM notes WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
