package main

import (
	"fmt"
)

func main() {
	var f, c float64
	fmt.Scan(&f)
	c = (f - 32) * (5.0 / 9.0)

	fmt.Println(c)
}
