package main

import (
	"fmt"
)

func main() {

	nasci := 2005
	parar := 2024

	for {
		if nasci == parar+1 {
			break
		}
		fmt.Println(nasci)
		nasci++

	}
}
