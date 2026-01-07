package main

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}

func FactorialRec(number int) int {
	if number < 0 {
		panic("factorial not defined for negative numbers")
	}
	if number == 0 {
		return 1
	}
	return number * FactorialRec(number-1)
}

func FactorialClosesure() func(int) int {
	result := 1

	return func(n int) int {
		for i := 1; i <= n; i++ {
			result *= i
		}
		return result
	}
}
