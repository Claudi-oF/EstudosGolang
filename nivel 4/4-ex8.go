package main

import "fmt"

func main() {
	amigo := map[string]string{
		"claudio":      "tocar violao",
		"lebron james": "jogar basquete",
		"neymar":       "jogar bola",
	}

	amigo = map[string]string{
		"amanda": "fazer comida",
	}

	for i, v := range amigo {
		fmt.Printf("Indice: %v, Valor: %v\n", i, v)
	}
}
