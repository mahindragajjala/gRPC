package main

import (
    "context"
    "log"
    "time"

    pb "yourmodule/proto" // Replace with actual import path

    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewHelloServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Mahindra"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", resp.GetMessage())
}
