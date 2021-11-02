package main

import (
	"fmt"
)

func main() {
	var n, x, y int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v%v\n", &x, &y)
		var divisao float64 = float64(x) / float64(y)
		if y == 0 {
			fmt.Println("divisao impossivel")
		} else {
			fmt.Printf("%.1f\n", divisao)
		}
	}
}
