package main

import (
    "context"
    "fmt"
    "log"
    "net"

    pb "yourmodule/proto" // Replace with actual import path

    "google.golang.org/grpc"
)

type HelloServer struct {
    pb.UnimplementedHelloServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
    log.Printf("Received: %v", req.GetName())
    return &pb.HelloResponse{
        Message: "Hello, " + req.GetName(),
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterHelloServiceServer(grpcServer, &HelloServer{})

    log.Println("Server listening on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
