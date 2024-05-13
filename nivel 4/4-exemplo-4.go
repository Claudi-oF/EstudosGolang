package main

import (
	"fmt"
)

func main() {
	// para otimizar o acrescimo de elementos no array, podemos usar o "make"
	// make([]T, len, cap); len = lenght; cap = capacity

	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3] = 1, 2, 3, 4

	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
	slice = append(slice, 11)

	fmt.Println(slice, len(slice), cap(slice))
	// quando atinge o maximo de capacidade, ele DOBRA o cap().

}
