package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	_ = scaner.Scan()
	string1 := scaner.Text()
	_ = scaner.Scan()
	string2 := scaner.Text()
	_ = scaner.Scan()
	string3 := scaner.Text()
	fmt.Println(string3)
	fmt.Println(string2)
	fmt.Println(string1)
}
