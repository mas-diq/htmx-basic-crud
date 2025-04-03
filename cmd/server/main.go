package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/joho/godotenv"
	"github.com/mas-diq/htmx-basic-crud/internals/configs"
	"github.com/mas-diq/htmx-basic-crud/internals/handlers"
	"github.com/mas-diq/htmx-basic-crud/internals/repositories"
	"github.com/mas-diq/htmx-basic-crud/internals/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize database
	db, err := configs.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create router
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("../../web/templates/**/*")

	// Initialize repositories
	noteRepo := repositories.NewNoteRepository(db)

	// Initialize services
	noteService := services.NewNoteService(noteRepo)

	// Initialize handlers
	noteHandler := handlers.NewNoteHandler(noteService)

	// Register routes
	r.GET("/", noteHandler.Index)
	r.GET("/notes", noteHandler.Index)
	r.GET("/notes/new", noteHandler.New)
	r.POST("/notes", noteHandler.Create)
	r.GET("/notes/:id", noteHandler.Show)
	r.GET("/notes/:id/edit", noteHandler.Edit)
	r.PUT("/notes/:id", noteHandler.Update)
	r.DELETE("/notes/:id", noteHandler.Delete)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
