package main

import (
	"fmt"
)

func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	if y == 1 {
		return x
	}
	i := y / 2
	f := y % 2
	p := pow(x, i)
	p *= p
	if f == 1 {
		p *= x
	}
	return p
}

func main() {
	fmt.Printf("Pow(2, 5): %v\n", pow(2, 5))
}
