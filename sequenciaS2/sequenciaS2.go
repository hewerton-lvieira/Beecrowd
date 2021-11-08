package main

import (
	"fmt"
)

func main() {
	var s, d float64
	s = 0.0
	d = 1.0
	for i := 1; i <= 39; i = i + 2 {
		if i <= 39 {
			s = s + (float64(i) / d)
			d = d * 2
		}
	}
	fmt.Printf("%.2f\n", s)
}
