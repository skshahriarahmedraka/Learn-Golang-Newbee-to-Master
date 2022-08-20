

# gRPC programming use
**@author: Davie**
**Copyright: Beijing Qianfeng Internet Technology Co., Ltd.**

## One, gRPC call
In the last lesson, we learned how to use the gRPC framework to implement service call programming. In the gRPC framework, such as the way we learned in the last lesson to pass data between the client and the server through the definition of the message structure, we call it "single RPC", also known as the simple mode. In addition, there is also an implementation of RPC calls in data flow mode in gRPC, which is exactly what we will learn in this lesson.

### 1.1, server stream RPC
In the RPC implementation of the server stream mode, after the server receives the client's request, the processing ends and returns a data response stream. After sending all the response data requested by the client, the status details of the server and optional tracking metadata are sent to the client. The server-side streaming RPC implementation case is as follows:
#### 1.1.1, service interface definition
Define the service interface in the .proto file, and use the server-side stream mode to define the service interface, as shown below:
```proto
...
//Order service service definition
service OrderService {
    rpc GetOrderInfos (OrderRequest) returns (stream OrderInfo) (); //server-side streaming mode
}
```
We can see that the difference between the data in the previous simple mode as the parameter and return value of the service interface is that the return value of the service interface here is decorated with stream. When the interface is called by stream modification, the server will return the data to the client in the form of a data stream.

#### 1.1.2 Compile .proto file and generate pb.go file
Use the gRPC plugin compilation command to compile the .proto file, the compilation command is as follows:
```go
protoc --go_out=plugins=grpc:. message.proto
```

#### 1.1.3 Automatically generated file changes
When it is different from the implementation of sending and carrying data in the data structure, the data sending and receiving in the stream mode is completed by using a new function method. Among the automatically generated go code programs, each service interface corresponding to the flow mode will automatically generate corresponding individual client and server programs, and the corresponding structure implementation. The specific programming is shown in the figure below:
##### 1.1.3.1 Server-side automatic generation
```go
type OrderService_GetOrderInfosServer interface {
	Send(*OrderInfo) error
	grpc.ServerStream
}

type orderServiceGetOrderInfosServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrderInfosServer) Send(m *OrderInfo) error {
	return x.ServerStream.SendMsg(m)
}
```
In the streaming mode, the server of the service interface provides the Send method to send the data in the form of a stream
##### 1.1.3.2 Automatic client generation
```go
type OrderService_GetOrderInfosClient interface {
	Recv() (*OrderInfo, error)
	grpc.ClientStream
}

type orderServiceGetOrderInfosClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrderInfosClient) Recv() (*OrderInfo, error) {
	m := new(OrderInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```
In the streaming mode, the client of the service interface provides the Recv() method to receive the streaming data sent by the server.

#### 1.1.4 Service Coding Implementation
After defining the service interface and compiling the generated code file, the defined service can be coded and implemented according to the rules. The specific service coding implementation is as follows:
```go
//Order service realization
type OrderServiceImpl struct {
}

//Get order information s
func (os *OrderServiceImpl) GetOrderInfos(request *message.OrderRequest, stream message.OrderService_GetOrderInfosServer) error {
	fmt.Println("server stream RPC mode")

	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "Clothes", OrderStatus: "paid"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "Snacks", OrderStatus: "paid"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "Food", OrderStatus: "Unpaid"},
	}
	for id, info := range orderMap {
		if (time.Now().Unix() >= request.TimeStamp) {
			fmt.Println("Order serial number ID:", id)
			fmt.Println("Order details:", info)
			//Send to the client through streaming mode
			stream.Send(&info)
		}
	}
	return nil
}
```
The GetOrderInfos method is the specific implementation of the service interface. Because it is a stream mode development, the server sends the data in the form of a stream. Therefore, the second parameter type of the method is OrderService_GetOrderInfosServer, and the parameter type is an interface, which contains the Send method. To allow streaming data to be sent. The specific implementation of the Send method is in the compiled pb.go file, and the grpc.SeverStream.SendMsg method is further called.

