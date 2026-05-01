package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Book struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	ISBN        string  `json:"isbn"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
}

type BookStore struct {
	db *sql.DB
}

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := &BookStore{db: db}
	router := mux.NewRouter()

	router.HandleFunc("/health", store.health).Methods(http.MethodGet)
	router.HandleFunc("/api/books", store.listBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id:[0-9]+}", store.getBook).Methods(http.MethodGet)
	router.HandleFunc("/api/books", store.createBook).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id:[0-9]+}", store.updateBook).Methods(http.MethodPut)
	router.HandleFunc("/api/books/{id:[0-9]+}", store.deleteBook).Methods(http.MethodDelete)

	addr := ":" + env("PORT", "8000")
	log.Printf("bookstore API started on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func openDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = env("DB_USER", "bookstore")
	cfg.Passwd = env("DB_PASSWORD", "bookstore")
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", env("DB_HOST", "localhost"), env("DB_PORT", "3306"))
	cfg.DBName = env("DB_NAME", "bookstore")
	cfg.ParseTime = true
	cfg.Collation = "utf8mb4_unicode_ci"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		if err := db.PingContext(ctx); err == nil {
			return db, nil
		}

		select {
		case <-ctx.Done():
			db.Close()
			return nil, errors.New("database is not ready")
		case <-ticker.C:
		}
	}
}

func (s *BookStore) health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := s.db.PingContext(ctx); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "database unavailable"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *BookStore) listBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.QueryContext(r.Context(), `
		SELECT id, title, author, isbn, price, description, created_at, updated_at
		FROM books
		ORDER BY id DESC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list books")
		return
	}
	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		book, err := scanBook(rows)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to read books")
			return
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list books")
		return
	}

	writeJSON(w, http.StatusOK, books)
}

func (s *BookStore) getBook(w http.ResponseWriter, r *http.Request) {
	id, err := bookID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	row := s.db.QueryRowContext(r.Context(), `
		SELECT id, title, author, isbn, price, description, created_at, updated_at
		FROM books
		WHERE id = ?`, id)

	book, err := scanBook(row)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to read book")
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func (s *BookStore) createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validateBook(book); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := s.db.ExecContext(r.Context(), `
		INSERT INTO books (title, author, isbn, price, description)
		VALUES (?, ?, ?, ?, ?)`,
		strings.TrimSpace(book.Title),
		strings.TrimSpace(book.Author),
		strings.TrimSpace(book.ISBN),
		book.Price,
		strings.TrimSpace(book.Description),
	)
	if isDuplicateISBN(err) {
		writeError(w, http.StatusConflict, "isbn already exists")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create book")
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to read created book")
		return
	}

	book, err = s.findBookByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to read created book")
		return
	}

	writeJSON(w, http.StatusCreated, book)
}

func (s *BookStore) updateBook(w http.ResponseWriter, r *http.Request) {
	id, err := bookID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validateBook(book); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := s.db.ExecContext(r.Context(), `
		UPDATE books
		SET title = ?, author = ?, isbn = ?, price = ?, description = ?
		WHERE id = ?`,
		strings.TrimSpace(book.Title),
		strings.TrimSpace(book.Author),
		strings.TrimSpace(book.ISBN),
		book.Price,
		strings.TrimSpace(book.Description),
		id,
	)
	if isDuplicateISBN(err) {
		writeError(w, http.StatusConflict, "isbn already exists")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update book")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update book")
		return
	}
	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	book, err = s.findBookByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to read updated book")
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func (s *BookStore) deleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := bookID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	result, err := s.db.ExecContext(r.Context(), "DELETE FROM books WHERE id = ?", id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete book")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete book")
		return
	}
	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *BookStore) findBookByID(ctx context.Context, id int64) (Book, error) {
	row := s.db.QueryRowContext(ctx, `
		SELECT id, title, author, isbn, price, description, created_at, updated_at
		FROM books
		WHERE id = ?`, id)

	return scanBook(row)
}

type bookScanner interface {
	Scan(dest ...any) error
}

func scanBook(scanner bookScanner) (Book, error) {
	var book Book
	var createdAt time.Time
	var updatedAt time.Time

	err := scanner.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.ISBN,
		&book.Price,
		&book.Description,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return Book{}, err
	}

	book.CreatedAt = createdAt.Format(time.RFC3339)
	book.UpdatedAt = updatedAt.Format(time.RFC3339)
	return book, nil
}

func validateBook(book Book) error {
	if strings.TrimSpace(book.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(book.Author) == "" {
		return errors.New("author is required")
	}
	if strings.TrimSpace(book.ISBN) == "" {
		return errors.New("isbn is required")
	}
	if book.Price < 0 {
		return errors.New("price cannot be negative")
	}

	return nil
}

func bookID(r *http.Request) (int64, error) {
	return strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
}

func isDuplicateISBN(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func env(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	return value
}
