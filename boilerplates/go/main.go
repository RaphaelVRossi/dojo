package main

import (
	"fmt"
	"math"
)

// https://judge.beecrowd.com/pt/problems/view/1018
func main() {
	entrada := 0.0

	fmt.Scanln(&entrada)

	fmt.Printf("%.0f\n", entrada)

	notas := []float64{100, 50, 20, 10, 5, 2, 1}

	for _, nota := range notas {
		sobra := calc(entrada, nota)
		entrada -= sobra * nota
	}
}

func calc(entrada float64, nota float64) float64 {
	sobra := math.Floor(entrada / nota)

	fmt.Printf("%.0f nota(s) de R$ %.0f,00\n", sobra, nota)

	return sobra
}
