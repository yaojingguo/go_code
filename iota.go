package main

import (
  "fmt"
)

const ( // iota is reset to 0
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const ( // iota is reset to 0
	a = 1 << iota  // a == 1
	b = 1 << iota  // b == 2
	c = 3          // c == 3  (iota is not used but still incremented)
	d = 1 << iota  // d == 8
)

const ( // iota is reset to 0
	u         = iota * 42  // u == 0     (untyped integer constant)
	v float64 = iota * 42  // v == 42.0  (float64 constant)
	w         = iota * 42  // w == 84    (untyped integer constant)
)

const x = iota  // x == 0  (iota has been reset)
const y = iota  // y == 0  (iota has been reset)

func main() {
  fmt.Printf("v type %T, w type %T\n", v, w)
}
