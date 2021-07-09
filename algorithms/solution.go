package algorithms

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := 0
	xCopy := x
	for ; x > 0; {
		y = 10 * y + x % 10
		x = x / 10
	}

	fmt.Printf("x: %d, y: %d", x, y)


	return xCopy == y
}