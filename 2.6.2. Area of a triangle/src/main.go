package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	s := 0.5 * a * b

	fmt.Println(s)
}
