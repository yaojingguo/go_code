package main

import (
  "fmt"
)

const (
	First int64 = iota
	Second
)

const (
  One int64 = 1
  Two = 2
)

const (
  A, B = iota, iota
  C, D
)

const (
  X int64 = 10
  Y
)
type ByteSize float64

func main() {
  fmt.Printf("X type: %T, Y type: %T\n", X, Y)
  fmt.Printf("A: %v, B: %v, C: %v, D: %v\n", A, B, C, D)
  fmt.Printf("First type: %T, Second type: %T\n", First, Second)
  fmt.Printf("One type: %T, Two type: %T\n", One, Two)
}