#### 1.1.5 Service registration and monitoring processing
The monitoring and processing of the service is the same as the content learned in the previous article, and the steps are still the same:
```go
func main() {
	server := grpc.NewServer()
	//register
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
```

#### 1.1.6 Client data reception
The server uses the Send method to send the data in the form of a stream, and the client can use the Recv() method to receive the stream data. Because the data is constantly lost, the for infinite loop is used to read the data stream. When the io is read When .EOF, it means the end of the stream data. The client data reading is implemented as follows:
```go
...
for {
		orderInfo, err := orderInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("End of reading")
			return
		}
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("The information read:", orderInfo)
	}
...
```

#### 1.1.7 Running results
According to the sequence, run the server.go file and client.go file in turn to get the running result.
##### 1.1.7.1 Server running results
```go
 Server-side streaming RPC mode
Order serial number ID: 201907300001
Order details: {201907300001 Clothes paid{} [] 0}
Order serial number ID: 201907310001
Order details: {201907310001 Snacks paid{} [] 0}
Order serial number ID: 201907310002
Order details: {201907310002 Unpaid food{} [] 0}
```

##### 1.1.7.2 Client running results
```go
Client request RPC call: server stream mode
Information read: OrderId:"201907310001" OrderName:"\351\233\266\351\243\237" OrderStatus:"\345\267\262\344\273\230\346\254\276" 
Information read: OrderId:"201907310002" OrderName:"\351\243\237\345\223\201" OrderStatus:"\346\234\252\344\273\230\346\254\276" 
Information read: OrderId:"201907300001" OrderName:"\350\241\243\346\234\215" OrderStatus:"\345\267\262\344\273\230\346\254\276" 
End of reading
```

### 1.2, client streaming mode
The above demonstrates how the server returns data in the form of a data stream. Correspondingly, there is also a form in which the client sends request data in the form of a stream.

#### 1.2.1 Definition of Service Interface
Similar to the server side, the RPC service declaration format of the client stream mode is to use stream to modify the receiving parameters of the service interface, as shown below:
```proto
...
//Order service service definition
service OrderService {
    rpc AddOrderList (stream OrderRequest) returns (OrderInfo) (); //Client stream mode
}
```

#### 1.2.2 Compile .proto file
Use the compile command to compile the .protow file. The interface of the service interface is also automatically generated in the client streaming mode.
##### 1.2.2.1 Automatically generated service flow interface implementation
```go
type OrderService_AddOrderListServer interface {
	SendAndClose(*OrderInfo) error
	Recv() (*OrderRequest, error)
	grpc.ServerStream
}

type orderServiceAddOrderListServer struct {
	grpc.ServerStream
}

func (x *orderServiceAddOrderListServer) SendAndClose(m *OrderInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceAddOrderListServer) Recv() (*OrderRequest, error) {
	m := new(OrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```
The SendAndClose and Recv methods are methods owned by the server object in the client stream mode.

##### 1.2.2.2 Automatically generated client stream interface implementation
```go
type OrderService_AddOrderListClient interface {
	Send(*OrderRequest) error
	CloseAndRecv() (*OrderInfo, error)
	grpc.ClientStream
}

type orderServiceAddOrderListClient struct {
	grpc.ClientStream
}

func (x *orderServiceAddOrderListClient) Send(m *OrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceAddOrderListClient) CloseAndRecv() (*OrderInfo, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(OrderInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```
Send and CloseAndRecv are methods owned by the client object in the client stream mode.

#### 1.2.3 Realization of Service
The specific implementation of the service interface of the client stream mode is as follows:
```go
//Order service realization
type OrderServiceImpl struct {
}

//Add order information service implementation
func (os *OrderServiceImpl) AddOrderList(stream message.OrderService_AddOrderListServer) error {
	fmt.Println("Client Stream RPC Mode")

	for {
		//Read data information from the stream
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("End of reading data")
			result := message.OrderInfo{OrderStatus: "End of reading data"}
			return stream.SendAndClose(&result)
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		//Print the received data
		fmt.Println(orderRequest)
	}
}
```

#### 1.2.4 Service registration and monitoring processing
The same service registration and monitoring processing method is still used to register and monitor the service.
```go
func main() {

	server := grpc.NewServer()
	//register
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
```

