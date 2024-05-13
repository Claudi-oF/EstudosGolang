package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6, 7, 8, 9, 10},
		[]int{11, 12},
	}
	// as duplas chaves [][] significa que tem um slice que contem slices de ints, por exemplo. Necessariamente,
	// dentro da ss tem que ter somente um tipo determinado, nao pode ter int com strings por exemplo

	fmt.Println(ss[2][0]) // slice[slice de posicao 2][pos do elemento do slice de posicao 2]

	fmt.Printf("\n")

	for x := 0; x <= 2; x++ {
		for y := 0; y < len(ss[x]); y++ {
			fmt.Printf("%v ", ss[x][y])
		}
		fmt.Printf("\n")
	} // percorreu todos os slices de "ss"
}
