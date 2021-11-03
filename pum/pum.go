package main

import (
	"fmt"
)

func main() {
	var n, cont int
	cont = 1
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		cont1 := cont
		cont2 := cont + 1
		cont3 := cont + 2
		fmt.Printf("%v %v %v PUM\n", cont1, cont2, cont3)
		cont = cont + 4
	}
}