#### 1.2.5 Client implementation
The client calls the send method to stream data to the server. The specific implementation is as follows:
```go
...
//Call the service method
	addOrderListClient, err := orderServiceClient.AddOrderList(context.Background())
	if err != nil {
		panic(err.Error())
	}
	//Call method to send stream data
	for _, info := range orderMap {
		err = addOrderListClient.Send(&info)
		if err != nil {
			panic(err.Error())
		}
	}

	for {
		orderInfo, err := addOrderListClient.CloseAndRecv()
		if err == io.EOF {
			fmt.Println("The end of reading data")
			return
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(orderInfo.GetOrderStatus())
	}
```

#### 1.2.6 Program operation
##### 1.2.6.1 Server
Run the case, the program output is as follows:
```go
 Client streaming RPC mode
201907300001 clothes paid
201907310001 Snacks paid
201907310002 Food unpaid
 End of reading data 
 Client streaming RPC mode
201907300001 clothes paid
201907310001 Snacks paid
201907310002 Food unpaid
 End of reading data 
```

##### 1.2.6.2 Client
The output of the client running program is as follows:
```go
Client request RPC call: client stream mode
 End of reading data 
 Reading data is over 
```

### 1.3, two-way flow mode
The server-side streaming mode and the client-side streaming mode have been discussed above. If the client and server streaming modes are combined, it is the third mode, the two-way streaming mode. That is, when the client sends data, it is sent as stream data, and the data returned by the server is also sent as a stream, so it is called a two-way stream mode.
#### 1.3.1 Definition of two-way streaming service
```go
//Order service service definition
service OrderService {
    rpc GetOrderInfos (stream OrderRequest) returns (stream OrderInfo) (); //Two-way stream mode
}
```

#### 1.3.2 Compile .proto file
##### 1.3.2.1 Server-side interface implementation
```go
type OrderService_GetOrderInfosServer interface {
	Send(*OrderInfo) error
	Recv() (*OrderRequest, error)
	grpc.ServerStream
}

type orderServiceGetOrderInfosServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrderInfosServer) Send(m *OrderInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceGetOrderInfosServer) Recv() (*OrderRequest, error) {
	m := new(OrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```

##### 1.3.2.2 Client interface implementation
```go
type OrderService_GetOrderInfosClient interface {
	Send(*OrderRequest) error
	Recv() (*OrderInfo, error)
	grpc.ClientStream
}

type orderServiceGetOrderInfosClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrderInfosClient) Send(m *OrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceGetOrderInfosClient) Recv() (*OrderInfo, error) {
	m := new(OrderInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```

#### 1.3.3 Service Implementation
```go
//Implement grpc bidirectional stream mode
func (os *OrderServiceImpl) GetOrderInfos(stream message.OrderService_GetOrderInfosServer) error {

	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Data Reading End")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}

		fmt.Println(orderRequest.GetOrderId())
		orderMap := map[string]message.OrderInfo{
			"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "Clothes", OrderStatus: "paid"},
			"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "Snacks", OrderStatus: "paid"},
			"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "Food", OrderStatus: "Unpaid"},
		}

		result := orderMap[orderRequest.GetOrderId()]
		//send data
		err = stream.Send(&result)
		if err == io.EOF {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
```

#### 1.3.4 Server and client programming implementation
##### 1.3.4.1 Server implementation
```go
func main() {
	server := grpc.NewServer()
	//register
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
```

##### 1.3.4.2 Client implementation
```go
func main() {

	//1, Dail connection
	conn, err := grpc.Dial("localhost:8092", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	fmt.Println("Client request RPC call: two-way stream mode")
	orderIDs := []string{"201907300001", "201907310001", "201907310002"}

	orderInfoClient, err := orderServiceClient.GetOrderInfos(context.Background())
	for _, orderID := range orderIDs {
		orderRequest := message.OrderRequest{OrderId: orderID}
		err := orderInfoClient.Send(&orderRequest)
		if err != nil {
			panic(err.Error())
		}
	}

	//closure
	orderInfoClient.CloseSend()

	for {
		orderInfo, err := orderInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("End of reading")
			return
		}
		if err != nil {
			return
		}
		fmt.Println("The information read:", orderInfo)
	}
}
```

