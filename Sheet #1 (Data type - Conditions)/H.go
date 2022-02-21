package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Scanf("%f %f", &a, &b)
	fmt.Printf("floor %.f / %.f = %.f\n", a, b, math.Floor(a / b))
	fmt.Printf("ceil %.f / %.f = %.f\n", a, b, math.Ceil(a / b))
	fmt.Printf("round %.f / %.f = %.f\n", a, b, math.Round(a / b))
}