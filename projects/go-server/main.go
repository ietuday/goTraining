package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handle form submission
func formHandle(w http.ResponseWriter, r *http.Request) {

	// Allow only POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get values from form
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Response
	fmt.Fprintf(w, "Post request successful\n")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// Handle /hello route
func helloHandle(w http.ResponseWriter, r *http.Request) {

	// Check correct path
	if r.URL.Path != "/hello" {
		http.Error(w, "404: Not Found", http.StatusNotFound)
		return
	}

	// Allow only GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {

	// Serve static files (HTML, CSS, JS)
	fileServer := http.FileServer(http.Dir("./static"))

	// Routes
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandle)
	http.HandleFunc("/hello", helloHandle)

	fmt.Println("Server starting at http://localhost:8080")

	// Start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}