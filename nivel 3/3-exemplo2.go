package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue //para a execucao atual e vai pra proxima

			//break - literalmente quebra o laço de repetição

		}
		fmt.Println(i)
	}
}
