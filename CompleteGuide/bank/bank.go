package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, error = fileops.GetFloatFromFile(accountBalanceFile)
	if error != nil {
		fmt.Println("ERROR")
		fmt.Println(error)
		fmt.Println("+++++++++++++++++++++++++++++++++")
		//panic("Can't continue.......................")
	}
	fmt.Println("Welcome to the GO Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	for {
		presentOptions()
		var choice int
		fmt.Print("Your Choice: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Account Balance: ", accountBalance)

		case 2:
			var depositAmount float64
			fmt.Print("Amount to Deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}
			fmt.Println("User Input to deposit : ", depositAmount)
			accountBalance += depositAmount
			fmt.Println("Updated Balance : ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawAmount float64
			fmt.Print("Amount to Deposit: ")
			fmt.Scan(&withdrawAmount)
			fmt.Println("User Input to deposit : ", withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid Amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Updated Balance : ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Good Bye")
			fmt.Println("Thanks for Choosing Our Bank")
			return
		}
	}

}


