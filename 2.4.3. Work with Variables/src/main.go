package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	x2 := x * x
	x4 := x2 * x2
	x6 := x4 * x2
	fmt.Println(x6)
}
