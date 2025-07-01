type wrappedStream struct {
    grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
    err := w.ServerStream.RecvMsg(m)
    log.Printf("⬅️  Client Sent: %v", m)
    return err
}

func (w *wrappedStream) SendMsg(m interface{}) error {
    log.Printf("➡️  Server Sent: %v", m)
    return w.ServerStream.SendMsg(m)
}

func bidiStreamInterceptor(
    srv interface{},
    ss grpc.ServerStream,
    info *grpc.StreamServerInfo,
    handler grpc.StreamHandler,
) error {
    log.Printf("🔄 Bidirectional stream started: %s", info.FullMethod)
    err := handler(srv, &wrappedStream{ss})
    log.Printf("🛑 Bidirectional stream ended: %s", info.FullMethod)
    return err
}


Register with:
grpc.NewServer(grpc.StreamInterceptor(bidiStreamInterceptor))
