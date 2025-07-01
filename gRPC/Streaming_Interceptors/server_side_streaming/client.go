func main() {
    conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
    client := stockpb.NewStockServiceClient(conn)

    stream, _ := client.GetStockUpdates(context.Background(), &stockpb.StockRequest{Symbol: "NIFTY50"})

    for {
        msg, err := stream.Recv()
        if err == io.EOF {
            break
        }
        fmt.Printf("📈 Price: %v at %v\n", msg.Price, msg.Timestamp)
    }
}
