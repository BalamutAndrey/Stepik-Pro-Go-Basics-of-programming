package main

import (
	"fmt"
)

func main() {
	var num float64
	fmt.Scan(&num)
	temp := int(num)
	result := num - float64(temp)

	fmt.Println(result)
}
