package main

import (
	"fmt"
)

func fibonacciii(valor int) {
	var fn, f1, f2, cont int
	f1 = 0
	f2 = 1
	cont = 2
	fmt.Printf("%v", f1)
	fmt.Printf(" %v", f2)
	fn = f1 + f2
	f1 = f2
	f2 = fn
	for cont < valor {
		cont++
		fmt.Printf(" %v", fn)
		fn = f1 + f2
		f1 = f2
		f2 = fn
	}
	fmt.Printf("\n")
}
func main() {
	var x int
	fmt.Scanln(&x)
	fibonacciii(x)
}
