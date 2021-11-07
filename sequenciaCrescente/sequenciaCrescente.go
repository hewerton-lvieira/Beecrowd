package main

import (
	"fmt"
)

func main() {
	var x int
	x = 1
	for x != 0 {
		fmt.Scanln(&x)
		for i := 1; i <= x; i++ {
			if i != x {
				fmt.Printf("%v ", i)
			} else {
				fmt.Printf("%v\n", i)
			}
		}
	}

}
