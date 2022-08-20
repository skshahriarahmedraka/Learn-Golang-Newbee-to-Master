

# Error handling

> @author：Han Ru
>
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

In actual engineering projects, we hope to quickly locate the problem through the error message of the program, but we don't like the redundant and verbose error handling code. The `Go` language does not provide `try...catch` exception handling methods like `Java` and `C#` languages, but throws upwards layer by layer through the return value of the function. This design encourages engineers to explicitly check for errors in the code instead of ignoring them. The advantage is to avoid missing errors that should be handled. But it brings a drawback, making the code verbose.

## 1.1 What is an error

What is wrong?

An error refers to a problem where there may be a problem. For example, when opening a file fails, this situation is expected.

An exception refers to a problem where there shouldn't be a problem. For example, a null pointer is referenced, which is unexpected. It can be seen that errors are part of the business process, while exceptions are not.



Errors in Go are also a type. Errors are represented by the built-in `error` type. Just like other types, such as int, float64,. Error values ​​can be stored in variables, returned from functions, and so on.

## 1.2 Demonstration error

Let's start with a sample program that tries to open a file that does not exist.

Sample code:

```go
package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Open("/test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
  //Read or write files according to f
    fmt.Println(f.Name(), "opened successfully")
}
```

> There are functions to open files in the os package:
>
> ​ func Open(name string) (file \*File, err error)
>
> If the file has been successfully opened, the Open function will return to file processing. If an error occurs when opening the file, a non-nil error will be returned.

​	

If a function or method returns an error, then by convention, it must be the last value returned by the function. Therefore, the value returned by the `Open` function is the last value.

The idiomatic way of handling errors is to compare the returned error with nil. A nil value indicates that no error occurred, while a non-nil value indicates that an error occurred. In our example, we check if the error is nil. If it is not nil, we just print the error and return from the main function.

operation result:

```
open /test.txt: No such file or directory
```

We get an error stating that the file does not exist.

## 1.3 Error type indication

Go language provides a very simple error handling mechanism through the built-in error interface.

Let's go a little deeper and see how to define the construction of error types. The error is an interface type with the following definition,

```go
type error interface {
    Error() string
}
```

It contains a method with an Error() string. Any type that implements this interface can be used as an error. This method provides a description of the error.

When printing an error, the fmt.Println function internally calls the Error() method to obtain the description of the error. This is how the error description is printed on one line.

**Different ways to extract more information from the error**

Now that we know that errors are an interface type, let's see how to extract more information about errors.

In the above example, we just printed the description of the error. If what we want is the actual path of the file that caused the error. One possible method is to parse the error string. This is the output of our program,

```
open /test.txt: No such file or directory  
```

We can parse this error message and get the file path "/test.txt" from it. But this is a bad method. In the new version of the language, the error description can be changed at any time, and our code will be interrupted.

Is there a way to reliably obtain the file name? The answer is yes, it can be done, the standard Go library uses different ways to provide more information about the error. Let's take a look.

 1. Assert the underlying structure type and get more information from the structure field

If you read the documentation of the open function carefully, you can see that it returns an error of type PathError. PathError is a struct type, and its implementation in the standard library is as follows,

```go
type PathError struct {  
    Op string
    Path string
    Err error
}

func (e *PathError) Error() string {return e.Op + "" + e.Path + ": "+ e.Err.Error()}  
```

From the above code, you can understand that PathError implements the error interface by declaring the `Error()string` method. This method connects the operation, path, and actual error and returns it. So we got the error message,

```
open /test.txt: No such file or directory 
```

The path field of the PathError structure contains the path of the file that caused the error. Let us modify the program written above and print out the path.

Modify the code:

```go
package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Open("/test.txt")
    if err, ok := err.(*os.PathError); ok {
        fmt.Println("File at path", err.Path, "failed to open")
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}
```

In the above program, we use type assertions to get the basic value of the error interface. Then we use the error to print the path. This program outputs,

```
File at path /test.txt failed to open  
```

2. Assert the underlying structure type and use methods to get more information

The second way to get more information is to assert the underlying type and get more information by calling methods of the struct type.

Sample code:

```go
type DNSError struct {  
    ...
}

func (e *DNSError) Error() string {  
    ...
}
func (e *DNSError) Timeout() bool {  
    ... 
}
func (e *DNSError) Temporary() bool {  
    ... 
}
```

