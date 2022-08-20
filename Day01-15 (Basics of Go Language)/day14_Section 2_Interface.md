

# Interface

> @author：Han Ru
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## 1.1 What is an interface?

The general definition of an interface in the object-oriented world is "the interface defines the behavior of an object". It indicates what the specified object should do. The method (implementation details) to achieve this behavior is specific to the object.

In Go, an interface is a set of method signatures. When a type provides definitions for all methods in an interface, it is said to implement an interface. It is very similar to OOP. The interface specifies the methods that the type should have, and the type determines how to implement these methods.

> It defines all common methods together, any other type as long as they implement these methods is to implement this interface
>
> An interface defines a set of methods. If an object implements all methods of an interface, the object implements the interface.



## 1.2 Interface definition syntax

Define the interface

```go
/* Define interface*/
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}

/* Define structure */
type struct_name struct {
   /* variables */
}

/* Implement interface method*/
func (struct_name_variable struct_name) method_name1() [return_type] {
   /* Method implementation*/
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
   /* Method implementation*/
}
```

Sample code:

```go
package main

import (
    "fmt"
)

type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func main() {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()

}
```

operation result:

```go
I am Nokia, I can call you!
I am iPhone, I can call you!
```

-interface can be implemented by any object
-An object can implement any number of interfaces
-Any type implements an empty interface (we define it like this: interface{}), which is an interface containing 0 methods

## 1.3 interface value



```go
package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human //Anonymous field
	school string
	loan float32
}
type Employee struct {
	Human //Anonymous field
	company string
	money float32
} //Human implements the Sayhi method
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
} //Human implements the Sing method
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
} //Employee rewrites the SayHi method of Human
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

// Interface Men is implemented by Human, Student and Employee
// Because these three types implement these two methods
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}
	//Define variable i of type Men
	var i Men
	//i can store Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")
	//i can also store Employee
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")
	// defines slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//T These three are different types of elements, but they implement the same interface interface
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x {
		value.SayHi()
	}
}
```

operation result:

```go
	This is Mike, a Student:
	Hi, I am Mike you can call me on 222-222-XXX
	La la la la... November rain
	This is Tom, an Employee:
	Hi, I am Sam, I work at Things Ltd.. Call me on 444-222-XXX
	La la la la... Born to be wild
	Let's use a slice of Men and see what happens
	Hi, I am Paul you can call me on 111-222-XXX
	Hi, I am Sam, I work at Golang Inc.. Call me on 444-222-XXX
	Hi, I am Mike you can call me on 222-222-XXX
```

So what value can be stored in the interface? If we define an interface variable, then this variable can store any type of object that implements this interface. For example, in the above example, we define a variable m of Men interface type, then the value of Human, Student or Employee can be stored in m

> Of course, it is also possible to use pointers
>
> However, the interface object cannot call the properties of the implementation object

**interface function parameters**

Interface variables can hold any object that implements the interface type. This provides us with some additional thinking about writing functions (including methods). Can we define interface parameters to allow functions to accept various types of parameters?

**Embed in interface**

```go
package main

import "fmt"

type Human interface {
	Len()
}
type Student interface {
	Human
}

type Test struct {
}

func (h *Test) Len() {
	fmt.Println("Success")
}
func main() {
	var s Student
	s = new(Test)
	s.Len()
}
```



Sample code:

```go
package test

import (
	"fmt"
)

type Controller struct {
	M int32
}

type Something interface {
	Get()
	Post()
}

func (c *Controller) Get() {
	fmt.Print("GET")
}

func (c *Controller) Post() {
	fmt.Print("POST")
}
```

```go
package main

import (
	"fmt"
	"test"
)

type T struct {
	test.Controller
}

func (t *T) Get() {
	//new(test.Controller).Get()
	fmt.Print("T")
}
func (t *T) Post() {
	fmt.Print("T")
}
func main() {
	var something test.Something
	something = new(T)
	var t T
	tM = 1
	// t.Controller.M = 1
	something.Get()
}
```



The Controller implements all the Something interface methods. When the Controller structure is called in the structure T, T is equivalent to the inheritance in Java, and T inherits the Controller. Therefore, T does not need to rewrite all the methods in the Something interface. Because the parent constructor has already implemented the interface.

