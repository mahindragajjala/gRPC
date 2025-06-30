// client.go
package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	args := Args{A: 6, B: 7}
	var result Result

	err = client.Call("Calculator.Multiply", args, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Multiplication Result: %d × %d = %d\n", args.A, args.B, result.Product)
}
