package main

import (
	"fmt"
)

func main() {

	switch {
	case 10 > 20:
		fmt.Printf("10 maior que 20")
	case 10 == 20:
		fmt.Printf("10 igual que 20")
	case 10 < 20:
		fmt.Printf("10 menor que 20")
	}
}
