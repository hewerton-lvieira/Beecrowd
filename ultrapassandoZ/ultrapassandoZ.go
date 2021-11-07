package main

import (
	"fmt"
)

func main() {
	var x, z, i, somatorio int
	fmt.Scanln(&x, &z)
	for z <= x {
		fmt.Scanln(&z)
	}
	for somatorio < z {
		somatorio += x + i
		i++
	}
	fmt.Println(i)

}
