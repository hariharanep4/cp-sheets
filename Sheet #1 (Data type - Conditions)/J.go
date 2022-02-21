package main

import (
	"fmt"
)

func main() {
	var a, b int64

	fmt.Scanf("%d %d", &a, &b)

	if a % b == 0 || b % a == 0 {
		fmt.Println("Multiples")
	} else {
		fmt.Println("No Multiples")
	}
}