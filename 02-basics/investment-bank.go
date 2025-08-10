package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter your investement amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("What is your expected return rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("For how many years do you plan to invest:")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)

	fmt.Printf("Future Value: %.2f \n", futureValue)
	fmt.Printf("Future Value with inflation : %.2f \n", futureRealValue)

}
