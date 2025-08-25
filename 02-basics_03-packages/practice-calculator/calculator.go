package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {

	revenue := retrieveUserInput("Enter your revenue")
	expenses := retrieveUserInput("Enter your expenses")
	taxRate := retrieveUserInput("What is your tax rate (in precentage)")

	earningsBeforeTax, profit, ratio := calcEarnings(revenue, expenses, taxRate)

	printValues(earningsBeforeTax, profit, ratio)
	writeToFile(earningsBeforeTax, profit, ratio)
}

func printValues(earningsBeforeTax float64, profit float64, ratio float64) {

	fmt.Println("Earnings before taxes: ", earningsBeforeTax)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}

func calcEarnings(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))
	ratio := math.RoundToEven(earningsBeforeTax / profit)

	return earningsBeforeTax, profit, ratio
}

func retrieveUserInput(text string) float64 {

	var value float64
	fmt.Printf("%s: ", text)
	_, err := fmt.Scan(&value)

	var invalidInputErr error
	if err != nil {
		invalidInputErr = errors.New("invalid input")
		fmt.Println(invalidInputErr)
		os.Exit(1)
	}

	if value <= 0 {
		invalidInputErr = errors.New("invalid value")
		fmt.Println(invalidInputErr)
		os.Exit(1)
	}

	return float64(value)
}

func writeToFile(revenue float64, expenses float64, taxRate float64) {

	stringifiedVal := fmt.Sprintf(`
	Revenue: %f2
	Expenses: %f2
	Tax Rate: %f2
	`, revenue, expenses, taxRate)
	err := os.WriteFile("output.txt", []byte(stringifiedVal), 0644)

	if err != nil {
		err := errors.New("something went wrong writing to file. Please try again")
		fmt.Println(err)
		os.Exit(1)
	}

}
