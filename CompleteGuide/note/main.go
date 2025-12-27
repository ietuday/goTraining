package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/note/note"
	"example.com/note/todo"
)

// Interfaces
type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

// Item = any type that can Display + Save
type Item interface {
	Saver
	Displayer
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Create a Note
	title := mustGetUserInput(reader, "Note title: ")
	content := mustGetUserInput(reader, "Note content: ")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("❌ Note error:", err)
		return
	}

	// Create a Todo
	todoText := mustGetUserInput(reader, "Todo text: ")
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("❌ Todo error:", err)
		return
	}

	// Use interface-based common logic (no duplication)
	if err := HandleAndSave(userNote); err != nil {
		fmt.Println("❌ Saving note failed:", err)
		return
	}
	fmt.Println("✅ Note saved successfully!")

	if err := HandleAndSave(userTodo); err != nil {
		fmt.Println("❌ Saving todo failed:", err)
		return
	}
	fmt.Println("✅ Todo saved successfully!")

	// -------------------------------
	// any + extracting types examples
	// -------------------------------
	fmt.Println("\n--- any + type extraction demo ---")

	var values []any
	values = append(values, userNote)          // stored as any
	values = append(values, userTodo)          // stored as any
	values = append(values, "hello")           // string
	values = append(values, 123)               // int
	values = append(values, true)              // bool

	for _, v := range values {
		// 1) Try extracting to interface (Item) using type assertion
		if it, ok := v.(Item); ok {
			fmt.Printf("\nValue %T satisfies Item => handle it:\n", v)
			_ = HandleAndSave(it)
			continue
		}

		// 2) If not Item, use type switch to handle different types
		switch val := v.(type) {
		case string:
			fmt.Println("string:", val)
		case int:
			fmt.Println("int:", val)
		case bool:
			fmt.Println("bool:", val)
		default:
			fmt.Printf("unknown type: %T value=%v\n", v, v)
		}
	}

	// 3) Concrete type extraction from interface (type switch on Item slice)
	fmt.Println("\n--- Item slice + concrete type extraction ---")
	items := []Item{userNote, userTodo}

	for _, it := range items {
		switch x := it.(type) {
		case note.Note:
			fmt.Println("Detected NOTE -> Title:", x.Title)
		case todo.Todo:
			fmt.Println("Detected TODO -> Text:", x.Text)
		default:
			fmt.Printf("Unknown Item: %T\n", it)
		}
	}
}

func HandleAndSave(data Item) error {
	data.Display()
	return data.Save()
}
