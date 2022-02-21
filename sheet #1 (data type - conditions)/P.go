package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64

	fmt.Scanf("%f", &n)

	digits := int64(math.Log10(n))
	fmt.Println(digits)
}