package main

import (
  "fmt"
  "time"
  "math/rand"
)

var (
  Web1 = fakeSearch("web1")
  Web2 = fakeSearch("web2")
  Image1 = fakeSearch("image1")
  Image2 = fakeSearch("image2")
  Video1 = fakeSearch("video1")
  Video2 = fakeSearch("video2")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
  return func(query string) Result {
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    return Result(fmt.Sprintf("%s result for %q\n", kind, query))
  }
}

func main() {
  var results []Result
  c := make(chan Result)
  query := "golang"
  go func() { c <- First(query, Web1, Web2) } ()
  go func() { c <- First(query, Image1, Image2) } ()
  go func() { c <- First(query, Video1, Video2) } ()
  timeout := time.After(200 * time.Millisecond)
  for i := 0; i < 3; i++ {
    select {
    case result := <-c:
      results = append(results, result)
    case <-timeout:
      fmt.Println("timed out")
      return
    }
  }
  fmt.Println(results)
  return
}

func First(query string, replicas ...Search) Result {
  c := make(chan Result)
  searchReplica := func(i int) { c <- replicas[i](query) }
  for i := range replicas {
    go searchReplica(i)
  }
  return <-c
}


