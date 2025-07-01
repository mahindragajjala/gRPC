package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "sync"

    "google.golang.org/grpc"
    "usermgmtpb"
)

type userServer struct {
    usermgmtpb.UnimplementedUserServiceServer
    mu    sync.Mutex
    users map[string]*usermgmtpb.User
    count int
}

func newUserServer() *userServer {
    return &userServer{
        users: make(map[string]*usermgmtpb.User),
    }
}

func (s *userServer) CreateUser(ctx context.Context, req *usermgmtpb.CreateUserRequest) (*usermgmtpb.CreateUserResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.count++
    id := fmt.Sprintf("user-%d", s.count)
    req.User.Id = id
    s.users[id] = req.User

    log.Printf("Created user: %s\n", id)
    return &usermgmtpb.CreateUserResponse{Id: id}, nil
}

func (s *userServer) GetUser(ctx context.Context, req *usermgmtpb.GetUserRequest) (*usermgmtpb.GetUserResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    user, exists := s.users[req.Id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }

    return &usermgmtpb.GetUserResponse{User: user}, nil
}

func (s *userServer) ListUsers(ctx context.Context, req *usermgmtpb.ListUsersRequest) (*usermgmtpb.ListUsersResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    var userList []*usermgmtpb.User
    for _, user := range s.users {
        userList = append(userList, user)
    }

    return &usermgmtpb.ListUsersResponse{Users: userList}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    usermgmtpb.RegisterUserServiceServer(grpcServer, newUserServer())

    log.Println("Server started at :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
