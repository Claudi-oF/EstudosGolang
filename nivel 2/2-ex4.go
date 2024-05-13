package main

import (
	"fmt"
)

var x int = 100

func main() {
	fmt.Printf("%d \t %b \t %#x \n", x, x, x)

	y := x << 1 //desloca o bit para a esquerda

	fmt.Printf("%d \t %b \t %#x", y, y, y)

}
