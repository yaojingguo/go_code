package main

import (
	"fmt"
)

type Greeter interface {
	Hello(name string)
}
type HandlerFunc func(name string)

func (f HandlerFunc) Hello(name string) {
	f(name)
}

type Yao struct{}

func (y Yao) Hello(name string) {
	fmt.Printf("hello, %s\n", name)
}

func invoke(h Greeter, name string) {
	h.Hello(name)
}

func hi(name string) {
	fmt.Printf("hi, %s\n", name)
}

func main() {
	y := Yao{}
	invoke(y, "xiaoyu")

	invoke(HandlerFunc(hi), "xiaoyu")
}
