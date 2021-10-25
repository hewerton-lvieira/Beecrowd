package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scanln(&n)
	var a1 []int
	a1 = []int{0}
	for i := 1; i <= n; i++ {
		fmt.Scanln(&x)
		a1 = append(a1, x)
	}
	for i := 1; i <= n; i++ {
		if a1[i] == 0 {
			fmt.Println("NULL")
		} else if a1[i]%2 == 0 {
			fmt.Print("EVEN ")
		} else {
			fmt.Print("ODD ")
		}
		if a1[i] > 0 {
			fmt.Println("POSITIVE")
		} else if a1[i] == 0 {

		} else {
			fmt.Println("NEGATIVE")
		}
	}
}
