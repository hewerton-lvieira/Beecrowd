package main

import (
	"fmt"
)

func main() {
	var s1 []int
	s1 = []int{0}
	var x, cont, maior int
	for i := 1; i <= 100; i++ {
		fmt.Scanln(&x)
		s1 = append(s1, x)
	}
	maior = s1[1]
	for i := 1; i <= 100; i++ {
		if maior < s1[i] {
			maior = s1[i]
		}
	}
	for i := 1; s1[i] != maior; i++ {
		cont++
	}
	fmt.Printf("%v\n%v\n", maior, (cont + 1))

}
