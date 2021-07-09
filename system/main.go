package main

import (
	"github.com/yaojingguo/system/util"
)

func main() {
	util.Info()
	numbers := []int{1, 2, 3}
	target := make([], 10, 10)
	target = append(target, numbers...)
}
