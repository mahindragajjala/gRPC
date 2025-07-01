user_grpc.pb.go

Purpose:
This file contains:
                The gRPC client and server interface code
                Server registration helpers
                Client stubs that the client uses to make RPCs
                Interface that the server must implement

Server Interface 
          type UserServiceServer interface {
              GetUser(context.Context, *UserRequest) (*UserResponse, error)
          }


Server Registration:
          func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer)


Client Interface:
            type UserServiceClient interface {
                GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
            }
On Server:
          userpb.RegisterUserServiceServer(grpcServer, &server{})
On Client:
          client := userpb.NewUserServiceClient(conn)
          response, err := client.GetUser(ctx, req)
