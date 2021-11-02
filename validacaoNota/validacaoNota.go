package main

import (
	"fmt"
)

func main() {
	var nota1, nota2 float64
	fmt.Scanln(&nota1)
	if nota1 > 10 || nota1 < 0 {
		for nota1 > 10 || nota1 < 0 {
			fmt.Println("nota invalida")
			fmt.Scanln(&nota1)
		}
		fmt.Scanln(&nota2)
		if nota2 > 10 || nota2 < 0 {
			for nota2 > 10 || nota2 < 0 {
				fmt.Println("nota invalida")
				fmt.Scanln(&nota2)
			}
		}
		media := (nota1 + nota2) / 2
		fmt.Println("media =", media)
	}
}