As you can see from the code above, DNSError struct has two methods Timeout() bool and Temporary() bool, which return a boolean value indicating whether the error is due to timeout or temporary.

Let us write a program that asserts *DNSError and call these methods to determine whether the error is temporary or timed out.

```go
package main

import (  
    "fmt"
    "net"
)

func main() {  
    addr, err := net.LookupHost("golangbot123.com")
    if err, ok := err.(*net.DNSError); ok {
        if err.Timeout() {
            fmt.Println("operation timed out")
        } else if err.Temporary() {
            fmt.Println("temporary error")
        } else {
            fmt.Println("generic error: ", err)
        }
        return
    }
    fmt.Println(addr)
}
```

In the above program, we are trying to obtain the ip address of an invalid domain name, which is an invalid domain name. golangbot123.com. We get the potential value of the error by declaring it to enter *net.DNSError.

In our example, the error is neither temporary nor due to a timeout, so the program will print it out,

```
generic error: lookup golangbot123.com: no such host  
```

If the error is temporary or timed out, then the corresponding If statement will be executed and we can handle it appropriately.

3. Direct comparison

The third way to get more detailed information about the error is to compare directly with the variable of the wrong type. Let us understand this problem through an example.

The Glob function of the filepath package is used to return the names of all files matching the pattern. When there is an error in the pattern, the function will return an error ErrBadPattern.

ErrBadPattern is defined in the filepath package, as described below:

```go
var ErrBadPattern = errors.New("syntax error in pattern")  
```

errors.New() is used to create new errors.

When there is an error in the pattern, ErrBadPattern is returned by the Glob function.

Let's write a small program to check this error:

```go
package main

import (  
    "fmt"
    "path/filepath"
)

func main() {  
    files, error := filepath.Glob("[")
    if error != nil && error == filepath.ErrBadPattern {
        fmt.Println(error)
        return
    }
    fmt.Println("matched files", files)
}
```

operation result:

```
syntax error in pattern  
```

**Don't ignore errors**

Never ignore an error. Ignoring mistakes can cause trouble. Let me rewrite an example that lists the names of all files that match the pattern, ignoring the error handling code.

```go
package main

import (  
    "fmt"
    "path/filepath"
)

func main() {  
    files, _ := filepath.Glob("[")
    fmt.Println("matched files", files)
}
```

We already know from the previous examples that the mode is invalid. I ignored the error returned by the Glob function by using the blank identifier in the line number.

```
matched files []  
```

Since we ignored this error, the output looks like no file matching this pattern, but in fact the pattern itself is malformed. So don't ignore the error.



## 1.4 Custom errors

To create a custom error, you can use the New() function under the errors package, and the Errorf() function under the fmt package.

```go
//errors package:
func New(text string) error {}

//fmt package:
func Errorf(format string, a ...interface{}) error {}
```

Before using the New() function to create a custom error, let us understand how it is implemented. The implementation of the new functions in the error package is provided below.

```go
// Package errors implements functions to manipulate errors.
  package errors

  // New returns an error that formats as the given text.
  func New(text string) error {
      return &errorString{text}
  }

  // errorString is a trivial implementation of error.
  type errorString struct {
      s string
  }

  func (e *errorString) Error() string {
      return es
  }
```

Now that we know how the New() function works, let us use it in our own program to create a custom error.

We will create a simple program to calculate the area of ​​a circle. If the radius is negative, an error will be returned.

```go
package main

import (  
    "errors"
    "fmt"
    "math"
)

func circleArea(radius float64) (float64, error) {  
    if radius <0 {
        return 0, errors.New("Area calculation failed, radius is less than zero")
    }
    return math.Pi * radius * radius, nil
}

func main() {  
    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of ​​circle %0.2f", area)
}
```

operation result:

```
Area calculation failed, radius is less than zero 
```

Use Errorf to add more information to the error

The above program works well, but if we print out the actual radius that caused the error, it's not good. This is where the Errorf function of the fmt package comes in. This function formats the error according to a format specifier and returns a string as the value to satisfy the error.

Use the Errorf function to modify the program.

```go
package main

import (  
    "fmt"
    "math"
)

func circleArea(radius float64) (float64, error) {  
    if radius <0 {
        return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
    }
    return math.Pi * radius * radius, nil
}

func main() {  
    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of ​​circle %0.2f", area)
}
```

