package main

import "fmt"

func removeDuplicate(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

func main() {
	nums := []int{4, 4, 6, 3, 7, 5, 4, 7, 9, 3, 5}
	fmt.Println(nums)
	result := removeDuplicate(nums)
	fmt.Println("Without Duplicate: ", result)
}
