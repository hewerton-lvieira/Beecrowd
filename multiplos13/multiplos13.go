package main

import (
	"fmt"
)

func main() {
	var x, y, somatorio int
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	if x <= y {
		for i := x; i <= y; i++ {
			if i%13 != 0 {
				somatorio = somatorio + i
			}
		}
	} else if y < x {
		for i := y; i <= x; i++ {
			if i%13 != 0 {
				somatorio = somatorio + i
			}
		}
	}
	fmt.Println(somatorio)
}
