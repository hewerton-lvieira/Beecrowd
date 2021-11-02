package main

import (
	"fmt"
)

func main() {
	var time1, time2, decisao, contador, inter, gremio int
	inter = 0
	gremio = 0
	empate := 0
	var s1 []int
	contador = 1
	fmt.Scanf("%v%v\n", &time1, &time2)
	s1 = append(s1, time1)
	s1 = append(s1, time2)
	fmt.Println("Novo grenal (1-sim 2-nao)")
	fmt.Scanln(&decisao)
	for decisao == 1 {
		contador++
		fmt.Scanf("%v%v\n", &time1, &time2)
		s1 = append(s1, time1)
		s1 = append(s1, time2)
		fmt.Println("Novo grenal (1-sim 2-nao)")
		fmt.Scanln(&decisao)
	}
	for i := 0; i < len(s1); i = i + 2 {
		if s1[i] > s1[i+1] {
			inter++
		} else if s1[i] < s1[i+1] {
			gremio++
		} else if s1[i] == s1[i+1] {
			empate++
		}
	}
	fmt.Println(contador, "grenais")
	fmt.Printf("Inter:%v\n", inter)
	fmt.Printf("Gremio:%v\n", gremio)
	fmt.Printf("Empates:%v\n", empate)
	if inter > gremio {
		fmt.Println("Inter venceu mais")
	} else if gremio > inter {
		fmt.Println("Gremio venceu mais")
	} else {
		fmt.Println("Nao houve vencedor")
	}
}
