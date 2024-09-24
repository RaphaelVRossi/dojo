package main

import (
	"fmt"
)

// https://judge.beecrowd.com/pt/problems/view/2006
func main() {
	p1, s1, s2, s3, s4, s5 := 0, 0, 0, 0, 0, 0

	fmt.Scanln(&p1)

	count := 0
	fmt.Scanln(&s1, &s2, &s3, &s4, &s5)
	for _, num := range []int{s1, s2, s3, s4, s5} {
		if num == p1 {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}
