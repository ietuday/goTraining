package main

import (
	"errors"
	"fmt"
	"math"
)

func secondLargest(nums []int) (int, error) {
	if len(nums) < 2 {
		return 0, errors.New("Need atleast two numbers")

	}
	max1, max2 := math.MinInt, math.MinInt

	fmt.Println	("MAX1 and MAX2", max1, max2)

	for _, n := range nums {
		if n > max1 {
			max2 = max1
			max1 = n
		} else if n > max2 && n != max1 {
			max2 = n
		}
	}

	if max2 == math.MinInt {
		return 0, errors.New("no second element largest (all elements same.....)")
	}

	return max2, nil
}

func main() {
	nums := []int{5, 1, 9, 6, 9, 7, 1}
	fmt.Println(nums)
	max2, error := secondLargest(nums)
	if error != nil {
		fmt.Println("Error: ", error)
		return
	}

	fmt.Println("Second Largest: ", max2)
}
