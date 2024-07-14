package main

import "fmt"

func main() {
	var num, sum1, sum2 int
	fmt.Scan(&num)
	sum1 = num / 1000
	sum1 = (sum1 / 100) + ((sum1 / 10) % 10) + (sum1 % 10)
	sum2 = ((num / 100) % 10) + ((num / 10) % 10) + (num % 10)

	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
