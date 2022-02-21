package main

import (
	"fmt"
)

func main() {
	var a, b, c int64 

	fmt.Scanf("%d %d %d", &a, &b, &c)

	var max, min int64 = -100000, 100000

	// max := -100000
	// min := 100000

	if a >= b && a >= c {
		max = a 
	} else if b >= a && b >= c {
		max = b 
	} else {
		max = c 
	}

	if a <=b && a <= c {
		min = a 
	} else if b <= a && b <= c {
		min = b 
	} else {
		min = c 
	}

	fmt.Printf("%d %d\n", min, max)
}