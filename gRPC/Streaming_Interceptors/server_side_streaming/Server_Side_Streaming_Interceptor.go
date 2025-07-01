//This interceptor logs incoming stream requests and responses.
func streamingLoggerInterceptor(
    srv interface{},
    ss grpc.ServerStream,
    info *grpc.StreamServerInfo,
    handler grpc.StreamHandler,
) error {
    log.Printf("📥 Stream started: %s, IsClientStream: %v, IsServerStream: %v", info.FullMethod, info.IsClientStream, info.IsServerStream)
    
    wrapped := &loggingServerStream{ServerStream: ss}
    err := handler(srv, wrapped)

    log.Printf("📤 Stream ended: %s, Error: %v", info.FullMethod, err)
    return err
}

type loggingServerStream struct {
    grpc.ServerStream
}

func (s *loggingServerStream) SendMsg(m interface{}) error {
    log.Printf("Sending message to client: %v", m)
    return s.ServerStream.SendMsg(m)
}

Register the interceptor in the gRPC server:
server := grpc.NewServer(
    grpc.StreamInterceptor(streamingLoggerInterceptor),
)
