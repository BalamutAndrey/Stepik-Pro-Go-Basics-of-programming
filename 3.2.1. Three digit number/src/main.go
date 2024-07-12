package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num >= 100 && num < 1000 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
