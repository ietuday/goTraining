package main

import "fmt"

func countFreq(nums []int) map[int]int {
	freq := make(map[int]int)

	for _, n := range nums {
		freq[n]++
	}
	return freq
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 3}
	fmt.Println(nums)
	freq := countFreq(nums)

	for value, count := range freq {
		fmt.Printf("%d => %d", value, count)
		fmt.Println("")

	}
}
