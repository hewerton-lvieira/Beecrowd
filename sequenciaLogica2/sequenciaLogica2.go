package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	for i := 1; i <= y; i++ {
		j := 1
		for j <= x {
			for j <= x {
				if j < x {
					fmt.Printf("%v ", i)
				} else {
					fmt.Printf("%v", i)
				}
				j++
				i++
			}

		}
		i--
		fmt.Printf("\n")
	}
}
