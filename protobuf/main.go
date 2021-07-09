package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	{
		var val int64 = 1
		bytes := proto.EncodeVarint(uint64(val))
		fmt.Printf("0x%x\n", bytes)
	}
	{
		var val int64 = 300
		bytes := proto.EncodeVarint(uint64(val))
		fmt.Printf("0x%x\n", bytes)
	}
	{
		var val int64 = -1
		bytes := proto.EncodeVarint(uint64(val))
		fmt.Printf("0x%x\n", bytes)
	}
}
