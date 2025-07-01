fmt.Println("\n-- Streaming Users (via StreamAllUsers) --")
stream, err := client.StreamAllUsers(context.Background(), &usermgmtpb.StreamAllUsersRequest{})
if err != nil {
    log.Fatalf("Error on stream: %v", err)
}

for {
    user, err := stream.Recv()
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatalf("Stream recv error: %v", err)
    }

    log.Printf("Streamed User: %s (%s)", user.GetId(), user.GetName())
}
