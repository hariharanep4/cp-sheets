package main

import (
	"fmt"
)

const PI float64 = 3.141592653

func main() {
	var radius float64

	fmt.Scanf("%f", &radius)

	area := PI * (radius * radius)

	fmt.Printf("%.9f\n", area)

}