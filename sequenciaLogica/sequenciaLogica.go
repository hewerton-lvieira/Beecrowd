package main

import (
	"fmt"
)

func main() {
	var n int
	var v1, v2, v3, i, cont1, cont2 int
	v1, v2, v3, i, cont1, cont2 = 1, 0, 0, 0, 0, 0
	logica := 1
	fmt.Scanln(&n)
	for i < n {
		i++
		v1 = i
		v2++
		cont1++
		v3++
		fmt.Printf("%v %v %v\n", int(v1), int(v2), int(v3))
		v1 = i
		v2++
		cont1++
		v3++
		fmt.Printf("%v %v %v\n", int(v1), int(v2), int(v3))
		v2 = v2 + cont1 - 1
		v3 = logica*6 + v3 - 1
		cont2++
		logica = logica + (cont2 + 1)
	}

}
