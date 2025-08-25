package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/ishowdarkside/golang/bank/fileops"
)

func main() {

	accountBalance, _ := fileops.ReadFloatFromFile("balance.txt", 2000)
	fmt.Println(randomdata.Country(2))

	fmt.Println("Welcome to Go Bank!")

	for {

		PresentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			{
				fmt.Println("Your balance is:", accountBalance)
			}
		case 2:
			{

				var depositAmount float64
				fmt.Print("Your deposit:")
				fmt.Scan(&depositAmount)

				if depositAmount <= 0 {
					fmt.Println("------------------------")
					fmt.Println("Invalid depoist amount.")
					continue
				}
				accountBalance += depositAmount
				fileops.WriteFloatToFile(accountBalance, "balance.txt")
				fmt.Println("Balance updated:", accountBalance)
			}

		case 3:
			{
				var withdrawAmount float64
				fmt.Print("How much would you like to withdraw: ")
				fmt.Scan(&withdrawAmount)

				if withdrawAmount <= 0 {
					fmt.Println("------------------------")
					fmt.Println("Invalid withdraw amount.")
					continue
				}

				if withdrawAmount > accountBalance {
					fmt.Println("------------------------")
					fmt.Println("You don't have enough money on your balance to perform this operation")
					continue
				} else {
					accountBalance -= withdrawAmount

					fileops.WriteFloatToFile(accountBalance, "balance.txt")
					fmt.Println("Balance updated: ", accountBalance)
				}

			}
		default:
			{
				fmt.Print("Goodbye!")
				return
			}

		}
	}
}
