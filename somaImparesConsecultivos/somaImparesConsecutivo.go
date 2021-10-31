package main

import (
	"fmt"
)

func main() {
	var n, a, b, somatorio int
	fmt.Scanln(&n)
	var s1 []int
	var s2 []int
	for j := 0; j < n; j++ {
		somatorio = 0
		fmt.Scanf("%v%v\n", &a, &b)
		if a >= b {
			for i := b + 1; i < a; i++ {
				if i%2 != 0 {
					s1 = append(s1, i)
				}
			}
		} else {
			for i := a + 1; i < b; i++ {
				if i%2 != 0 {
					s1 = append(s1, i)
				}
			}
		}
		for h := 0; h < len(s1); h++ {
			somatorio += s1[h]
		}
		s2 = append(s2, somatorio)
		for h := 0; h < len(s1); h++ {
			s1[h] = 0
		}
	}
	for x := 0; x < (len(s2)); x++ {
		fmt.Println(s2[x])
	}
}
