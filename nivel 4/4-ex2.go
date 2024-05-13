package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice {
		fmt.Println("indice:", i, " valor:", v)
	}

	fmt.Printf("%T", slice)
}
