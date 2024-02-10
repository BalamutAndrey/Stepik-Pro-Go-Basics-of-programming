package main

import (
	"fmt"
)

func main() {
	var a, b, n int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&n)
	price := ((a * 100) + b) * n
	fmt.Println(price/100, price%100)
}
