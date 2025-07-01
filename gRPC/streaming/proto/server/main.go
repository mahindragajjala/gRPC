func (s *userServer) StreamAllUsers(req *usermgmtpb.StreamAllUsersRequest, stream usermgmtpb.UserService_StreamAllUsersServer) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    for _, user := range s.users {
        if err := stream.Send(user); err != nil {
            return err
        }
        time.Sleep(time.Millisecond * 500) // simulate delay
    }

    return nil
}

//This sends each user one at a time to the client via stream.Send().