operation result:

```
Area calculation failed, radius -20.00 is less than zero  
```

Use struct types and fields to provide more information about the error

You can also use the struct type that implements the error interface as an error. This gives us more flexibility in error handling. In our example, if we want to access the radius that caused the error, then the only way now is to parse the error description area calculation failure, radius -20.00 is less than zero. This is not a correct approach, because if the description changes, our code will break.

We will use the standard library strategy explained in the previous tutorial, under "Assert the underlying structure type and get more information from the struct field", and use the struct field to provide access to the radius that caused the error. We will create a struct type that implements the error interface and use its fields to provide more information about the error.

The first step is to create a struct type to represent the error. The naming convention for error types is that the name should end with the text Error. Let's name the struct type areaError

```go
type areaError struct {  
    err string
    radius float64
}
```

The above struct type has a field radius, which stores the value of the radius responsible for the error, and the error field stores the actual error message.

The next step is to implement the error interface

```go
func (e *areaError) Error() string {  
    return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}
```

In the above code snippet, we use a pointer receiver area error to implement the Error() string method of the error interface. This method prints out the radius and error description.

```go
package main

import (  
    "fmt"
    "math"
)

type areaError struct {  
    err string
    radius float64
}

func (e *areaError) Error() string {  
    return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {  
    if radius <0 {
        return 0, &areaError{"radius is negative", radius}
    }
    return math.Pi * radius * radius, nil
}

func main() {  
    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
        if err, ok := err.(*areaError); ok {
            fmt.Printf("Radius %0.2f is less than zero", err.radius)
            return
        }
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of ​​circle %0.2f", area)
}
```

Program output:

```
Radius -20.00 is less than zero
```

Use the structure type method to provide more information about the error

In this section, we will write a program to calculate the area of ​​a rectangle. If the length or width is less than 0, this program will output an error.

The first step is to create a structure to represent the error.

```go
type areaError struct {  
    err string //error description
    length float64 //length which caused the error
    width float64 //width which caused the error
}
```

The above error structure type contains an error description field, and the length and width that caused the error.

Now that we have the error type, let us implement the error interface and add some methods to the error type to provide more information about the error.

```go
func (e *areaError) Error() string {  
    return e.err
}

func (e *areaError) lengthNegative() bool {  
    return e.length <0
}

func (e *areaError) widthNegative() bool {  
    return e.width <0
}
```

In the above code snippet, we return the error description of the `Error() string` method. When the length is less than 0, the lengthNegative() bool method returns true; when the width is less than 0, the widthNegative() bool method returns true. These two methods provide more information about the error. In this case, they say whether the area calculation failed because the length is negative or the width is negative. Therefore, we use the struct error type method to provide more information about the error.

The next step is to write the area calculation function.

```go
func rectArea(length, width float64) (float64, error) {  
    err := ""
    if length <0 {
        err += "length is less than zero"
    }
    if width <0 {
        if err == "" {
            err = "width is less than zero"
        } else {
            err += ", width is less than zero"
        }
    }
    if err != "" {
        return 0, &areaError{err, length, width}
    }
    return length * width, nil
}
```

The above rectArea function checks whether the length or width is less than 0, and if it returns an error message, it returns the area of ​​the rectangle as nil.

Main function:

```go
func main() {  
    length, width := -5.0, -9.0
    area, err := rectArea(length, width)
    if err != nil {
        if err, ok := err.(*areaError); ok {
            if err.lengthNegative() {
                fmt.Printf("error: length %0.2f is less than zero\n", err.length)

            }
            if err.widthNegative() {
                fmt.Printf("error: width %0.2f is less than zero\n", err.width)

            }
        }
        fmt.Println(err)
        return
    }
    fmt.Println("area of ​​rect", area)
}
```

operation result:

```
error: length -5.00 is less than zero  
error: width -9.00 is less than zero 
```



## 1.5 panic() and recover()

