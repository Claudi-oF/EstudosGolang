package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	fmt.Printf("\n")

	y := 0

	for y < 10 {
		fmt.Println(y)
		y++
	}
}
