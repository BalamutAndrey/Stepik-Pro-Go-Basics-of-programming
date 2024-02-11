package main

import (
	"fmt"
	"math"
)

func main() {
	var kat1, kat2, gip float64
	fmt.Scan(&kat1, &kat2)
	gip = math.Sqrt((math.Pow(kat1, 2) + math.Pow(kat2, 2)))
	per := kat1 + kat2 + gip

	fmt.Println(per)
}
