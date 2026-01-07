package main

func Calculate(a, b int, op func(int, int) int) int {
	return op(a, b)
}
