package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		valor1 := i
		valor2 := math.Pow(float64(i), 2)
		valor3 := math.Pow(float64(i), 3)
		fmt.Printf("%v %v %v\n", int(valor1), int(valor2), int(valor3))
	}
}
