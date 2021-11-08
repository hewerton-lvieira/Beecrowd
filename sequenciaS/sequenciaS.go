package main

import (
	"fmt"
)

func main() {
	var s float64
	s = 0.0
	for i := 1; i <= 100; i++ {
		s = s + (1 / float64(i))
	}
	fmt.Printf("%.2f\n", s)
}
