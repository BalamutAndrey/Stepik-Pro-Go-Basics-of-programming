package main

import (
	"fmt"
)

func main() {
	var price1, price2, price3, price4 int
	fmt.Scan(&price1)
	fmt.Scan(&price2)
	fmt.Scan(&price3)
	fmt.Scan(&price4)
	price := (price1 + price2 + price3 + price4) * 3

	fmt.Println(price)
}
