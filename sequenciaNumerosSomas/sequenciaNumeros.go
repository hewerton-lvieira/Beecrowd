package main

import (
	"fmt"
)

func main() {
	var m, n, somatorio int
	var s1 []int

	for i := 1; i >= 0; i++ {
		fmt.Scanf("%v%v\n", &m, &n)
		if m <= 0 || n <= 0 {
			i = -10
		} else {
			s1 = append(s1, m)
			s1 = append(s1, n)
		}
	}
	for i := 0; i < (len(s1)); i = i + 2 {
		var s2 []int
		if s1[i] < s1[i+1] {
			for k := s1[i]; k <= s1[i+1]; k++ {
				s2 = append(s2, k)
			}
		} else {
			for k := s1[i+1]; k <= s1[i]; k++ {
				s2 = append(s2, k)
			}
		}
		for i := 0; i < len(s2); i++ {
			somatorio += s2[i]
		}
		for i := 0; i < len(s2); i++ {
			fmt.Printf("%v ", s2[i])
		}
		fmt.Printf("Sum=%v\n", somatorio)
		somatorio = 0
	}
}
