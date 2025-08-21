package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func writeToFile(balance float64) {

	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func readFile() (float64, error) {

	byteValue, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
		return 1000, errors.New("failed to read file")
	}

	balanceText := string(byteValue)
	numVal, conversionErr := strconv.ParseFloat(balanceText, 64)

	if conversionErr != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return numVal, nil

}

func main() {

	accountBalance, err := readFile()

	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Print("------------------------------------")
		return
	}

	fmt.Println("Welcome to Go Bank!")

	for {

		fmt.Println("-----------------------")
		fmt.Println("What do you want to do?")
		fmt.Println(`
1. Check Balance
2. Deposit money
3. Withdraw money
4. Exit
	`)

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
				writeToFile(accountBalance)
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

					writeToFile(accountBalance)
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
