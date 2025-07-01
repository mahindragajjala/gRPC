type chatServer struct {
    chatpb.UnimplementedChatServiceServer
}

func (s *chatServer) ChatStream(stream chatpb.ChatService_ChatStreamServer) error {
    for {
        msg, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }

        log.Printf("💬 Received from %s: %s", msg.GetUser(), msg.GetMessage())

        response := &chatpb.ChatMessage{
            User:      "Server",
            Message:   "Received: " + msg.GetMessage(),
            Timestamp: time.Now().Format(time.RFC3339),
        }

        if err := stream.Send(response); err != nil {
            return err
        }
    }
}
