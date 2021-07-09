package main

import "fmt"
import "sort"

type Person struct {
	no int
}

type people []Person

func (p people) Less(i, j int) bool {
	return p[i].no <= p[j].no
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Len() int {
	return len(p)
}

func main() {
	slice := []Person{{no: 100},{no: 10}}
	fmt.Println(slice)
	modify(slice)
	fmt.Println(slice)
}

func modify(p people) {
	sort.Sort(p)
}
