package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	if num1%num2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
