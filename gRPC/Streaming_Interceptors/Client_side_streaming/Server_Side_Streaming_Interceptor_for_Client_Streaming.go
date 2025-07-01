//This interceptor wraps the RecvMsg method to
//log each message received from the client.

func loggingClientStreamInterceptor(
    srv interface{},
    ss grpc.ServerStream,
    info *grpc.StreamServerInfo,
    handler grpc.StreamHandler,
) error {
    log.Printf("📥 Client stream started: %s", info.FullMethod)
    wrapped := &loggingStream{ServerStream: ss}
    err := handler(srv, wrapped)
    log.Printf("📤 Client stream ended: %s, Error: %v", info.FullMethod, err)
    return err
}

type loggingStream struct {
    grpc.ServerStream
}

func (s *loggingStream) RecvMsg(m interface{}) error {
    err := s.ServerStream.RecvMsg(m)
    log.Printf("🔄 Received from client: %v", m)
    return err
}

//Register it like this:
                    server := grpc.NewServer(
                        grpc.StreamInterceptor(loggingClientStreamInterceptor),
                    )
