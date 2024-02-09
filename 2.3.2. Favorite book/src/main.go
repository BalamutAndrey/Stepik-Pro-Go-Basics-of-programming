package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	_ = scaner.Scan()
	name := scaner.Text()
	fmt.Println(name, "- лучшая книга!")
}
