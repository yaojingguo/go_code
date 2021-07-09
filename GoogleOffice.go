package main


import "fmt"

type Office int

const (
  Boston Office = iota
  NewYork
)

var officePlace = map[Office]string {
  Boston: "Boston",
  NewYork: "New York",
}

func (o Office) String() string {
  return "Google, " + officePlace[o]
}

func main() {
  fmt.Printf("Hello, %s\n", Boston)
}
