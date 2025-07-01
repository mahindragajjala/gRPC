In gRPC (using Protocol Buffers), a message is used to define 
structured data types that are sent between the client and server.

message MessageName {
  type fieldName = tag;
}


message User {
  string name = 1;
  int32 age = 2;
  string email = 3;
}


This defines a structured message User with three fields.

 Rules:
 Feature              Description                                                          
 Types            `int32`, `int64`, `string`, `bool`, `float`, `double`, `bytes`, etc. 
 Tag              Unique number per field (1–536870911), required for serialization    
 Default Values   All fields are optional by default in `proto3` (unset = zero value)  
 Nested Messages  Messages can be nested inside each other                             
 Repeated Fields  Represent arrays/lists                                               


message Address {
  string city = 1;
  string state = 2;
}

message User {
  string id = 1;
  string name = 2;
  int32 age = 3;
  Address address = 4;              // Nested message
  repeated string hobbies = 5;      // List of strings
}
