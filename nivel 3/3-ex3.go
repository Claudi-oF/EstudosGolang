package main

import (
	"fmt"
)

func main() {

	x := 0

	for i := 2005; i <= 2024; i++ {
		fmt.Println(i)

		if x == 1 {
			fmt.Printf("%v ano \n\n", x)
		}

		if x != 1 {
			fmt.Printf("%v anos \n\n", x)
		}

		x++
	}
}
