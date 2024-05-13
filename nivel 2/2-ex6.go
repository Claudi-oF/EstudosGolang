package main

import (
	"fmt"
)

const (
	a = iota + 2025
	b
	c
	d

	//com iota pode declarar os proximos numeros ou nao se quiser pra funcionar:
	//a = iota + 2025
	//b = iota + 2025
	//c = iota + 2025
	//d = iota + 2025
)

func main() {
	fmt.Println(a, b, c, d)
}
