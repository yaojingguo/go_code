package main

import "context"
import "fmt"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	Work(ctx)
	fmt.Printf("ctx: %q\n", ctx)
	//if false {
	//  cancel()
	//}
}

func Work(ctx context.Context) {
}
