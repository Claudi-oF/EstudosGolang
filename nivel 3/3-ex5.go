package main

import (
	"fmt"
)

var resto int

func main() {

	for i := 10; i < 100; i++ {
		resto = i % 4

		fmt.Printf("de %v: %v\n", i, resto)
	}
}
