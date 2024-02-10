package main

import (
	"fmt"
)

func main() {
	var a, b int
	var result float64
	fmt.Scan(&a, &b)
	result = float64(a+b) / 2.0

	fmt.Println(result)
}