## Two, TLS verification and Token authentication
In the last lesson, we learned and mastered the four flow modes of the grpc-go framework. In the actual production environment, a fully functional service includes not only basic method invocation and data interaction functions, but also authorization and authentication, data tracking, load balancing, etc. In this lesson, let's take a look at how to implement authorization authentication and how to intercept processing in addition to the gRPC call process.

### 2.1 Authorization
Two authorization methods are supported by default in gRPC, namely: SSL/TLS authentication method and Token-based authentication method.

#### 2.1.1 SSL/TLS authentication method
The full name of SSL is Secure Sockets Layer, also known as the Secure Sockets Layer. It is a standard security protocol used to establish an encrypted link between the client and the server during the communication process.
The full name of TLS is Transport Layer Security, and TLS is an upgraded version of SSL. In the process of use, it is often used to combine SSL and TLS to write SSL/TLS.
In short, SSL/TLS is a security protocol used for encryption in network communications.
##### 2.1.1.1 The working principle of SSL/TLS
The use of SSL/TLS protocol to securely encrypt the communication connection is achieved through asymmetric encryption. The so-called asymmetric encryption method is also called public key encryption, and the key pair is composed of two keys: a public key and a private key. The private key and the public key exist as a pair, the private key is generated first, and the corresponding public key is generated from the private key. The public key can be made public, and the private key should be properly stored.

In the encryption process: If the client wants to initiate a link to the server, it will first request the public key to be encrypted from the server. After obtaining the public key, the client uses the public key to encrypt the information, and the server receives the encrypted information, uses the private key to decrypt the information and performs other subsequent processing to complete the entire channel encryption and data transmission process.
##### 2.1.1.2 Making a certificate
You can install openssl on your local computer and generate the corresponding certificate.
```openssl
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

##### 2.1.1.3 Programming to implement server

```go
type MathManager struct {
}

