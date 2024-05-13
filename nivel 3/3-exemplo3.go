package main

import (
	"fmt"
)

var a interface{}

func main() {
	x := 10

	switch true { //se tal caso for "true", retorna a operação abaixo
	case x < 1:
		fmt.Printf("x e menor que 1 \n")
	case x == 1:
		fmt.Printf("x e igual que 1 \n")
	case x > 1:
		fmt.Printf("x e maior que 1 \n")
	}

	nome := "claudio"

	switch nome {
	case "joao": //se "joao" for igual a nome, entao retorna a operacao a baixo
		fmt.Printf("seu nome é joao \n")
	case "enrique":
		fmt.Printf("seu nome é enrique \n")
	case "claudio":
		fmt.Printf("seu nome é claudio \n")
		fallthrough //fallthrough: se o de cima ta certo, entao o debaixo tambem vai estar
	case "lindo":
		fmt.Printf("ele e lindo \n")

	default:
		//é o padrao quando nao tem outra alternativa

	}

	a = true

	switch a.(type) {
	case int:
		fmt.Printf("é inteiro\n")
	case float64:
		fmt.Printf("é float\n")
	case bool:
		fmt.Printf("é booleano\n")
	case string:
		fmt.Printf("é string\n")
	}

}
