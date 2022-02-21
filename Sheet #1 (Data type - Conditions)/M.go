package main

import (
	"fmt"
)

func main() {
	var c byte

	fmt.Scanf("%c", &c)

	if c >= 48 && c <= 57 {
		fmt.Println("IS DIGIT")
	} else {
		if c >= 65 && c <= 90 {
			fmt.Println("ALPHA")
			fmt.Println("IS CAPITAL")
		} else {
			if c >= 97 && c <= 122 {
				fmt.Println("ALPHA")
				fmt.Println("IS SMALL")
			}
		}
	}
}