package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(((n % 10) + (n % 100 / 10) + (n % 1000 / 100)))
}
