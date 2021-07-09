package main

import (
  "encoding/hex"
  "fmt"
  "os"
)

func f(args ...int) {
  for i, v := range args {
    fmt.Printf("%v: %v\n", i, v)
  }
}

func main() {
  h := hex.Dumper(os.Stdout)
  defer h.Close()
  fmt.Fprintf(h, "Hello, 世界\n")
}
