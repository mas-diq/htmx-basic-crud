-- Create notes table
CREATE TABLE IF NOT EXISTS notes (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert some sample data
INSERT INTO notes (title, content) VALUES 
    ('Welcome to Notes App', 'This is a simple note-taking application built with Go, Gin, HTMX, Alpine.js, and DaisyUI.'),
    ('Getting Started', 'You can create, read, update, and delete notes using this application.'),
    ('HTMX', 'HTMX allows you to access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes.'),
    ('Alpine.js', 'Alpine.js offers you the reactive and declarative nature of big frameworks like Vue or React at a much lower cost.');