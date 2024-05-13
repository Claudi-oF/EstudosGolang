package main

import (
	"fmt"
)

func main() {
	qualquercoisa := map[int]string{
		123: "top",
		456: "meio top",
		789: "menos top",
	}

	fmt.Println(qualquercoisa)

	for i, v := range qualquercoisa {
		fmt.Println(i, v) // maps nao fica em ordem!!!
	}
}
