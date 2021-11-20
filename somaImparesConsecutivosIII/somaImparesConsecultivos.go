package main

import (
	"fmt"
)

func main() {
	var n, x, y int
	fmt.Scanln(&n)
	var s1 []int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d%d\n", &x, &y)
		somatorio := 0

		for f := 1; f <= y; f++ {
			if x%2 != 0 {
				somatorio += x
				f++
			}
			f--
			x++
		}
		s1 = append(s1, somatorio)
	}
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}
}
