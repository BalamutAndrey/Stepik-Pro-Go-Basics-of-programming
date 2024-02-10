package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	_ = scaner.Scan()
	num1, _ := strconv.Atoi(scaner.Text())
	_ = scaner.Scan()
	num2, _ := strconv.Atoi(scaner.Text())
	_ = scaner.Scan()
	num3, _ := strconv.Atoi(scaner.Text())
	sqr := num1 * num2 * num3
	fmt.Println(sqr)
}
