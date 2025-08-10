package main

import (
	"fmt"
	"math"
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("What is your tax rate (in precentage): ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))
	ratio := math.RoundToEven(earningsBeforeTax / profit)

	fmt.Println("Earnings before taxes: ", earningsBeforeTax)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)

}