Two built-in functions panic and recover are introduced in Golang to trigger and terminate the exception handling process, and the keyword defer is introduced to delay the execution of the function behind defer.
The deferred function (function after defer) will not be executed until the function containing the defer statement is executed, regardless of whether the function containing the defer statement ends normally through return or ends abnormally due to panic. You can execute multiple defer statements in a function, and their execution order is opposite to the declaration order.
When the program is running, if it encounters a situation where a null pointer is referenced, a subscript is out of bounds, or the panic function is explicitly called, the execution of the panic function is triggered first, and then the delay function is called. The caller continues to pass the panic, so the process has been repeated in the call stack: the function stops executing, the delayed execution function is called, and so on. If there is no call to the recover function in the delayed function all the way, it will reach the starting point of the coroutine, the coroutine ends, and then all other coroutines, including the main coroutine (similar to the main thread in the C language, the coroutine ID is 1).

panic:
 1. Built-in functions
 2. If a panic statement is written in function F, the code to be executed afterwards will be terminated. If there is a list of defer functions to be executed in function F where panic is located, it will be executed in the reverse order of defer
 3. Return the caller G of function F. In G, the code after the statement of calling function F will not be executed. If there is a list of defer functions to be executed in function G, they are executed in the reverse order of defer. Here, defer is somewhat similar to try- finally in catch-finally
 4. Until the goroutine exits completely and reports an error

recover:
 1. Built-in functions
 2. Used to control the panicking behavior of a goroutine, capture panic, thereby affecting the behavior of the application
 3. General call suggestions
 a). In the defer function, use recever to terminate the panicking process of a gojroutine, thereby restoring normal code execution
 b). You can get the error passed through panic

To put it simply: Go can throw a panic exception, and then catch the exception through recover in defer, and then process it normally.



Errors and exceptions are the difference between error and panic in terms of Golang mechanism. The same is true for many other languages, such as C++/Java, which has no error but errno, and no panic but throw.

Golang errors and exceptions can be converted to each other:

1. Errors turn to exceptions. For example, the program logically tries to request a certain URL, and it tries at most three times. If the request fails during the three tries, it is an error. If the third attempt is unsuccessful, the failure is promoted to an exception.
2. Exception to error. For example, after the exception triggered by panic is recovered by recover, the variable of type error in the return value is assigned so that the upper-level function can continue the error handling process.

 

**Under what circumstances are expressed in error, and under what circumstances are expressed in exceptions, there must be a set of rules, otherwise it is easy for everything to be wrong or everything to be abnormal. **

The scope (scenario) of exception handling is given below:

1. Null pointer reference
2. Subscript out of bounds
3. Divisor is 0
4. Branches that should not appear, such as default
5. Input should not cause a function error


In other scenarios, we use error handling, which makes our functional interface very refined. For exceptions, we can choose to recover from a suitable upstream and print stack information so that the deployed program will not terminate.

 

**Note: Golang error handling has always been criticized by many people. Some people complain that half of the code is "if err != nil {/ print && error handling / }", which seriously affects the normal processing logic. When we distinguish between errors and exceptions and design functions according to rules, readability and maintainability will be greatly improved. **

 

## 1.6 Correct posture for error handling

**Position 1: When there is only one reason for failure, don’t use error**

Let's look at a case:

```go
func (self *AgentContext) CheckHostType(host_type string) error {
    switch host_type {
    case "virtual_machine":
        return nil
    case "bare_metal":
        return nil
    }
    return errors.New("CheckHostType ERROR:" + host_type)
}
```

 

We can see that there is only one reason for the failure of the function, so the return value type should be bool instead of error. Refactor the code: 

```go
func (self *AgentContext) IsValidHostType(hostType string) bool {
    return hostType == "virtual_machine" || hostType == "bare_metal"
}
```

 

Note: In most cases, there is more than one reason for the failure, especially for I/O operations. Users need to know more error information. At this time, the return value type is no longer a simple bool, but an error.

 

**Position 2: When there is no failure, no error is used**

Error is so popular in Golang that many people use error when designing functions, even if there is no reason for failure.
Let's take a look at the sample code:



```go
func (self *CniParam) setTenantId() error {
    self.TenantId = self.PodNs
    return nil
}
```

 

For the above function design, there will be the following calling code:

```go
err := self.setTenantId()
if err != nil {
    // log
    // free resource
    return errors.New(...)
}
```

 

According to our correct posture, refactor the code:

```go
func (self *CniParam) setTenantId() {
    self.TenantId = self.PodNs
}
```

 

So the calling code becomes:

```go
self.setTenantId()
```

 

**Posture 3: error should be placed at the end of the return value type list**

For the return value type error, it is used to pass error information, and it is usually placed last in Golang.

