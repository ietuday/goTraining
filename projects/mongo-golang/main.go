package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Author      string             `bson:"author" json:"author"`
	ISBN        string             `bson:"isbn" json:"isbn"`
	Price       float64            `bson:"price" json:"price"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type BookStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func main() {
	client, err := openMongo()
	if err != nil {
		log.Fatal(err)
	}
	defer disconnectMongo(client)

	dbName := env("MONGO_DATABASE", "bookstore")
	store := &BookStore{
		client:     client,
		collection: client.Database(dbName).Collection("books"),
	}

	if err := store.ensureIndexes(context.Background()); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/health", store.health).Methods(http.MethodGet)
	router.HandleFunc("/api/books", store.listBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", store.getBook).Methods(http.MethodGet)
	router.HandleFunc("/api/books", store.createBook).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id}", store.updateBook).Methods(http.MethodPut)
	router.HandleFunc("/api/books/{id}", store.deleteBook).Methods(http.MethodDelete)

	addr := ":" + env("PORT", "8000")
	log.Printf("MongoDB bookstore API started on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func openMongo() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env("MONGO_URI", "mongodb://localhost:27017")))
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		if err := client.Ping(ctx, nil); err == nil {
			return client, nil
		}

		select {
		case <-ctx.Done():
			disconnectMongo(client)
			return nil, errors.New("database is not ready")
		case <-ticker.C:
		}
	}
}

func disconnectMongo(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Printf("failed to disconnect MongoDB client: %v", err)
	}
}

func (s *BookStore) ensureIndexes(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := s.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "isbn", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

func (s *BookStore) health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := s.client.Ping(ctx, nil); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "database unavailable"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *BookStore) listBooks(w http.ResponseWriter, r *http.Request) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := s.collection.Find(r.Context(), bson.D{}, opts)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list books")
		return
	}
	defer cursor.Close(r.Context())

	books := make([]Book, 0)
	if err := cursor.All(r.Context(), &books); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to read books")
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

	book, err := s.findBookByID(r.Context(), id)
	if errors.Is(err, mongo.ErrNoDocuments) {
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

	now := time.Now().UTC()
	book.ID = primitive.NilObjectID
	book.Title = strings.TrimSpace(book.Title)
	book.Author = strings.TrimSpace(book.Author)
	book.ISBN = strings.TrimSpace(book.ISBN)
	book.Description = strings.TrimSpace(book.Description)
	book.CreatedAt = now
	book.UpdatedAt = now

	result, err := s.collection.InsertOne(r.Context(), book)
	if isDuplicateKey(err) {
		writeError(w, http.StatusConflict, "isbn already exists")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create book")
		return
	}

	book.ID = result.InsertedID.(primitive.ObjectID)
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

	update := bson.M{
		"$set": bson.M{
			"title":       strings.TrimSpace(book.Title),
			"author":      strings.TrimSpace(book.Author),
			"isbn":        strings.TrimSpace(book.ISBN),
			"price":       book.Price,
			"description": strings.TrimSpace(book.Description),
			"updated_at":  time.Now().UTC(),
		},
	}

	result, err := s.collection.UpdateByID(r.Context(), id, update)
	if isDuplicateKey(err) {
		writeError(w, http.StatusConflict, "isbn already exists")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update book")
		return
	}
	if result.MatchedCount == 0 {
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

	result, err := s.collection.DeleteOne(r.Context(), bson.M{"_id": id})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete book")
		return
	}
	if result.DeletedCount == 0 {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *BookStore) findBookByID(ctx context.Context, id primitive.ObjectID) (Book, error) {
	var book Book
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	return book, err
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

func bookID(r *http.Request) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(mux.Vars(r)["id"])
}

func isDuplicateKey(err error) bool {
	var writeException mongo.WriteException
	return errors.As(err, &writeException) && writeException.HasErrorCode(11000)
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
