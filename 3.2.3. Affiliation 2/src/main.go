package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num <= -3 || num >= 7 {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}
