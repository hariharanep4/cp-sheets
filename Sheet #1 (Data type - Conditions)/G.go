package main

import (
	"fmt"
)

func main() {
	var n int64

	fmt.Scanf("%d", &n)

	fmt.Printf("%d\n", (n * (n + 1)) / 2)
}