```go
resp, err := http.Get(url)
if err != nil {
    return nill, err
}
```

 

The same goes for bool as the return value type.

```go
value, ok := cache.Lookup(key) 
if !ok {
    // ...cache[key] does not exist... 
}
```

 

**Position 4: The error value is defined uniformly, instead of following the feeling**

When many people write code, they return errors.New(value) everywhere, and the error value may have different forms when expressing the same meaning. For example, the error value of "record does not exist" may be:

1. "record is not existed."
2. "record is not exist!"
3. "###record is not existed!!!"
4. ...

This causes the same error value to be scattered in a large piece of code. When the upper function needs to process a specific error value in a unified manner, it needs to roam all the lower code to ensure that the error value is unified. Unfortunately, sometimes there is a fish in the net, and this This method seriously hinders the reconstruction of the wrong value.

Therefore, we can refer to the C/C++ error code definition file and add an error object definition file to each package of Golang, as shown below:

```go
var ERR_EOF = errors.New("EOF")
var ERR_CLOSED_PIPE = errors.New("io: read/write on closed pipe")
var ERR_NO_PROGRESS = errors.New("multiple Read calls return no data or error")
var ERR_SHORT_BUFFER = errors.New("short buffer")
var ERR_SHORT_WRITE = errors.New("short write")
var ERR_UNEXPECTED_EOF = errors.New("unexpected EOF")
```

  

**Position 5: When the error is passed layer by layer, log is added layer by layer**

Adding logs at all levels is very convenient for fault location.

Note: As for finding faults through testing, not logs, it is still difficult for many teams to do so. If you or your team can do it, then please ignore this gesture.



**Position 6: Use defer for error handling**

We generally handle errors by judging the value of error. If the current operation fails, we need to destroy the resources that have been created in this function. The sample code is as follows:

```go
func deferDemo() error {
    err := createResource1()
    if err != nil {
        return ERR_CREATE_RESOURCE1_FAILED
    }
    err = createResource2()
    if err != nil {
        destroyResource1()
        return ERR_CREATE_RESOURCE2_FAILED
    }

    err = createResource3()
    if err != nil {
        destroyResource1()
        destroyResource2()
        return ERR_CREATE_RESOURCE3_FAILED
    }

    err = createResource4()
    if err != nil {
        destroyResource1()
        destroyResource2()
        destroyResource3()
        return ERR_CREATE_RESOURCE4_FAILED
    } 
    return nil
}
```

When Golang code is executed, if it encounters a defer closure call, it is pushed onto the stack. When the function returns, the closure will be called in the order of last in, first out.
**The parameter of the closure is passed by value, and the external variable is passed by reference, so the value of the external variable err in the closure becomes the latest err value when the external function returns. **
According to this conclusion, we refactor the above sample code:

```go
func deferDemo() error {
    err := createResource1()
    if err != nil {
        return ERR_CREATE_RESOURCE1_FAILED
    }
    defer func() {
        if err != nil {
            destroyResource1()
        }
    }()
    err = createResource2()
    if err != nil {
        return ERR_CREATE_RESOURCE2_FAILED
    }
    defer func() {
        if err != nil {
            destroyResource2()
                   }
    }()

    err = createResource3()
    if err != nil {
        return ERR_CREATE_RESOURCE3_FAILED
    }
    defer func() {
        if err != nil {
            destroyResource3()
        }
    }()

    err = createResource4()
    if err != nil {
        return ERR_CREATE_RESOURCE4_FAILED
    }
    return nil
}
```

**Position Seven: Don’t return an error immediately when you try a few times to avoid failure.**

If the error occurs accidentally, or caused by unpredictable problems. A wise choice is to retry the failed operation, sometimes succeeding on the second or third attempt. When retrying, we need to limit the retry interval or the number of retries to prevent unlimited retry.

Two cases:

1. When we usually go online, we try to request a certain URL, sometimes there is no response the first time, when we refresh again, there is a pleasant surprise.
2. A QA of the team once suggested that when Neutron's attach operation fails, it is best to try three times. This proved to be effective in the environment at that time.

 

**Position 8: When the upper function does not care about errors, it is recommended not to return error**

For some functions related to resource cleanup (destroy/delete/clear), if the sub-function has an error, just print the log without further feeding the error to the upper-level function, because in general, the upper-level function does not care about the execution result, or Even if you care, there is nothing you can do, so we recommend designing related functions to not return error.

 

