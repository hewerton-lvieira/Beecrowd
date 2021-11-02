package main

import (
	"fmt"
)

func main() {
	var nota1, nota2 float64
	fmt.Scanln(&nota1)
	for nota1 < 0 || nota1 > 10 {
		fmt.Println("nota invalida")
		fmt.Scanln(&nota1)
	}
	fmt.Scanln(&nota2)
	for nota2 < 0 || nota2 > 10 {
		fmt.Println("nota invalida")
		fmt.Scanln(&nota2)
	}
	media := (nota1 + nota2) / 2
	fmt.Printf("media = %.2f\n", media)
	var decisao int
	fmt.Println("novo calculo (1-sim 2-nao)")
	fmt.Scanln(&decisao)
	for decisao != 2 {
		if decisao == 1 {
			fmt.Scanln(&nota1)
			for nota1 < 0 || nota1 > 10 {
				fmt.Println("nota invalida")
				fmt.Scanln(&nota1)
			}
			fmt.Scanln(&nota2)
			for nota2 < 0 || nota2 > 10 {
				fmt.Println("nota invalida")
				fmt.Scanln(&nota2)
			}
			media := (nota1 + nota2) / 2
			fmt.Printf("media = %.2f\n", media)
			fmt.Println("novo calculo (1-sim 2-nao)")
			fmt.Scanln(&decisao)
		} else if decisao != 1 && decisao != 2 {
			fmt.Println("novo calculo (1-sim 2-nao)")
			fmt.Scanln(&decisao)
		}
	}
}
