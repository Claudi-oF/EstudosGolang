package main

import (
	"fmt"
)

type banana int

var x banana = 87

func main() {
	fmt.Println(x)
	fmt.Printf("%T\t", x)

	x = 42

	fmt.Println(x)
}
