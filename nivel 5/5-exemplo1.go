package main

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "Joao",
		sobrenome: "Vitor",
		fumante:   true,
	}

	fmt.Println(cliente1)
}
