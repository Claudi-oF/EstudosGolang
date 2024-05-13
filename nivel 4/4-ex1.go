package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array)

	for i, v := range array {
		fmt.Println("indice:", i, " valor:", v)
	}

	fmt.Printf("%T", array) //tipo do array
}
