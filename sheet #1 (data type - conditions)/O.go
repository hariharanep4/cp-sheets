package main

import (
	"fmt"
)

func main() {
	var a, b int64
	var o byte

	fmt.Scanf("%d%c%d", &a, &o, &b)

	if o == '+' {
		fmt.Println(a + b)
	} else if o == '-' {
		fmt.Println(a - b)
	} else if o == '*' {
		fmt.Println(a * b)
	} else {
		fmt.Println(a / b)
	}
}