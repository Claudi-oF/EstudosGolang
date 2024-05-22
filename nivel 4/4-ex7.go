package main

import (
	"fmt"
)

func main() {
	x := [][]string{
		[]string{"Claudio", "Felipe", "Augusto"},
		[]string{"Filho", "Eduardo", "Nogueira"},
		[]string{"Violao", "Natacao", "Desenhar"},
	}

	for i := 0; i < 3; i++ {
		for y := 0; y < 3; y++ {
			fmt.Printf("%v ", x[y][i])
		}
		fmt.Printf("\n")
	}
}
