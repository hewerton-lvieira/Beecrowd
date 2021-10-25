package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var x, y, z, mediaPonderada float64
	var s1 []float64
	s1 = []float64{0}
	for i := 0; i < n; i++ {
		fmt.Scanf("%v%v%v\n", &x, &y, &z)
		s1 = append(s1, x)
		s1 = append(s1, y)
		s1 = append(s1, z)
	}
	for i := 3; i <= len(s1)-1; i = i + 3 {
		mediaPonderada = ((s1[i-2] * 2) + (s1[i-1] * 3) + (s1[i] * 5)) / 10
		fmt.Printf("%.1f\n", mediaPonderada)
	}

}
