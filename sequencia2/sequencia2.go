package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 9; i += 2 {
		j := 7
		fmt.Printf("I=%v J=%v\n", i, j)
		j -= 1
		fmt.Printf("I=%v J=%v\n", i, j)
		j -= 1
		fmt.Printf("I=%v J=%v\n", i, j)
	}
}
