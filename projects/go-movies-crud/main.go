package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies = []Movie{
	{
		ID:    "1",
		Isbn:  "438227",
		Title: "Movie One",
		Director: &Director{
			ID:        "101",
			FirstName: "John",
			LastName:  "Doe",
		},
	},
	{
		ID:    "2",
		Isbn:  "454555",
		Title: "Movie Two",
		Director: &Director{
			ID:        "102",
			FirstName: "Steve",
			LastName:  "Smith",
		},
	},
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", getMovies).Methods(http.MethodGet)
	router.HandleFunc("/api/movies/{id}", getMovie).Methods(http.MethodGet)
	router.HandleFunc("/api/movies", createMovie).Methods(http.MethodPost)
	router.HandleFunc("/api/movies/{id}", updateMovie).Methods(http.MethodPut)
	router.HandleFunc("/api/movies/{id}", deleteMovie).Methods(http.MethodDelete)

	log.Println("server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			writeJSON(w, http.StatusOK, movie)
			return
		}
	}

	writeJSON(w, http.StatusNotFound, map[string]string{"error": "movie not found"})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(1000000))
	if movie.Director != nil && movie.Director.ID == "" {
		movie.Director.ID = strconv.Itoa(rand.Intn(1000000))
	}

	movies = append(movies, movie)
	writeJSON(w, http.StatusCreated, movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updated Movie

	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	for index, movie := range movies {
		if movie.ID == params["id"] {
			updated.ID = movie.ID
			if updated.Director != nil && updated.Director.ID == "" {
				updated.Director.ID = movie.DirectorID()
			}

			movies[index] = updated
			writeJSON(w, http.StatusOK, updated)
			return
		}
	}

	writeJSON(w, http.StatusNotFound, map[string]string{"error": "movie not found"})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	writeJSON(w, http.StatusNotFound, map[string]string{"error": "movie not found"})
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}

func (m Movie) DirectorID() string {
	if m.Director == nil {
		return ""
	}

	return m.Director.ID
}
