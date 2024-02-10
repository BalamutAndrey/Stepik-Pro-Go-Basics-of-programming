package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	n1 := n % 10
	n2 := n % 100 / 10
	n3 := n % 1000 / 100
	result := (n1 * 100) + (n2 * 10) + n3
	fmt.Println(result)
}