func (mm *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {
	fmt.Println("Server Add Method")
	result := request.Args1 + request.Args2
	fmt.Println("The calculation result is:", result)
	response = new(message.Response)
	response.Code = 1;
	response.Message = "Successful execution"
	return response, nil
}

func main() {

	//TLS authentication
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
	if err != nil {
		grpclog.Fatal("Failed to load the certificate file", err)
	}

	//Instantiate grpc server, turn on TLS authentication
	server := grpc.NewServer(grpc.Creds(creds))

	message.RegisterMathServiceServer(server, new(MathManager))

	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
```

##### 2.1.1.3 Programming to realize the client
```
func main() {

	//TLS connection
	creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "go-grpc-example")
	if err != nil {
		panic(err.Error())
	}
	//1, Dail connection
	conn, err := grpc.Dial("localhost:8092", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	serviceClient := message.NewMathServiceClient(conn)

	addArgs := message.RequestArgs{Args1: 3, Args2: 5}

	response, err := serviceClient.AddMethod(context.Background(), &addArgs)
	if err != nil {
		grpclog.Fatal(err.Error())
	}

	fmt.Println(response.GetCode(), response.GetMessage())
}
```

#### 2.1.2 Token-based authentication method

##### 2.1.2.1 Introduction to Token Authentication
In the process of web application development, we often use another authentication method for identity verification, that is: Token authentication. Token-based authentication is stateless, and there is no need to store user information services in the server or session.

##### 2.1.2.2 Token authentication process
The main process of identity verification based on Token authentication is: before the client sends a request, it first initiates a request to the server, and the server returns a generated token to the client. The client saves the token and uses it for each subsequent request, carrying the token parameter. The server will first verify the token before processing the request. Only if the token verification is successful, will it process and return the relevant data.

##### 2.1.2.3 Custom Token Authentication of gRPC
In gRPC, developers are allowed to customize their own authentication rules and pass

```go
grpc.WithPerRPCCredentials()
```

Set up custom authentication rules. The WithPerRPCCredentials method receives a PerRPCCredentials type parameter. If you look further, you can know that PerRPCCredentials is an interface, defined as follows:

```go
type PerRPCCredentials interface {
    GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
    RequireTransportSecurity() bool
}
```

Therefore, developers can implement the above interface to define their own token information.

In this case, our custom token authentication structure is as follows:
```go
//token authentication
type TokenAuthentication struct {
	AppKey string
	AppSecret string
}

//Organize token information
func (ta *TokenAuthentication) RequestMetaData(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid": ta.AppKey,
		"appkey": ta.AppSecret,
	}, nil
}

//Whether to perform secure transmission based on TLS authentication
func (a *TokenAuthentication) RequireTransportSecurity() bool {
	return true
}
```

It should be noted that the appid and appkey fields in the RequestMetaData method need to be kept in lower case, not upper case. The RequireTransportSecurity method is used to set whether to perform secure transmission based on tls authentication.

When the client connects, we pass in the custom token authentication information as a parameter.
```go
auth := TokenAuthentication{
		AppKey: "hello",
		AppSecret: "20190812",
}
conn, err := grpc.Dial("localhost:8093", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
if err != nil {
	panic(err.Error())
}
```

##### 2.1.2.4 Server
To realize the judgment of the token request parameters in the calling method of the server, the token authentication information can be obtained through metadata. The specific implementation is as follows:
```go
func (mm *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {

	//Through metadata
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "No Token authentication information")
	}

	var appKey string
	var appSecret string

	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}

	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "hello" || appSecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Token is illegal")
	}
	fmt.Println("Server Add Method")
	result := request.Args1 + request.Args2
	fmt.Println("The calculation result is:", result)
	response = new(message.Response)
	response.Code = 1;
	response.Message = "Successful execution"
	return response, nil
}
```

Finally, run the project and the token authentication is successful. If the client modifies the token information and runs it again, it will prompt that the token is illegal.

## Third, the use of interceptors
### 3.1, demand
In the previous lesson, we learned to use two authentication methods in the gRPC framework: TLS verification and Token verification.

However, in the server-side method, each method needs to judge the token. The program efficiency is too low. You can optimize the processing logic. Before calling the specific method of the server, first intercept and perform token verification judgment. This method is called interceptor processing.

In addition to the token verification judgment processing here, log processing can also be performed.

### 3.2, Interceptor
To use an interceptor, you first need to register.
In grpc programming implementation, you can add interceptor settings in NewSever, in grpc framework you can set a custom interceptor through the UnaryInterceptor method, and return ServerOption. The specific code is as follows:
```go
grpc.UnaryInterceptor()
```
UnaryInterceptor() receives an UnaryServerInterceptor type, continue to check the source code definition, you can find that UnaryServerInterceptor is a func, defined as follows:
```go
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
```

With the above code, if the developer needs to register a custom interceptor, he needs to customize the definition of UnaryServerInterceptor.

### 3.3, custom UnaryServerInterceptor
Next, customize the implementation of func, which conforms to the standard of UnaryServerInterceptor, and implement the verification logic of the token in the definition of the func. The custom implementation of func is as follows:
```go
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	//Through metadata
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "No Token authentication information")
	}

	var appKey string
	var appSecret string
	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}
	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "hello" || appSecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Token is illegal")
	}
	//Pass the token verification and continue to process the request
	return handler(ctx, req)
}
```

In the custom TokenInterceptor method definition, consistent with the verification logic of the previous method call in the service, the token authentication information carried in the request header is taken from the metadata and verified whether it is correct. If the token verification is passed, the subsequent logic of the request will continue to be processed, and the subsequent continued processing can be processed by grpc.UnaryHandler. grpc.UnaryHandler is also a method, and its specific implementation is the service method implemented by the developer. The source code of grpc.UnaryHandler interface definition is as follows:
```go
type UnaryHandler func(ctx context.Context, req interface{}) (interface{}, error)
```

### 3.4, interceptor registration
Register the interceptor when the server calls grpc.NewServer. The details are as follows:
```go
server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(TokenInterceptor))
```

### 3.5、Project operation
Run the server.go program and client.go program in turn to get the correct result of the program. Modifying the value of the token can verify the effect of the interceptor in the illegal situation of the token.

