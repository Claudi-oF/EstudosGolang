package main

import (
	"fmt"
)

type banana int

var x banana

var y int

func main() {
	fmt.Println("Valor de x: ", x)
	fmt.Printf("Tipo de x: %T\n", x)

	y = int(x)

	fmt.Println("Valor de y: ", y)
	fmt.Printf("Tipo de y: %T", y)

}
