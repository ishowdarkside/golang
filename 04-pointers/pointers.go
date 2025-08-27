package main

import "fmt"

func main() {

	var age int
	var agePointer *int = &age
	fmt.Println("Age", *agePointer)

	getAdultYears(&age)
	fmt.Println(age)
}

func getAdultYears(age *int) {

	*age -= 12

}
