package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(balance float64, fileName string) {

	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func ReadFloatFromFile(fileName string, defaultValue float64) (float64, error) {

	byteValue, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
		return defaultValue, errors.New("failed to read file")
	}

	stringifiedValue := string(byteValue)
	numVal, conversionErr := strconv.ParseFloat(stringifiedValue, 64)

	if conversionErr != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return numVal, nil

}
