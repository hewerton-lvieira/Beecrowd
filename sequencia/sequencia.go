package main

import (
	"fmt"
)

func main() {
	i := 1
	for j := 60; j >= 0; j = j - 5 {
		fmt.Printf("I=%v J=%v\n", i, j)
		i += 3
	}
}
