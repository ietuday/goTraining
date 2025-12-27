package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrInvalidTitle   = errors.New("invalid title: cannot be empty")
	ErrInvalidContent = errors.New("invalid content: cannot be empty")
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func New(title, content string) (Note, error) {
	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)

	if title == "" {
		return Note{}, ErrInvalidTitle
	}
	if content == "" {
		return Note{}, ErrInvalidContent
	}

	return Note{
		Title:   title,
		Content: content,
	}, nil
}

func (n Note) Display() {
	fmt.Println("\n--- NOTE ---")
	fmt.Println("Title  :", n.Title)
	fmt.Println("Content:", n.Content)
}

func (n Note) Save() error {
	data, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("note.json", data, 0644)
}
