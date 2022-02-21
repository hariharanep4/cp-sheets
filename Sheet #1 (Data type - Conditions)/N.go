package main

import (
	"fmt"
)

func main() {
	var x byte

	fmt.Scanf("%c", &x)

	if x >= 65 && x <= 90 {
		fmt.Printf("%c\n", x + 32)
	} else {
		if x >= 95 && x <= 122 {
			fmt.Printf("%c\n", x - 32)
		}
	}
}