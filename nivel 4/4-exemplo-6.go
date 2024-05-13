package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"claudio": 12377983,
		"jo":      14359286,
	}

	fmt.Println(amigos)            // todos os amigos
	fmt.Println(amigos["claudio"]) // somente o "claudio"

	amigos["pedro"] = 992690301 // "pedro" recebe seu numero

	fmt.Println(amigos)
	fmt.Println(amigos["pedro"])

	fmt.Println(amigos["robertinho"], "\n\n")
	// nao existe esse nome, entao aparecera o valor 0.
	// isso significa que o numero de "robertinho" eh 0 ou significa que nao existe?

	sera, ok := amigos["fantasma"]
	fmt.Println(sera, ok)
	// quando o bool "ok" da false, entao o valor "fantasma" nao existe no mapa!!!
	// isso tira a nossa duvida do robertinho =D
}
