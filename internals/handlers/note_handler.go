package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mas-diq/htmx-basic-crud/internals/services"
	"github.com/mas-diq/htmx-basic-crud/internals/utils"
)

// NoteHandler handles HTTP requests for notes
type NoteHandler struct {
	noteService services.NoteService
}

// NewNoteHandler creates a new note handler
func NewNoteHandler(noteService services.NoteService) *NoteHandler {
	return &NoteHandler{noteService}
}

// Index renders the notes index page
func (h *NoteHandler) Index(c *gin.Context) {
	notes, err := h.noteService.GetAllNotes()
	if err != nil {
		utils.InternalServerError(c, "Failed to fetch notes")
		return
	}

	utils.HTMLResponse(c, http.StatusOK, "notes/index.html", gin.H{
		"title": "Notes",
		"notes": notes,
	})
}

// New renders the note creation form
func (h *NoteHandler) New(c *gin.Context) {
	utils.HTMLResponse(c, http.StatusOK, "notes/create.html", gin.H{
		"title": "Create Note",
	})
}

// Create handles the note creation
func (h *NoteHandler) Create(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if title == "" {
		utils.BadRequest(c, "Title is required")
		return
	}

	note, err := h.noteService.CreateNote(title, content)
	if err != nil {
		utils.InternalServerError(c, "Failed to create note")
		return
	}

	// Check if request is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		utils.HTMLResponse(c, http.StatusOK, "notes/index.html", gin.H{
			"title": "Notes",
			"notes": []interface{}{note},
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/notes")
}

// Show renders a specific note
func (h *NoteHandler) Show(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "Invalid note ID")
		return
	}

	note, err := h.noteService.GetNoteByID(id)
	if err != nil {
		if err == services.ErrNoteNotFound {
			utils.NotFound(c)
		} else {
			utils.InternalServerError(c, "Failed to fetch note")
		}
		return
	}

	utils.HTMLResponse(c, http.StatusOK, "notes/show.html", gin.H{
		"title": note.Title,
		"note":  note,
	})
}

// Edit renders the note edit form
func (h *NoteHandler) Edit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "Invalid note ID")
		return
	}

	note, err := h.noteService.GetNoteByID(id)
	if err != nil {
		if err == services.ErrNoteNotFound {
			utils.NotFound(c)
		} else {
			utils.InternalServerError(c, "Failed to fetch note")
		}
		return
	}

	utils.HTMLResponse(c, http.StatusOK, "notes/edit.html", gin.H{
		"title": "Edit " + note.Title,
		"note":  note,
	})
}

// Update handles the note update
func (h *NoteHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "Invalid note ID")
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")

	if title == "" {
		utils.BadRequest(c, "Title is required")
		return
	}

	note, err := h.noteService.UpdateNote(id, title, content)
	if err != nil {
		if err == services.ErrNoteNotFound {
			utils.NotFound(c)
		} else {
			utils.InternalServerError(c, "Failed to update note")
		}
		return
	}

	// Check if request is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		utils.HTMLResponse(c, http.StatusOK, "notes/show.html", gin.H{
			"title": note.Title,
			"note":  note,
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/notes/"+strconv.FormatInt(note.ID, 10))
}

// Delete handles the note deletion
func (h *NoteHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequest(c, "Invalid note ID")
		return
	}

	err = h.noteService.DeleteNote(id)
	if err != nil {
		if err == services.ErrNoteNotFound {
			utils.NotFound(c)
		} else {
			utils.InternalServerError(c, "Failed to delete note")
		}
		return
	}

	// Check if request is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		c.Status(http.StatusOK)
		return
	}

	c.Redirect(http.StatusSeeOther, "/notes")
}
