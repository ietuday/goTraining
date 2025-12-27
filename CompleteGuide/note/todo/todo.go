package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrInvalidText = errors.New("invalid todo text: cannot be empty")

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return Todo{}, ErrInvalidText
	}
	return Todo{Text: text}, nil
}

func (t Todo) Display() {
	fmt.Println("\n--- TODO ---")
	fmt.Println("Text:", t.Text)
}

func (t Todo) Save() error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("todo.json", data, 0644)
}
