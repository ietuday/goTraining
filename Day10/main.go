package main

import (
	"bufio"
	"fmt"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func writeText(filename string, lines []string) {
	f, err := os.Create(filename)
	must(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		must(err)
	}
	must(w.Flush())
}

func readText(filename string) []string {
	f, err := os.Open(filename)
	must(err)
	defer f.Close()

	var out []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		out = append(out, sc.Text())
	}
	must(sc.Err())
	return out
}

func main() {
	fmt.Println("Day10 Started")
	filename := "notes.txt"
	writeText(filename, []string{
		"Go is fun!",
		"File I/O is easy.",
		"Next up: JSON.",
	})
	fmt.Println("Wrote", filename)

	lines := readText(filename)
	fmt.Println("Read lines:")
	for i, line := range lines {
		fmt.Printf("%d) %s\n", i+1, line)
	}

}
