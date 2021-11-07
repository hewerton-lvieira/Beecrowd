package main

import (
	"fmt"
)

func main() {
	var n, n1 int
	fmt.Scanln(&n)
	n1 = n - 1
	for n1 != 0 {
		n = n * n1
		n1--
	}
	fmt.Println(n)
}
