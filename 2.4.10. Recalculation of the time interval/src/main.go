package main

import (
	"fmt"
)

func main() {
	var min int
	fmt.Scan(&min)
	fmt.Println(min, "мин - это", min/60, "час", min%60, "минут.")
}
