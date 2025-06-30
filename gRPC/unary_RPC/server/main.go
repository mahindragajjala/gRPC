package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "grpc-hello-world/proto" // import the generated proto package
)

// Implement the server
type HelloServer struct {
    proto.UnimplementedHelloServiceServer
}

// Implement the SayHello method
func (s *HelloServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
    message := fmt.Sprintf("Hello, %s!", req.GetName())
    return &proto.HelloResponse{Message: message}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterHelloServiceServer(grpcServer, &HelloServer{})

    fmt.Println("gRPC Server is running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
