package main

import "fmt"

type tranformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	morenumbers := []int{5, 1, 22}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
	tranformFn1 := getTransformerFunction(&numbers)
	tranformFn2 := getTransformerFunction(&morenumbers)

	transferedNumbers := transformNumbers(&numbers, tranformFn1)
	moretransferedNumbers := transformNumbers(&numbers, tranformFn2)

	fmt.Println(transferedNumbers)
	fmt.Println(moretransferedNumbers)

	result := Calculate(10, 5, func(x, y int) int {
		return x * y
	})

	fmt.Println(result) // 50

	mul2 := makeMultiplier(2)
	mul3 := makeMultiplier(3)
	mul10 := makeMultiplier(10)
	fmt.Println(transformNumbers(&numbers, mul2))
	fmt.Println(transformNumbers(&numbers, mul3))
	fmt.Println(transformNumbers(&numbers, mul10))

	countMul2 := makeCountingMultiplier(2)
	fmt.Println(transformNumbers(&numbers, countMul2))
	fmt.Println(transformNumbers(&numbers, countMul2))
	fmt.Println(factorial(5))
	fact := func(n int) int {
		result := 1
		for i := 1; i <= n; i++ {
			result *= i
		}
		return result
	}

	fmt.Println(fact(5)) // 120
	fmt.Println(FactorialRec(6))

	fmt.Println(sumup(1,2,3,5,23))
	fmt.Println(sumup(numbers...))

}

func getTransformerFunction(numbers *[]int) tranformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformNumbers(numbers *[]int, tranform tranformFn) []int {
	dNumbers := []int{}
	for index, val := range *numbers {
		fmt.Println(index)
		dNumbers = append(dNumbers, tranform(val))

	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func makeMultiplier(factor int) tranformFn {
	return func(n int) int {
		return n * factor
	}
}

func makeCountingMultiplier(factor int) tranformFn {
	count := 0
	return func(n int) int {
		count++
		fmt.Println("call #", count, "factor:", factor)
		return n * factor
	}
}
