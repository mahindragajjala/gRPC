package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "google.golang.org/grpc"
    "grpc-hello-world/proto" // import the generated proto package
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    client := proto.NewHelloServiceClient(conn)

    // Make a unary request
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    res, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Mahindra"})
    if err != nil {
        log.Fatalf("Error when calling SayHello: %v", err)
    }

    fmt.Println("Server Response:", res.GetMessage())
}
