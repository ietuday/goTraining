package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email,omitempty"` // omit if empty
	Hobbies []string `json:"hobbies"`
}

func main() {
	p := Person{
		Name: "Uday",
		Age:  35,
		Hobbies: []string{
			"GO", "Docker", "Kubernetes",
		},
	}

	data, err := json.MarshalIndent(p, "", " ")

	if err != nil {
		panic(err)
	}

	fmt.Println("JSON output:\n" + string(data))

	// Save to file
	if err := os.WriteFile("person.json", data, 0644); err != nil {
		panic(err)
	}
	fmt.Println("Wrote person.json")

	// ----- Unmarshal JSON -> Go -----
	raw := `{"name":"Amit","age":30,"hobbies":["Travel","Guitar"]}`
	var q Person
	if err := json.Unmarshal([]byte(raw), &q); err != nil {
		panic(err)
	}
	fmt.Printf("Decoded struct: %+v\n", q)

	// Read from file and unmarshal
	fileBytes, err := os.ReadFile("person.json")
	if err != nil {
		panic(err)
	}
	var fromFile Person
	if err := json.Unmarshal(fileBytes, &fromFile); err != nil {
		panic(err)
	}
	fmt.Printf("From file struct: %+v\n", fromFile)

}
