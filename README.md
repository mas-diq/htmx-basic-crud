# Go Notes App

A simple CRUD application for notes built with Go, Gin, HTMX, Alpine.js, DaisyUI, and MySQL.

## Tech Stack

- **Backend**: Go with Gin web framework
- **Frontend**: HTMX with Alpine.js for interactivity
- **CSS Framework**: TailwindCSS with DaisyUI components
- **Database**: MySQL

## Project Structure

```
├───cmd
│   └───server            # Main application entry point
├───internals
│   ├───configs           # Application configurations
│   ├───domain            # Domain models
│   ├───handlers          # HTTP handlers
│   ├───middlewares       # HTTP middlewares
│   ├───repositories      # Data access layer
│   ├───services          # Business logic
│   └───utils             # Utility functions
├───migrations            # Database migrations
├───tests
│   ├───integrations      # Integration tests
│   └───unit              # Unit tests
└───web
    ├───static            # Static assets (CSS, JS)
    └───templates         # HTML templates
```

## Features

- Create, read, update, and delete notes
- Responsive design with DaisyUI components
- Dark/light mode toggle
- Real-time UI updates with HTMX
- Form validation with Alpine.js
- MySQL database for data persistence

## Getting Started

### Prerequisites

- Go 1.19+
- MySQL
- Git

### Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/go-notes-app.git
cd go-notes-app
```

2. Set up the environment variables by creating a `.env` file:
```
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=notes_db
```

3. Create the database and run migrations:
```bash
# Create database
mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS notes_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# Run migration (manually or with a migration tool)
mysql -u root -p notes_db < migrations/001_create_notes_table.sql
```

4. Install dependencies and build the application:
```bash
go mod download
go build -o ./bin/server ./cmd/server
```

5. Run the application:
```bash
./bin/server
```

6. Access the application at `http://localhost:8080`

## Development

To run the application in development mode with hot reloading, you can use [Air](https://github.com/cosmtrek/air):

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run the application with Air
air
```

## Testing

Run unit tests:
```bash
go test ./tests/unit/...
```

Run integration tests:
```bash
go test ./tests/integrations/...
```

## License

This project is licensed under the MIT License.