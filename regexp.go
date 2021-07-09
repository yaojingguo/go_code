package main

import "fmt"
import "regexp"

func main() {
	pattern := "(foo|foobar)"
	s := "foobarbaz"
	byteSlice := []byte(s)

	rPCRE, _ := regexp.Compile(pattern)
	rPOSIX, _ := regexp.CompilePOSIX(pattern)

	fmt.Println(string(rPCRE.Find(byteSlice)))
	fmt.Println(rPCRE.FindString(s))
	fmt.Println(rPCRE.FindStringSubmatch(s))

	fmt.Println(string(rPOSIX.Find(byteSlice)))
	fmt.Println(rPOSIX.FindString(s))
	fmt.Println(rPOSIX.FindStringSubmatch(s))
}
