package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    "usermgmtpb"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    client := usermgmtpb.NewUserServiceClient(conn)

    // Create a new user
    user := &usermgmtpb.User{
        Name:  "Mahindra",
        Age:   30,
        Email: "mahindra@example.com",
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    createResp, err := client.CreateUser(ctx, &usermgmtpb.CreateUserRequest{User: user})
    if err != nil {
        log.Fatalf("CreateUser failed: %v", err)
    }
    log.Printf("User created with ID: %s", createResp.Id)

    // Fetch the created user
    getResp, err := client.GetUser(ctx, &usermgmtpb.GetUserRequest{Id: createResp.Id})
    if err != nil {
        log.Fatalf("GetUser failed: %v", err)
    }
    log.Printf("User details: %+v", getResp.User)

    // List all users
    listResp, err := client.ListUsers(ctx, &usermgmtpb.ListUsersRequest{})
    if err != nil {
        log.Fatalf("ListUsers failed: %v", err)
    }

    log.Println("User list:")
    for _, u := range listResp.Users {
        log.Printf("- %s: %s (%d)", u.Id, u.Name, u.Age)
    }
}
