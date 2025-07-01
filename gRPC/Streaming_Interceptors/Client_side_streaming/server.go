type logServer struct {
    logpb.UnimplementedLogServiceServer
}

func (s *logServer) UploadLogs(stream logpb.LogService_UploadLogsServer) error {
    var count int32
    for {
        entry, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(&logpb.UploadStatus{
                Result: "Logs received",
                Count:  count,
            })
        }
        if err != nil {
            return err
        }

        log.Printf("📩 Log: [%s] %s", entry.GetTimestamp(), entry.GetMessage())
        count++
    }
}
