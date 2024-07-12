package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num == 0 {
		fmt.Println("0")
	} else if num > 0 {
		fmt.Println("1")
	} else {
		fmt.Println("-1")
	}
}
