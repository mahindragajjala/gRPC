package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "grpc-demo/proto/userpb"
)

// Server implementation
type server struct {
    userpb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
    fmt.Printf("Received request for UserID: %s\n", req.GetUserId())
    return &userpb.UserResponse{
        Name:  "Mahindra",
        Age:   30,
        Email: "mahindra@example.com",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    userpb.RegisterUserServiceServer(grpcServer, &server{})

    fmt.Println("Server is running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
