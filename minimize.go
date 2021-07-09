package main

import (
	"fmt"
)

func compute(partition []float64) float64 {
	var cost, n, i float64
	for j, x := range partition {
		n = n + x
		i = float64(j + 1)
		cost = cost + x*x/2 + i*x
	}
	cost = cost - 3*n/2
	return cost
}

func main() {
	a1 := []float64{3, 3, 3}
	a2 := []float64{4, 3, 2}

	fmt.Println("placing 3 express stations among 9 local stations")
	fmt.Printf("  uniform placement: %f\n", compute(a1))
	fmt.Printf("  optimal placement: %f\n", compute(a2))

	b1 := []float64{5, 5, 5, 5, 5}
	b2 := []float64{7, 6, 4, 4, 3}
	fmt.Println("placing 5 express stations among 25 local stations")
	fmt.Printf("  uniform placement: %f\n", compute(b1))
	fmt.Printf("  optimal placement: %f\n", compute(b2))
}
