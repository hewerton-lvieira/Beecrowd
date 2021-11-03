package main

import (
	"fmt"
)

func main() {
	var contAl, contGa, contDi, i int
	for i != 4 {
		fmt.Scanln(&i)
		if i == 1 {
			contAl++
		} else if i == 2 {
			contGa++
		} else if i == 3 {
			contDi++
		} else if i == 4 {

		} else {

		}

	}
	fmt.Println("MUITO OBRIGADO")
	fmt.Println("Alcool:", contAl)
	fmt.Println("Gasolina:", contGa)
	fmt.Println("Diesel:", contDi)
}
