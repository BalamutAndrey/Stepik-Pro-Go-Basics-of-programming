package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	_ = scaner.Scan()
	separator := scaner.Text()
	_ = scaner.Scan()
	string1 := scaner.Text()
	_ = scaner.Scan()
	string2 := scaner.Text()
	_ = scaner.Scan()
	string3 := scaner.Text()

	fmt.Println(string1 + separator + string2 + separator + string3)
}
