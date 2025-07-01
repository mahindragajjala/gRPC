package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "google.golang.org/grpc"
    "grpc-demo/proto/userpb"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    c := userpb.NewUserServiceClient(conn)

    req := &userpb.UserRequest{
        UserId: "U001",
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    res, err := c.GetUser(ctx, req)
    if err != nil {
        log.Fatalf("Error when calling GetUser: %v", err)
    }

    fmt.Printf("Response from Server: Name=%s, Age=%d, Email=%s\n",
        res.GetName(), res.GetAge(), res.GetEmail())
}
