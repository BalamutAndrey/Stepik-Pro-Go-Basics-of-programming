package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	string1 := scanner.Text()
	_ = scanner.Scan()
	string2 := scanner.Text()
	if string1 == string2 {
		fmt.Println("Пароль принят")
	} else {
		fmt.Println("Пароль не принят")
	}
}
