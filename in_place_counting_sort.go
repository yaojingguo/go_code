// Solution to ALGS4 exercise 5.1.17
package main

import (
	"fmt"
)

const (
	BITS_PER_BYTE = 8
	R             = 1 << BITS_PER_BYTE
)

// In-place key-indexed counting is stable.
func Sort(ss []string) {
	N := len(ss)
	count := make([]int, R+1)

	// Compute frequency counts
	for _, s := range ss {
		count[int(s[0])+1]++
	}

	// Compute cumulates
	for r := 0; r < R; r++ {
		count[r+1] = count[r+1] + count[r]
	}

	// Move data in-place
	i, r, moved := 0, 0, 0
	for r < R {
		if count[r] != 0 {
			r--
			break
		}
		r++
	}
	for moved < N {
		key := int(ss[i][0])
		fmt.Printf("i: %d, key: %x, ", i, key)
		if key == r {
			i++
			count[r]++
			if count[r] == count[r+1] {
				r++
			}
			fmt.Printf("increase i to %d\n", i)
		} else {
			j := count[key]
			ss[i], ss[j] = ss[j], ss[i]
			count[key] = j + 1
			fmt.Printf("swap %d with %d\n", i, j)
		}
		moved++
		fmt.Printf("ss: %v\n", ss)
	}
}

func main() {
	ss := []string{"ax", "b-", "ay", "b+", "az"}
	fmt.Printf("input: %v\n", ss)
	fmt.Printf("======================================================\n")
	Sort(ss)
	fmt.Printf("======================================================\n")
	fmt.Printf("output: %v\n", ss)
}
