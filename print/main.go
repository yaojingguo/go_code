// +build linux

package main

import (
	"encoding/json"
	"fmt"
)

type foo struct {
	Message    string `json:"message"`
	Ports      []int
	ServerName string `json:"server_name"`
}

func newFoo() (*foo, error) {
	return &foo{
		Message:    "foo loves bar",
		Ports:      []int{80},
		ServerName: "Foo",
	}, nil
}

func main() {
	res, err := newFoo()
	if err != nil {
		panic("new Foo")
	}

	out, err := json.Marshal(res)
	if err != nil {
		panic("Marshal")
	}
	fmt.Printf("string(out) = %+v\n", string(out))

	f := foo{quz: "QUZ"}
	q.quz
}

type foo struct {
	quz string
}
