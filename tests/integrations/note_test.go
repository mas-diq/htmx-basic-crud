package integrations

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mas-diq/htmx-basic-crud/internals/handlers"
	"github.com/mas-diq/htmx-basic-crud/internals/repositories"
	"github.com/mas-diq/htmx-basic-crud/internals/services"
)

// Setup test database connection
// Note: This requires a test database to be set up
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/notes_test?parseTime=true")
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to ping test database: %v", err)
	}

	// Clear test database
	_, err = db.Exec("TRUNCATE TABLE notes")
	if err != nil {
		t.Fatalf("Failed to truncate notes table: %v", err)
	}

	return db
}

// Setup Gin router for testing
func setupRouter(db *sql.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.LoadHTMLGlob("../../web/templates/**/*")

	noteRepo := repositories.NewNoteRepository(db)
	noteService := services.NewNoteService(noteRepo)
	noteHandler := handlers.NewNoteHandler(noteService)

	r.GET("/notes", noteHandler.Index)
	r.POST("/notes", noteHandler.Create)
	r.GET("/notes/:id", noteHandler.Show)
	r.PUT("/notes/:id", noteHandler.Update)
	r.DELETE("/notes/:id", noteHandler.Delete)

	return r
}

func TestCreateNoteIntegration(t *testing.T) {
	// Skip if integration tests should not run
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db := setupTestDB(t)
	defer db.Close()
	router := setupRouter(db)

	// Create a new note
	form := url.Values{}
	form.Add("title", "Integration Test Note")
	form.Add("content", "This is a test note for integration testing.")

	req, _ := http.NewRequest("POST", "/notes", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// For a regular form submission, not HTMX
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusSeeOther {
		t.Errorf("Expected status %d, got %d", http.StatusSeeOther, w.Code)
	}

	// Verify the note was created by fetching all notes
	req, _ = http.NewRequest("GET", "/notes", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check if the response contains the note title
	if !strings.Contains(w.Body.String(), "Integration Test Note") {
		t.Errorf("Expected response to contain the note title")
	}
}

func TestGetNoteIntegration(t *testing.T) {
	// Skip if integration tests should not run
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db := setupTestDB(t)
	defer db.Close()
	router := setupRouter(db)

	// First create a note
	_, err := db.Exec("INSERT INTO notes (title, content) VALUES (?, ?)", "Test Note", "Test Content")
	if err != nil {
		t.Fatalf("Failed to create test note: %v", err)
	}

	// Get the note's ID
	var id int64
	err = db.QueryRow("SELECT id FROM notes WHERE title = ?", "Test Note").Scan(&id)
	if err != nil {
		t.Fatalf("Failed to get note ID: %v", err)
	}

	// Request the note
	req, _ := http.NewRequest("GET", "/notes/"+string(id), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check if the response contains the note title
	if !strings.Contains(w.Body.String(), "Test Note") {
		t.Errorf("Expected response to contain the note title")
	}
}

// Note: For a complete test suite, add tests for update and delete operations
// These are omitted here for brevity
