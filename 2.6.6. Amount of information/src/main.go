package main

import (
	"fmt"
)

func main() {
	var bit float64
	fmt.Scan(&bit)
	byte := bit / 8
	kbyte := byte / 1024

	fmt.Println(kbyte)
}
