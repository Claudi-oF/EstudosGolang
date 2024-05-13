package main

import (
	"fmt"
)

func main() {
	slice := []string{"banana", "maçã", "pera"}

	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor:", valor)
	}

	fmt.Printf("\n")

	slice = append(slice, "macarrao DO NADA") //append: adiciona um elemento ao slice!!! WOOOOOOOOOW

	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor:", valor)
	}

	//indice é a posicao dentro do slice
	//valor é o valor da posicao "indice" do slice
	//range percorre o slice.
}
