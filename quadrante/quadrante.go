package main

import (
	"fmt"
)

func main() {
	var x, y int
	var s1 []int
	for i := 0; i >= 0; i++ {
		fmt.Scanf("%v%v\n", &x, &y)
		if x == 0 || y == 0 {
			i = -10
		} else {
			s1 = append(s1, x)
			s1 = append(s1, y)
		}
	}
	for i := 0; i < len(s1); i = i + 2 {
		if s1[i] > 0 && s1[i+1] > 0 {
			fmt.Println("primeiro")
		} else if s1[i] > 0 && s1[i+1] < 0 {
			fmt.Println("quarto")
		} else if s1[i] < 0 && s1[i+1] < 0 {
			fmt.Println("terceiro")
		} else if s1[i] < 0 && s1[i+1] > 0 {
			fmt.Println("segundo")
		}
	}
}
