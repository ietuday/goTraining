package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func writeJSONLines(path string, books []Book) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, b := range books {
		bytes, err := json.Marshal(b)
		if err != nil {
			return err
		}
		if _, err := w.Write(append(bytes, '\n')); err != nil {
			return err
		}
	}
	return w.Flush()
}

func readJSONLines(path string) ([]Book, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []Book
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var b Book
		if err := json.Unmarshal(sc.Bytes(), &b); err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, sc.Err()
}

func main() {
	books := []Book{
		{"The Go Programming Language", "Donovan", 2015},
		{"Go in Action", "Kennedy", 2015},
		{"Cloud Native Go", "A-Liang", 2021},
	}
	const file = "books.jsonl"

	if err := writeJSONLines(file, books); err != nil {
		panic(err)
	}
	fmt.Println("Wrote", file)

	got, err := readJSONLines(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read back:")
	for _, b := range got {
		fmt.Printf("%s by %s (%d)\n", b.Title, b.Author, b.Year)
	}
}
