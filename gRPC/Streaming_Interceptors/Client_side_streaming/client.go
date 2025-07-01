func main() {
    conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
    client := logpb.NewLogServiceClient(conn)

    stream, _ := client.UploadLogs(context.Background())

    logs := []string{"Init started", "DB connected", "Service ready"}
    for _, msg := range logs {
        _ = stream.Send(&logpb.LogEntry{
            Timestamp: time.Now().Format(time.RFC3339),
            Message:   msg,
        })
        time.Sleep(500 * time.Millisecond)
    }

    res, _ := stream.CloseAndRecv()
    fmt.Printf("✅ Upload Result: %s, Total Logs: %d\n", res.GetResult(), res.GetCount())
}
