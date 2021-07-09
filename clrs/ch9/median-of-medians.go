// a fork of https://github.com/ZachOrr/golang-algorithms/blob/master/searching/median-of-medians.go
package main

import (
	"fmt"
	"sort"
)

// medianOfMedians takes a list to be searched,
// the number i-thm smallest element (k), and
// the size of the subarrays (r) - r should be >= 3
func MedianOfMedians(elementList []int, ith int) int {
	n := len(elementList)

	// base case
	if n < 10 {
		sort.Ints(elementList)
		return elementList[ith-1]
	}

	// Step 1
	m := (n + 4) / 5
	medians := make([]int, m)

	// Step 2
	for i := 0; i < m; i++ {
		v := (i * 5) + 5
		var arr []int

		if v >= n {
			arr = make([]int, len(elementList[(i*5):]))
			copy(arr, elementList[(i*5):])
		} else {
			arr = make([]int, 5)
			copy(arr, elementList[(i*5):v])
		}

		sort.Ints(arr)
		medians[i] = arr[len(arr)/2]
	}

	// Step 3
	pivot := MedianOfMedians(medians, (len(medians)+1)/2)

	// Step 4
	var left, right []int
	for i := range elementList {
		if elementList[i] < pivot {
			left = append(left, elementList[i])
		} else if elementList[i] > pivot {
			right = append(right, elementList[i])
		}
	}

	// Step 5
	k := len(left) + 1
	switch {
	case ith == k:
		return pivot
	case ith < k:
		return MedianOfMedians(left, ith)
	default:
		return MedianOfMedians(right, ith-k)
	}
}

func main() {
	arr := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 20}
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	for _, i := range []int{5, 10, 7, 8} {
		val := MedianOfMedians(arr, i)
		expected := sortedArr[i-1]
		if val != expected {
			fmt.Printf("Expected %d, got %d\n", expected, val)
		}
	}
}
