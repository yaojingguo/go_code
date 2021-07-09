package main

import (
  "log"
  "os"
  "io"
  "encoding/xml"
  "encoding/json"
  "fmt"
)

type Lang struct {
  Name string
  Year int
  URL string
}

func do(f func(Lang)) {
  input, err := os.Open("lang.json")
  if err != nil {
    log.Fatal(err)
  }
  dec := json.NewDecoder(input)
  for {
    var lang Lang
    err := dec.Decode(&lang)
    if err != nil {
      if err == io.EOF {
        break
      }
      log.Fatal(err)
    }
    f(lang)
  }
}

func main() {
  do(func(lang Lang) {
    data, err := xml.MarshalIndent(lang, "", "  ")
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%s\n", data)
  })
}
