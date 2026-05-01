CREATE TABLE IF NOT EXISTS books (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(32) NOT NULL UNIQUE,
    price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO books (title, author, isbn, price, description) VALUES
    ('The Go Programming Language', 'Alan A. A. Donovan and Brian W. Kernighan', '9780134190440', 44.99, 'A practical guide to Go.'),
    ('Learning Go', 'Jon Bodner', '9781492077213', 49.99, 'Idiomatic Go for real-world projects.'),
    ('Concurrency in Go', 'Katherine Cox-Buday', '9781491941195', 39.99, 'Patterns and tools for concurrent Go programs.')
ON DUPLICATE KEY UPDATE
    title = VALUES(title),
    author = VALUES(author),
    price = VALUES(price),
    description = VALUES(description);
