

# RESTful API design standards and practices
**@author: Davie**
**Copyright: Beijing Qianfeng Internet Technology Co., Ltd.**

## background
In the previous few courses, some core functions and core mechanisms of go-micro have been introduced. The implementation is focused on the microservices. In the actual development process, microservices are only required to be deployed as background programs, and interaction and data need to be provided to web front-end user-end products as a whole. Therefore, let's take a look at how microservices interact with the web.

## Divide the calling range
In the overall system architecture, we will divide the system into the foreground and the background. The front desk is responsible for interacting with users, displaying data, and performing operations. The background is responsible for business logic processing, data persistence and other operations. In the process of system operation, function calls may occur in the foreground and the background, and the background and the background:
* Internal calls: The mutual calls between the various microservices in the background belong to the internal calls of the system background, which are called internal calls.

* External call: The interface request call between the foreground and the background is usually called from outside the city.

## Technology Selection
In development practice, we have different technical solutions for external calls and internal calls.

* RPC call: The internal calls between various services in the background. In order to achieve efficient service interaction, RPC is usually used for implementation.

* REST: For the scenario where the front-end client interacts with the back-end through the HTTP interface. Because it involves the management and operation of different resources, RESTful standards are often used for implementation.

## Go-Micro API Gateway
There is an API gateway function in the Micro framework. The role of the API gateway is to act as a proxy for the microservices, responsible for proxying the RPC method of the microservices into web requests supporting the HTTP protocol, and at the same time exposing the URL used by the client.

### Install Micro Tools
To use the api gateway function of go-micro. Need to download the Micro source code and install Mico.

#### Install Micro
It can be downloaded and installed directly through the go get command. The specific command is:
```go
go get -u github.com/micro/micro
```
When using the go get command to download the source code and the corresponding dependencies, the domestic network environment may cause the download failure error.

If a connection timeout error occurs during the installation process, you can manually download the corresponding dependent library, and then install it manually. Therefore, the solution is divided into two steps:
##### Install golang's net, crypt, text and other libraries
```
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc  
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net  
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text  
git clone https://github.com/golang/crypto.git $GOPATH/src/golang.org/x/crypto
```
Use the git clone command to download the required code base.

##### Install micro
```
go install github.com/micro/micro
```
 Use the go install command to install the micro framework and wait for the command execution to end.
 
#### Verify that Micro is installed successfully
After the micro series tool is installed successfully, you can check it through the command.
```go
micro --version
micro version 1.9.1
```
As above, the output of micro version 1.9.1 means that the micro installation is successful.


## How Micro API Works
The micro tool provides the function of building api gateway services, and is implemented by programming based on the go-micro framework. The core function is to proxy the service in the form of RPC into a WEB API request that supports the HTTP protocol.

## Run Micro api service
The micro api can be started by the following command:
```go
micro api
``` 
![Run micro service](./img/WX20190912-074346@2x.png)

### Reverse proxy API service start
In the Micro api function, it supports multiple methods of processing request routing, which we call Handler. Including: API Handler, RPC Handler, reverse proxy, Event Handler, RPC and other five methods. In this case, we use a reverse proxy to demonstrate.

#### Reverse proxy
* Format: /[service]
* Request/Response: HTTP method
* Set via --handler=proxy when micro api is started

Therefore, the micro api gateway service startup command in the form of a reverse proxy is:
```go
micro api --handler=http
```

In this case, we combine the reverse proxy of the micro api and the HTTP WEB request represented by REST.

## Install go-restful
You can implement RESTful-style path mapping by installing the go-restful library, thereby implementing HTTP WEB API services. The command to install go-restful is as follows:
```go
go get github.com/emicklei/go-restful
```
After downloading, you can view the corresponding source code in $GOPATH/src/github.com/emicklei/go-restful in the current system directory.

> We use a service that obtains a certain student's information as an example to explain the programming implementation of micro api.

## Service definition and compilation
Define the proto file of the student message body:
```proto
syntax ='proto3';

package proto;

message Student {
    string id = 1;
    string name = 2;
    int32 grade = 3;
    string classes = 4;
}

message Request {
    string name = 1;
}

service StudentService {
    rpc GetStudent (Request) returns (Student);
}
```
The Student, Request message body and rpc service are defined in the proto file. To use the micro api gateway function to compile the proto file, a micro file needs to be generated. To compile and generate this file, a new protoc-gen-micro library needs to be used. The command to install the protoc-gen-micro library is as follows:
```go
go get github.com/micro/protoc-gen-micro
```
To compile the proto file again, you need to specify two parameters: go_out and micro_out. The detailed commands are as follows:
```go
protoc --go_out=. --micro_out=. student.proto
```
After the above command is executed successfully, two go language files will be automatically generated: student.pb.go and student.micro.go.

The content generated in the micro.go file includes the instantiation of the service and the underlying implementation of the corresponding service method.

## Server implementation
We all know that normal web services process http requests through routing. The same here, we can parse the HTTP request interface through routing processing, and the service object contains routing processing methods. The detailed code is as follows:
```go
...
type StudentServiceImpl struct {
}

//Service implementation
func (ss *StudentServiceImpl) GetStudent(ctx context.Context, request *proto.Request, resp *proto.Student) error {

	//tom
	studentMap := map[string]proto.Student{
		"davie": proto.Student{Name: "davie", Classes: "Software Engineering Major", Grade: 80},
		"steven": proto.Student{Name: "steven", Classes: "Computer Science and Technology", Grade: 90},
		"tony": proto.Student{Name: "tony", Classes: "Computer Network Engineering", Grade: 85},
		"jack": proto.Student{Name: "jack", Classes: "Business Administration", Grade: 96},
	}

	if request.Name == "" {
		return errors.New("Request parameter error, please request again.")
	}

	//Get the corresponding student
	student := studentMap[request.Name]
	if student.Name != "" {
		fmt.Println(student.Name, student.Classes, student.Grade)
		*resp = student
		return nil
	}
	return errors.New("Relevant student information was not queried")
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.student"),
	)

	service.Init()
    proto.RegisterStudentServiceHandler(service.Server(), new(StudentServiceImpl))

	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
...
```
The server program implements the service and runs the service.

## REST mapping
Now, the RPC service has been written. We need to program the proxy function of the API to process HTTP requests.
In the rest.go file, implement rest mapping, the detailed code is as follows:
```go
type Student struct {
}

var (
	cli proto.StudentService
)

func (s *Student) GetStudent(req *restful.Request, rsp *restful.Response) {

	name := req.PathParameter("name")
	fmt.Println(name)
	response, err := cli.GetStudent(context.TODO(), &proto.Request{
		Name: name,
	})

	if err != nil {
		fmt.Println(err.Error())
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}

func main() {

	service := web.NewService(
		web.Name("go.micro.api.student"),
	)

	service.Init()

	cli = proto.NewStudentService("go.micro.srv.student", client.DefaultClient)

	student := new(Student)
	ws := new(restful.WebService)
	ws.Path("/student")
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{name}").To(student.GetStudent))

	wc := restful.NewContainer()
	wc.Add(ws)

	service.Handle("/", wc)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
```

