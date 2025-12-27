package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investementAmount, years float64 = 1000.0, 10.0
	expectedReturnRate := 5.5

	fmt.Print("Enter Investement Amount: ")
	fmt.Scan(&investementAmount)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	fmt.Print("Enter Expected ReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investementAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//fmt.Println("futureValue : ",futureValue)
	//fmt.Println("futureRealValue : ", futureRealValue)
	formattedString := fmt.Sprintf("futureValue : %.2f\n", futureValue)

	fmt.Printf("futureValue : %.2f\n", futureValue)
	fmt.Printf("futureRealValue : %.2f\n", futureRealValue)
	fmt.Print(formattedString)
}
