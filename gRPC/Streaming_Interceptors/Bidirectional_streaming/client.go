func main() {
    conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
    client := chatpb.NewChatServiceClient(conn)

    stream, _ := client.ChatStream(context.Background())

    // Sending and receiving concurrently
    go func() {
        messages := []string{"Hello", "How are you?", "Bye"}
        for _, m := range messages {
            _ = stream.Send(&chatpb.ChatMessage{
                User:      "Mahindra",
                Message:   m,
                Timestamp: time.Now().Format(time.RFC3339),
            })
            time.Sleep(1 * time.Second)
        }
        stream.CloseSend()
    }()

    for {
        resp, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("Receive error: %v", err)
        }
        fmt.Printf("📨 From Server: %s - %s\n", resp.GetUser(), resp.GetMessage())
    }
}
