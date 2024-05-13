package main

import (
	"fmt"
)

func main() {
	//                0.	  .1          .2         .3
	slice := []string{"mouse", "teclado", "monitor", "pc"}

	fatia := slice[0:len(slice)] // ele literalmente FATIA o slice e pega os valores entre o intervalo que voce fatiou

	// a funcao "len()" pega o tamanho TOTAL do slice, array, do que vc quiser

	fmt.Println(fatia)

	//para excluir um valor dentro de um slice:
	//                  0.	      1.        2.        3.      4.
	slice2 := []string{"macarrao", "batata", "carne", "cebola"}

	slice2 = append(slice2[:1], slice2[2:]...) // vai excluir a string "batata"

	fmt.Println(slice2)
	fmt.Printf("\n")

	//adicionando valores na slice:
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{10, 11, 12, 13}

	fmt.Println(umaslice)
	umaslice = append(umaslice, 5, 6, 7, 8, 9)
	fmt.Println(umaslice)

	//agora com outraslice
	fmt.Println(outraslice)
	outraslice = append(outraslice, 14, 15, 16, 17)
	fmt.Println(outraslice)

	//juntando as duas:
	umaslice = append(umaslice, outraslice...) // os "..." faz juntar as duas slices, sem eles da erro
	fmt.Println(umaslice)
}
