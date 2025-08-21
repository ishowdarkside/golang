package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter your investement amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("What is your expected return rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("For how many years do you plan to invest:")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)
	formattedValue := fmt.Sprintf("Future Value: %.2f \n", futureValue)
	formattedRealValue := fmt.Sprintf("Future Value with inflation : %.2f \n", futureRealValue)

	fmt.Print(`Sve
	BiH DAO
	DA JE VIDIM
	PONOSNU I LIJEPU
	KO U SNOVIMA!!!
	`)

	outputText(formattedRealValue)
	outputText(formattedValue)
}

func outputText(formattedValue string) {

	fmt.Print(formattedValue)

}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, realFutureValue
}
