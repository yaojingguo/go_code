package main

import (
  "fmt"
  "hash/crc32"
)

func main() {
  h := crc32.NewIEEE()
  fmt.Fprintf(h, "Hello, 世界\n")
  fmt.Printf("hash=%#x\n", h.Sum32())
}
