package main

import "fmt"

// reverse s[l:r] in-place (r exclusive)
func reverse[T any](s []T, l, r int) {
	for l < r-1 {
		s[l], s[r-1] = s[r-1], s[l]
		l++
		r--
	}
}

// A) RotateLeft: in-place, O(1) extra space
func RotateLeft[T any](s []T, k int) {
	n := len(s)
	if n == 0 {
		return
	}
	k %= n
	if k == 0 {
		return
	}
	reverse(s, 0, k)
	reverse(s, k, n)
	reverse(s, 0, n)
}

// B) PopFrontK: returns (popped, rest) with trimmed capacity
func PopFrontK[T any](s []T, k int) ([]T, []T) {
	if k <= 0 {
		return nil, append([]T(nil), s...) // trim cap
	}
	if k >= len(s) {
		return append([]T(nil), s...), nil
	}
	popped := append([]T(nil), s[:k]...) // stable copy
	rest := append([]T(nil), s[k:]...)   // force new array (releases old)
	return popped, rest
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	RotateLeft(a, 2)
	fmt.Println(a) // [3 4 5 1 2]

	popped, rest := PopFrontK([]int{10, 20, 30, 40, 50}, 2)
	fmt.Println(popped, rest) // [10 20] [30 40 50]
}
