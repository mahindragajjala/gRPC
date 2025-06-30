// server.go
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Service
type Calculator int

func (c *Calculator) Multiply(args Args, reply *Result) error {
	reply.Product = args.A * args.B
	return nil
}

func main() {
	calculator := new(Calculator)
	rpc.Register(calculator)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("RPC server listening on port 1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
