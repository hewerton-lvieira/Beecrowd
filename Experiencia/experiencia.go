package main

import (
	"fmt"
)

func main() {
	var quantia, n, coelho, sapo, rato int
	var tipo string
	fmt.Scanln(&n)
	s1 := []int{0}
	s2 := []string{"0"}
	for i := 1; i <= n; i++ {
		fmt.Scanf("%v%v\n", &quantia, &tipo)
		s1 = append(s1, quantia)
		s2 = append(s2, tipo)
	}
	for i := 1; i <= n; i++ {
		switch {
		case s2[i] == "C":
			coelho = coelho + s1[i]
		case s2[i] == "S":
			sapo = sapo + s1[i]
		case s2[i] == "R":
			rato = rato + s1[i]
		}
	}
	s := "%"
	total := coelho + rato + sapo
	fmt.Printf("Total: %v cobaias\n", (total))
	fmt.Printf("Total de coelhos: %v\n", (coelho))
	fmt.Printf("Total de ratos: %v\n", (rato))
	fmt.Printf("Total de sapos: %v\n", (sapo))
	fmt.Printf("Percentual de coelhos: %.2f %s\n", (((float64(coelho)) * 100.00) / float64(total)), s)
	fmt.Printf("Percentual de ratos: %.2f %s\n", ((float64(rato))*100.00)/(float64(total)), s)
	fmt.Printf("Percentual de sapos: %.2f %s\n", ((float64(sapo))*100.00)/(float64(total)), s)

}
