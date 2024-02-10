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
	n, _ := strconv.Atoi(scaner.Text())
	_ = scaner.Scan()
	k, _ := strconv.Atoi(scaner.Text())
	result := k / n
	fmt.Println(result)
}
