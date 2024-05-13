package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Primeiro ao terceiro item:\n")
	for i, v := range slice {
		if i < 3 {
			fmt.Printf("Indice: %v Valor: %v \n", i, v)
		}
	}
	fmt.Printf("\n")

	fmt.Printf("Do quinto ao ultimo item:\n")
	for i, v := range slice {
		if i > 3 {
			fmt.Printf("Indice: %v Valor: %v \n", i, v)
		}
	}
	fmt.Printf("\n")

	fmt.Printf("Do segundo ao sétimo item: \n")
	for i, v := range slice {
		if i > 0 && i < 7 {
			fmt.Printf("Indice: %v Valor: %v \n", i, v)
		}
	}
	fmt.Printf("\n")

	fmt.Printf("Do terceiro ao penultimo item: \n")
	for i, v := range slice {
		if i > 1 && i < len(slice)-1 {
			fmt.Printf("Indice: %v Valor: %v \n", i, v)
		}
	}

	fmt.Printf("Pode fazer dessa forma: \n") //modo mais facil
	//    posicao:    0  1  2  3  4  5  6  7  8
	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:8])
}
