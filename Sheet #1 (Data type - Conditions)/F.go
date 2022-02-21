package main

import (
	"fmt"
)

func main() {
	var n, m int64
	fmt.Scanf("%d %d", &n, &m)

	fmt.Println((n % 10) + (m % 10))
}