If the Controller does not implement the Something interface method, then T must implement all of its methods to call the method in Something.

If `something = new(test.Controller)`, the Get method in the Controller is called.

T can use variables defined in the Controller structure



## 1.4 Types of interfaces

**Interface and duck type:**

First look directly at the definition in Wikipedia:

> If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.

The translation is: if something looks like a duck, swims like a duck, and quacks like a duck, then it can be regarded as a duck.

Duck Typing, duck typing, is an object inference strategy of dynamic programming languages. It pays more attention to how the object can be used, rather than the type of the object itself. As a static language, Go language perfectly supports duck types through interfaces.

In static languages ​​such as Java and C++, an interface must be explicitly declared to be implemented before it can be used wherever this interface is needed. If you call a certain number in a program, but pass in a type that does not implement the other at all, it will not be passed during the compilation phase. This is why static languages ​​are safer than dynamic languages.

The difference between dynamic language and static language is reflected here. Static languages ​​can find type mismatch errors during compilation. Unlike dynamic languages, you must run to that line of code to report an error. Of course, static languages ​​require programmers to write programs in accordance with regulations during the coding phase, and specify data types for each variable. This increases the workload to a certain extent and lengthens the amount of code. Dynamic languages ​​do not have these requirements, which can allow people to focus more on the business, the code is also shorter, and it is faster to write. This point is clear to the students who write python.

As a modern static language, Go language has latecomer advantages. It introduces the convenience of dynamic language, and at the same time performs type checking of static language. It is very happy to write. Go adopts a compromise approach: it does not require the type to explicitly declare that an interface is implemented, as long as the relevant method is implemented, the compiler can detect it.

To sum up, the duck type is a dynamic language style. In this style, the effective semantics of an object is not inherited from a specific class or implements a specific interface, but by its "collection of current methods and properties." Decide. As a static language, Go implements the duck type through the interface. In fact, the Go compiler does the hidden conversion work in it.

**Polymorphism of Go language:**

Polymorphism in Go is achieved with the help of interfaces. As we have already discussed, interfaces can be implemented implicitly in Go. If the type provides definitions for all methods declared in the interface, then an interface is implemented. Let's see how to implement polymorphism with the help of interfaces.

Any type that defines all methods of an interface is said to implicitly implement the interface.

Variables of a type interface can hold any value that implements the interface. This attribute of the interface is used to implement polymorphism in Go.



## 1.5 Interface Assertion

As mentioned earlier, because the empty interface interface{} does not define any functions, all types in Go implement empty interfaces. When the formal parameter of a function is interface{}, then in the function, the formal parameter needs to be asserted to get its real type.

Syntax format:

```go
// safe type assertion


   <value of target type>
    ,
    <Boolean parameter>
      := 
     <expression>
      .(Target type) //non-safe type assertion 
      <value of target type>
        := 
       <expression>
        .(Target type) ``` Sample code: ```go package main import "fmt" func main() {var i1 interface{} = new (Student) s := i1.(Student) //Unsafe, if If the assertion fails, it will directly panic fmt.Println(s) var i2 interface{} = new(Student) s, ok := i2.(Student) //Safe, the assertion fails, and it will not panic, but the value of ok is false If ok {fmt.Println(s)}} type Student struct {} ``` There is actually another form of assertion, which is used to determine the type of interface by using the switch statement. Each case will be considered sequentially. When a case is hit, the statements in the case will be executed, so the order of the case statements is very important, because there are likely to be multiple cases matching. Sample code: ```go switch ins:=s.(type) {case Triangle: fmt.Println("Triangle...",ins.a,ins.b,ins.c) case Circle: fmt.Println( "Circle...",ins.radius) case int: fmt.Println("Integer data...")} ``` #### Summary interface object cannot call the properties of the interface implementation object Qianfeng Go language Learning group: 784190273 Corresponding video address: https://www.bilibili.com/video/av56018934 https://www.bilibili.com/video/av47467197 Source code: https://github.com/rubyhan1314/go_foundation
       
      
     
    
   

