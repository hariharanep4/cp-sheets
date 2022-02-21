package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int64

	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	fmt.Printf("Difference = %d\n", (a * b) - (c * d))
}