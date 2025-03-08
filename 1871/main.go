package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://judge.beecrowd.com/pt/problems/view/1018
func main() {
	var esq, dir, soma int

	for {
		fmt.Scanf("%d %d", &esq, &dir)
		if esq == 0 && dir == 0 {
			break
		}

		soma = esq + dir

		result := strconv.Itoa(soma)
		result = strings.ReplaceAll(result, "0", "")
		fmt.Printf("%s\n", result)
	}
}
