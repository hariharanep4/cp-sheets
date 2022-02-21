package main

import (
	"fmt"
)

func main() {
	var x, y int64
	fmt.Scanf("%d %d", &x, &y)
	fmt.Printf("%d + %d = %d\n", x, y, (x + y))
	fmt.Printf("%d * %d = %d\n", x, y, (x * y))
	fmt.Printf("%d - %d = %d\n", x, y, (x - y))
}