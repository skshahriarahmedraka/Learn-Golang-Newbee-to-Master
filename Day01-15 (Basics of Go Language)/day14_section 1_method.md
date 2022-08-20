

# method

> @author：Han Ru
>
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## 1.1 What is a method

There are both functions and methods in the Go language. A method is a function that contains a receiver. The receiver can be a value or a pointer of a named type or a structure type. All methods of a given type belong to the method set of that type

A method is just a function with a special receiver type, which is written between the func keyword and the method name. The receiver can be of struct type or non-struct type. The receiver can be accessed inside the method.



Methods can add new behaviors to user-defined types. The difference between it and a function is that a method has a receiver. Add a receiver to a function, then it becomes a method. The receiver can be a value receiver or a pointer receiver.

When calling a method, the value type can either call the method of the value receiver or the method of the pointer receiver; the pointer type can call the method of the pointer receiver or the method of the value receiver.

In other words, regardless of the type of the receiver of the method, the value and pointer of that type can be called, and it does not have to strictly conform to the type of the receiver.



## 1.2 Method syntax

Definition method syntax

```go
func (t Type) methodName(parameter list)(return list) {
  
}
func funcName(parameter list)(return list){
    
}
```

Example code:

```go
package main

import (  
    "fmt"
)

type Employee struct {  
    name string
    salary int
    currency string
}

/*
 displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {  
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {  
    emp1 := Employee {
        name: "Sam Adolf",
        salary: 5000,
        currency: "$",
    }
    emp1.displaySalary() //Calling displaySalary() method of Employee type
}
```

**You can define the same method name**

Sample code:

```go
package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}


func (r Rectangle) area() float64 {
	return r.width * r.height
}
//The method belongs to the method in the Circle type object
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of ​​r1 is: ", r1.area())
	fmt.Println("Area of ​​r2 is: ", r2.area())
	fmt.Println("Area of ​​c1 is: ", c1.area())
	fmt.Println("Area of ​​c2 is: ", c2.area())
}
```

operation result

```
Area of ​​r1 is: 24
Area of ​​r2 is: 36
Area of ​​c1 is: 314.1592653589793
Area of ​​c2 is: 1963.4954084936207
```

-Although the name of the method is exactly the same, if the receiver is different, then the method is different
-The recipient's field can be accessed in the method
-Calling method to access through., just like accessing fields in struct 

## 1.3 Methods and functions

Now that we have functions, why use methods?

Sample code:

```go
package main

import (  
    "fmt"
)

type Employee struct {  
    name string
    salary int
    currency string
}

/*
 displaySalary() method converted to function with Employee as parameter
*/
func displaySalary(e Employee) {  
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {  
    emp1 := Employee{
        name: "Sam Adolf",
        salary: 5000,
        currency: "$",
    }
    displaySalary(emp1)
}
```

> In the above program, the displaySalary method is converted into a function, and Employee struct is passed to it as a parameter. This program also produced the same output: Salary of Sam Adolf is $5000.

Why can we use functions to write the same program? There are several reasons

1. Go is not a purely object-oriented programming language, it does not support classes. Therefore, a method of a type is a way to achieve a behavior similar to a class.
2. Methods with the same name can be defined on different types, and functions with the same name are not allowed. Suppose we have a square and circular structure. You can define a method named Area on squares and circles. This is done in the program below.

## 1.4 Variable scope 

The scope is the scope in the source code of the constant, type, variable, function, or package represented by the declared identifier.

Variables in Go language can be declared in three places:

-Variables defined in functions are called local variables
-Variables defined outside the function are called global variables
-The variables in the function definition are called formal parameters

**Local Variables**

The variables declared in the function body are called local variables, and their scope is only in the function body, and the parameters and return value variables are also local variables.

**Global Variables**

Variables declared outside the function body are called global variables. Global variables with initial capital letters can be used in the entire package or even outside packages (after being exported).

```go
package main

import "fmt"

/* Declare global variables*/
var g int

func main() {

   /* Declare local variables*/
   var a, b int

   /* Initialization parameters*/
   a = 10
   b = 20
   g = a + b

   fmt.Printf("Result: a = %d, b = %d and g = %d\n", a, b, g)
}
```

`Results`

```go
Results: a = 10, b = 20 and g = 30
```

**Formal parameters**

Formal parameters will be used as local variables of the function

**Pointer as receiver**

If the pointer is not used as the receiver, it is actually just a copy, and it cannot really change the data in the receiver.

```go
func (b *Box) SetColor(c Color) {
	b.color = c
}
```

Sample code

```go
package main

import (
	"fmt"
)

type Rectangle struct {
	width, height int
}

func (r *Rectangle) setVal() {
	r.height = 20
}

func main() {
	p := Rectangle{1, 2}
	s := p
	p.setVal()
	fmt.Println(p.height, s.height)
}
```

result

```go
20 2
```

If there is no such *, the value is `2 2`

## 1.5 method inheritance

The method can be inherited. If the anonymous field implements a method, the struct containing the anonymous field can also call the method

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
}
type Employee struct {
	Human //Anonymous field
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}
```

operation result:

```go
Hi, I am Mark you can call me on 222-222-YYYY
Hi, I am Sam you can call me on 111-888-XXXX
```

## 1.6 method rewrite

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
}
type Employee struct {
	Human //Anonymous field
	company string
}

//Human defines method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee's method rewrites Human's method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}
func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}
```

operation result:

```go
Hi, I am Mark you can call me on 222-222-YYYY
Hi, I am Sam, I work at Golang Inc. Call me on 111-888-XXXX
```

-Methods can be inherited and rewritten
-When there is an inheritance relationship, call according to the principle of proximity



Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_advanced

