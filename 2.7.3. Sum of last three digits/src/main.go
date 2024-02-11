package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	result := (n % 10) + (n / 10 % 10) + (n / 100 % 10)
	fmt.Println(result)
}
