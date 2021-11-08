package main

import (
	"fmt"
)

func main() {
	var x, i, somatorio int
	for x >= 0 {
		fmt.Scanln(&x)
		if x >= 0 {
			somatorio += x
			i++
		}
	}
	media := (float64(somatorio)) / (float64(i))
	fmt.Printf("%.2f\n", media)
}
