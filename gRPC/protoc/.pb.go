This file contains:
          Go structs for each message in your .proto file
          Methods for marshaling/unmarshaling (serialize/deserialize)
          Helper functions for working with these messages


Contains:
For this .proto:
                  message UserRequest {
                    string user_id = 1;
                  }
                  
                  message UserResponse {
                    string name = 1;
                    int32 age = 2;
                    string email = 3;
                  }
The user.pb.go will generate:
                  type UserRequest struct {
                      UserId string
                  }
                  
                  type UserResponse struct {
                      Name  string
                      Age   int32
                      Email string
                  }

It also includes:
                  Reset()
                  String()
                  ProtoReflect()
                  Marshal/Unmarshal logic


Use case 
Used when building data to send/receive:
                              req := &userpb.UserRequest{
                                  UserId: "U001",
                              }
