package lib

import "fmt"

func ComputeSum() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func Caller1() {
	ComputeSum()
}

func Caller2() {
	Caller1()
}

func Entry() {
	fmt.Println("first invocation")
	ComputeSum()
	fmt.Println("second invocation")
	ComputeSum()

	Caller2()
}
