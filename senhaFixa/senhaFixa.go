package main

import (
	"fmt"
)

func main() {
	var senha int
	for i := 0; i >= 0; i++ {
		fmt.Scanln(&senha)
		if senha != 2002 {
			fmt.Println("Senha Invalida")
		} else {
			fmt.Println("Acesso Permitido")
			i = -10
		}
	}
}
