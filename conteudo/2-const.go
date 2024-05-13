package main

import (
	"fmt"
)

// o "const" nao define uma variavel DE CARA igual "var", "const a" pode ser qualquer coisa (int, float64...) ANTES de ser usada.
// depois de "a" ser usada, dai ela tem que ser definida.
const a = 10

var b = 10

var x float64 = 1.5
var y int8 = 40

func main() {
	fmt.Println(x)
	fmt.Printf("%T \n", a)

	x = a

	fmt.Println(x)
}
