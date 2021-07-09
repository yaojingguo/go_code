package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	err := sayHi()
	if err != nil {
		panic(err)
	}
}

// sayHi() returns the string "hi"
func sayHi() error {
	fmt.Println("hi")
	return nil
}
