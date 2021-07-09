
package main

import (
  "fmt"
  "code.google.com/p/go-tour/tree"
)

type Tree struct {
  Left *Tree
  Value int
  Right *Tree
}

func WalkImpl(t *tree.Tree, ch chan int) {
  if t == nil {
    return
  }
  WalkImpl(t.Left, ch)
  ch<-t.Value
  WalkImpl(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  WalkImpl(t, ch)
  close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  for {
    v1, ok1 := <-ch1
    v2, ok2 := <-ch2
    if ok1 == false &&  ok2 == false {
      return true
    }
    if ok1 == false ||  ok2 == false {
      return false
    }
    if v1 != v2 {
      return false
    }
  }
  return true
}

func main() {
  // Test Walk
  fmt.Println("Test Walk")
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for v := range ch {
    fmt.Println(v)
  }

  // Test same
  fmt.Println("Test Same")
  fmt.Println(Same(tree.New(1), tree.New(1)))
  fmt.Println(Same(tree.New(1), tree.New(2)))
}
