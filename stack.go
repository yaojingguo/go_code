package main

import "fmt"
import "runtime"
import "strings"
import "time"

func a() {
  buf := make([]byte, 2<<20)
	buf = buf[:runtime.Stack(buf, true)]
	for _, g := range strings.Split(string(buf), "\n\n") {
		sl := strings.SplitN(g, "\n", 2)
		if len(sl) != 2 {
			continue
		}
    stack := strings.TrimSpace(sl[1])
    fmt.Printf("g: %v\n", g)
    fmt.Printf("stack: %v\n", stack)
	}
}

func b() {
  a()
}

func main() {
  go func() {
    fmt.Println("sleeping...")
    time.Sleep(time.Second * 2)
  }()
  b()
}
