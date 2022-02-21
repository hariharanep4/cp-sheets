package main

import (
	"fmt"
)

func main() {
	var i int32
	var l int64
	var c byte
	var f float32
	var d float64

	fmt.Scanf("%d", &i)
	fmt.Scanf("%d", &l)
	fmt.Scanf("%c", &c)
	fmt.Scanf("%f", &f)
	fmt.Scanf("%f", &d)

	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", l)
	fmt.Printf("%c\n", c)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%.1f\n", d)
}