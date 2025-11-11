package main

import (
	"fmt"
	"time"
)

func main() {
	// Current time
	now := time.Now()
	fmt.Println("Current time:", now)

	// Format examples
	fmt.Println("1) Date only:", now.Format("02-Jan-2006"))
	fmt.Println("2) Time only:", now.Format("15:04:05"))
	fmt.Println("3) Date & Time:", now.Format("02-Jan-2006 15:04:05"))
	fmt.Println("4) Custom format:", now.Format("Monday, 02 January 2006, 03:04 PM"))

	// Individual components
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())

	// Add or subtract time
	nextWeek := now.AddDate(0, 0, 7)
	fmt.Println("Next week:", nextWeek.Format("02-Jan-2006"))

	lastMonth := now.AddDate(0, -1, 0)
	fmt.Println("Last month:", lastMonth.Format("02-Jan-2006"))

	// Sleep example (delay)
	fmt.Println("Waiting for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done waiting!")

	dateStr := "25-12-2025"
	layout := "02-01-2006" // pattern must match
	parsed, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed date:", parsed)
}
