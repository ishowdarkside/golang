package main

import (
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
	fmt.Printf("%s:", text)
	_, err := fmt.Scan(&value)
	if err != nil {
		os.Exit(1)
	}

	return float64(value)
}
