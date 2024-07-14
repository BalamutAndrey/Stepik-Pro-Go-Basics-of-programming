package main

import "fmt"

func main() {
	var num, a, b, c int
	fmt.Scan(&num)
	c = num % 10
	num /= 10
	b = num % 10
	a = num / 10
	if a != b && a != c && b != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
