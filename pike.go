package main

import (
  "fmt"
  "time"
  "math/rand"
)

type Message struct {
  str string
  wait chan bool}


func main() {
  c := fanIn(boring("Joe"), boring("Ann"))
  for i := 0; i < 5; i++ {
    msg1 := <-c; fmt.Println(msg1.str)
    msg2 := <-c; fmt.Println(msg2.str)
    msg1.wait <- true
    msg2.wait <- true
  }
  fmt.Println("You're boring; I'm leaving.")
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
  c := make(chan Message)
  go func() {
    for {
      select {
      case s := <-input1:
        c <- s
      case s:= <-input2:
        c <- s
      }
    }
  }()
  return c
}

func boring(msg string) <-chan Message {
  waitForIt := make(chan bool)
  c := make(chan Message)
  go func() {
    for i:=0; ; i++ {
      c <- Message{ fmt.Sprintf("%s %d", msg, i), waitForIt }
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
      <-waitForIt
    }
  }()
  return c
}