**Position Nine: When an error occurs, do not ignore useful return values**

Generally, when a function returns a non-nil error, other return values ​​are undefined, and these undefined return values ​​should be ignored. However, a small number of functions still return some useful return values ​​when an error occurs. For example, when an error occurs when reading a file, the Read function will return the number of bytes that can be read and the error message. In this case, the read string should be printed out together with the error message.

**Description: There must be a clear description of the return value of the function, so that others can use it. **

 

## 1.7 Correct posture for exception handling

**Position 1: In the program development stage, insist on quick mistakes**

A quick mistake is simply "let it hang". Only when it hangs will you know the mistake the first time. Before the early development and any release stage, the simplest and probably best method is to call the panic function to interrupt the execution of the program to force an error to occur, so that the error will not be ignored and can be repaired as soon as possible.

 

**Posture 2: After the program is deployed, the abnormality should be restored to avoid program termination**

In Golang, if a Goroutine is panic and does not recover, then the entire Golang process will exit abnormally. Therefore, once the Golang program is deployed, the exception that occurs under any circumstances should not cause the program to exit abnormally. We add a delayed recovery call to the upper function to achieve this goal, and whether to recover needs to be based on environment variables or configuration Depending on the file, recover is required by default.
This posture is similar to the assertion in the C language, but there are still differences: Generally, in the Release version, the assertion is defined as null and invalid, but an if check is required for exception protection, although this is not recommended in contract design. In Golang, recover can completely terminate the abnormal expansion process, saving time and effort.

We respond to the exception in the most reasonable way in the delayed function of calling recover:

1. Print the abnormal call information and key business information of the stack so that these problems remain visible;
2. Convert exceptions to errors so that the caller can restore the program to a healthy state and continue to run safely.

Let's look at a simple example:

```go
func funcA() error {
    defer func() {
        if p := recover(); p != nil {
            fmt.Printf("panic recover! p: %v", p)
            debug.PrintStack()
        }
    }()
    return funcB()
}

func funcB() error {
    // simulation
    panic("foo")
    return errors.New("success")
}

func test() {
    err := funcA()
    if err == nil {
        fmt.Printf("err is nil\\n")
    } else {
        fmt.Printf("err is %v\\n", err)
    }
}
```

 

We expect the output of the test function to be:

```
err is foo
```

In fact, the output of the test function is:

```
err is nil
```

 

The reason is that the panic exception handling mechanism does not automatically pass error information to error, so it must be explicitly passed in the funcA function. The code is as follows:

 

```go
func funcA() (err error) {
    defer func() {
        if p := recover(); p != nil {
            fmt.Println("panic recover! p:", p)
            str, ok := p.(string)
            if ok {
                err = errors.New(str)
            } else {
                err = errors.New("panic")
            }
            debug.PrintStack()
        }
    }()
    return funcB()
}
```

 

**Position 3: Use exception handling for branches that should not appear**

 When certain scenarios that shouldn't happen, we should call the panic function to trigger an exception. For example, when the program reaches a path that is logically impossible:

```go
switch s := suit(drawCard()); s {
    case "Spades":
    // ...
    case "Hearts":
    // ...
    case "Diamonds":
    // ... 
    case "Clubs":
    // ...
    default:
        panic(fmt.Sprintf("invalid suit %v", s))
}
```

 

**Position 4: Use panic design for functions that should not be problematic when entering parameters**

There should be no problems when entering parameters. Generally, it refers to hard coding. Let's first look at these two functions (Compile and MustCompile). The MustCompile function is a wrapper for the Compile function:

```go
func MustCompile(str string) *Regexp {
    regexp, error := Compile(str)
    if error != nil {
        panic(`regexp: Compile(` + quote(str) + `): `+ error.Error())
    }
    return regexp
}
```

Therefore, in the case of supporting both user input scenes and hard-coded scenes, the functions that generally support hard-coded scenes are packages of functions that support user-input scenes.
For situations that only support a single hard-coded scenario, panic is used directly in the function design, that is, there will be no error in the return value type list, which makes the function call processing very convenient (no boring "if err != nil {/ print && Error handling/}" code block).



Part of the content of this article is quoted from https://www.jianshu.com/p/f30da01eea97







Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_advanced

