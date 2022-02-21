package main

import (
	"fmt"
)

func main() {
	var firstNameOne, secondNameOne string
	var firstNameTwo, secondNameTwo string

	fmt.Scanf("%s %s", &firstNameOne, &secondNameOne)
	fmt.Scanf("%s %s", &firstNameTwo, &secondNameTwo)

	if secondNameOne == secondNameTwo {
		fmt.Println("ARE Brothers")
	} else {
		fmt.Println("NOT")
	}
}