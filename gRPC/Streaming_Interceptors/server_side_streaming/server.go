type stockServer struct {
    stockpb.UnimplementedStockServiceServer
}

func (s *stockServer) GetStockUpdates(req *stockpb.StockRequest, stream stockpb.StockService_GetStockUpdatesServer) error {
    symbol := req.GetSymbol()
    for i := 0; i < 5; i++ {
        price := &stockpb.StockPrice{
            Symbol:    symbol,
            Price:     rand.Float32() * 100,
            Timestamp: time.Now().Format(time.RFC3339),
        }
        if err := stream.Send(price); err != nil {
            return err
        }
        time.Sleep(1 * time.Second)
    }
    return nil
}
