package main

import (
	"fmt"
)

func main() {
	var a, n, soma int
	fmt.Scanln(&a, &n)
	for n <= 0 {
		fmt.Scanln(&n)
	}
	for i := 0; i < n; i++ {
		soma += a + i
	}
	fmt.Println(soma)
}
