Client                     Server
  │                          │
  │ --- SayHello(name) --->  │
  │                          │
  │ <-- "Hello, name!" ----  │
  │                          │




How it Works
              The .proto file defines HelloService and messages.
              protoc generates Go code (gRPC interfaces & types).
              Server implements the SayHello() method.
              Client calls SayHello() using a gRPC stub.
              Server receives the name and returns a greeting.


syntax = "proto3";

package hello;

option go_package = "grpc-hello-world/proto;proto";

// Define the service
service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

// Request message
message HelloRequest {
  string name = 1;
}

// Response message
message HelloResponse {
  string message = 1;
}


[ .proto file ]
     │
     ▼
[ protoc compiler ]
     │
     ├── --go_out → generates message types
     └── --go-grpc_out → generates gRPC service stubs
     ▼
[ Go code files ]
  ├── hello.pb.go         ← message types & helper methods
  └── hello_grpc.pb.go    ← gRPC service interface (client/server)





Client Side                                            Server Side
──────────────────────────────────────────────────────────────────────────────────

main.go (client)
  │
  │ grpc.Dial("localhost:50051") ────────────────▶ Establish TCP connection
  │                                                 and create gRPC client
  │
  │ NewHelloServiceClient(conn)
  │     │
  │     └─▶ Returns *helloServiceClient
  │
  │ client.SayHello(ctx, &HelloRequest{Name: "Mahindra"})
  │     │
  │     └─▶ helloServiceClient.SayHello()
  │             │
  │             └─▶ grpc.Invoke(...) 
  │                    │
  │                    └─▶ serializes HelloRequest using protobuf
  │                    │
  │                    └─▶ sends bytes via HTTP/2 request stream
  │
  │                                       ┌───────────────────────────────────┐
  ▼                                       ▼                                   │
                                         main.go (server)                    │
                                  ┌────> grpcServer.Serve(lis)              │
                                  │       │                                 │
                                  │       └─▶ Waits for incoming connection │
                                  │                                         │
                                  └──────── TCP received + HTTP/2 frame     │
                                          containing serialized request     │
                                          for "SayHello" RPC                │
                                                 │
                                                 ▼
                                    grpc-go internals parse:
                                     → method name = "SayHello"
                                     → decode HelloRequest
                                     → find registered service
                                                 │
                                                 ▼
                            HelloServer.SayHello(ctx, req *HelloRequest)
                                                 │
                            ┌────────────────────┘
                            │  Build HelloResponse
                            │    → "Hello, Mahindra!"
                            ▼
                  return &HelloResponse{Message: "..."} to gRPC runtime
                                                 │
                                 ┌───────────────┘
                                 │
                      gRPC Server serializes response
                      sends HTTP/2 frame back to client
                                 │
                                 ▼
main.go (client)
  ┌─────────────────────────────────────────────────────────────┐
  │ Receives HTTP/2 frame, deserializes HelloResponse from bytes│
  │ → res.GetMessage() = "Hello, Mahindra!"                     │
  └─────────────────────────────────────────────────────────────┘

