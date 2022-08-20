

# One, the concept of function

> @author：Han Ru
> 
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.



## 1.1 What is a function

Functions are blocks of code that perform specific tasks.

## 1.2 Function declaration

Go language has at least one main function

Syntax format:

```go
func funcName(parametername type1, parametername type2) (output1 type1, output2 type2) {
//Here is the processing logic code
//Return multiple values
return value1, value2
}
```

-func: The function is declared by func
-funcName: function name, function name and parameter list together constitute the function signature.
-parametername type: parameter list, the parameter is like a placeholder, when the function is called, you can pass the value to the parameter, this value is called the actual parameter. The parameter list specifies the parameter type, order, and number of parameters. Parameters are optional, which means that the function does not need to contain parameters.
-output1 type1, output2 type2: return type, the function returns a list of values. return_types is the data type of the column value. Some functions do not require a return value, in this case return_types is not necessary.
-The return value above declares two variables output1 and output2. If you don't want to declare it, you can just use the two types.
-If there is only one return value and the return value variable is not declared, then you can omit the parentheses that include the return value (that is, a return value may not declare the return type)
-Function body: the code collection of the function definition.

## 1.3 Use of functions

Sample code:

```go
package main

import "fmt"

func main() {
   /* Define local variables*/
   var a int = 100
   var b int = 200
   var ret int

   /* Call the function and return the maximum value*/
   ret = max(a, b)

   fmt.Printf( "Maximum value is: %d\n", ret)
}

/* The function returns the maximum value of two numbers*/
func max(num1, num2 int) int {
   /* Define local variables*/
   var result int

   if (num1> num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}
```

operation result:

```go
The maximum value is: 200
```





# Two, the parameters of the function

## 2.1 The use of parameters

Formal parameters: When defining a function, it is used to receive data passed in from the outside, called formal parameters, or formal parameters for short.

Actual parameter: The actual data passed to the formal parameter when the function is called is called actual parameter, or actual parameter for short.

Function call:

	A: The function name must match
	
	B: The actual participating formal parameters must correspond one-to-one: order, number, type

## 2.2 Variable parameters

Go functions support variable parameters. Functions that accept variable parameters have an indefinite number of parameters. In order to do this, you first need to define the function to accept variable parameters:

```go
func myfunc(arg ...int) {}
```

`arg...int` tells Go that this function accepts an indefinite number of parameters. Note that the types of these parameters are all int. In the function body, the variable arg is a slice of int:

```go
for _, n := range arg {
fmt.Printf("And the number is: %d\n", n)
}
```

## 2.3 Parameter passing

The parameters of go language functions also exist **value transfer** and **reference transfer**

Function application scenarios

**Pass by value**

```go
package main

import (
   "fmt"
   "math"
)

func main(){
   /* Declare function variables*/
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* Use function*/
   fmt.Println(getSquareRoot(9))

}
```

**Pass by reference**

This involves the so-called pointer. We know that variables are stored at a certain address in memory, and modifying variables is actually modifying the contents of the variable address.
live. Only when the add1 function knows the address of the x variable can the value of the x variable be modified. Therefore, we need to pass the address &x of x to the function, and change the type of the function's parameter from int to *int, that is, to the pointer type, in order to modify the value of the x variable in the function. At this time, the parameters are still passed by copy, but the copy is a pointer. Please see the example below

```go
package main
import "fmt"
//A simple function that implements the operation of parameter +1
func add1(a *int) int {// Please note,
*a = *a+1 // modify the value of a
return *a // return new value
} f
unc main() {
x := 3
fmt.Println("x = ", x) // should output "x = 3"
x1 := add1(&x) // call add1(&x) to pass the address of x
fmt.Println("x+1 = ", x1) // should output "x+1 = 4"
fmt.Println("x = ", x) // should output "x = 4"
}
```

-Passing pointers allows multiple functions to operate on the same object.
-Passing pointers is relatively lightweight (8bytes), just pass the memory address, we can use pointers to pass large structures. If you use parameter values ​​to pass, relatively more system overhead (memory and time) will be spent on each copy. So when you want to pass large structures, using pointers is a wise choice.
-**The implementation mechanisms of the three types of slice and map in Go are similar to pointers**, so they can be passed directly instead of passing the pointer after taking the address. (Note: If the function needs to change the length of the slice, it still needs to take the address to pass the pointer)



# Three, the return value of the function

## 3.1 What is the return value of a function

After a function is called, the execution result returned to the caller is called the return value of the function.

The caller needs to use a variable to receive the result

## 3.2 A function can return multiple values

A function can have no return value, one return value, or multiple values.

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Mahesh", "Kumar")
   fmt.Println(a, b)
}
```

```go
func SumAndProduct(A, B int) (add int, Multiplied int) {
add = A+B
Multiplied = A*B
return
}
```
## 3.3 Blank identifier

_ Is a blank identifier in Go. It can replace any value of any type. Let us look at the usage of this blank identifier.

For example, the result returned by the rectProps function is the area and perimeter. If we only need the area but not the perimeter, we can use a blank identifier.

Sample code:

```go
package main

import (  
    "fmt"
)

func rectProps(length, width float64) (float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}
func main() {  
    area, _ := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f ", area)
}
```



# Fourth, the scope of the function

Scope: The range in which the variable can be used.

## 4.1 Local Variables

A variable defined inside a function is called a local variable

Where the variable is defined, it can only be used in which scope, beyond this scope, we think that the variable will be destroyed.

## 4.2 Global Variables

A variable defined outside a function is called a global variable

All functions can be used, and share this data



# Five, the nature of the function

A function is also a data type in the Go language, which can be used as a parameter of another function, or as a return value of another function.

# Six, defer function

## 6.1 What is delay?

That is, the defer statement, the defer statement is used to execute a function call, before the function, the defer statement returns.

## 6.2 Delay function

You can add multiple defer statements in the function. When the function is executed to the end, these defer statements will be executed in reverse order, and finally the function returns. Especially when you are doing some operations to open resources, you need to return in advance if you encounter an error. You need to close the corresponding resources before returning, otherwise it is easy to cause resource leakage and other problems

-If there are many calls to defer, then defer uses the `last in, first out` mode
-When leaving the method, execute (it will also execute when an error is reported)

```go
func ReadWrite() bool {
    file.Open("file")
    defer file.Close()
    if failureX {
          return false
    } i
    f failureY {
          return false
    } 
    return true
}
```

Finally execute `file.Close()`

Sample code:

```go
package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer fmt.Println(b)
	fmt.Println(a)
}
```

operation result:

```go
1
2
```

Sample code:

```go
package main

import (  
    "fmt"
)

func finished() {  
    fmt.Println("Finished finding largest")
}

func largest(nums []int) {  
    defer finished()    
    fmt.Println("Started finding largest")
    max := nums[0]
    for _, v := range nums {
        if v> max {
            max = v
        }
    }
    fmt.Println("Largest number in", nums, "is", max)
}

func main() {  
    nums := []int{78, 109, 2, 563, 300}
    largest(nums)
}
```

operation result:

```
Started finding largest  
Largest number in [78 109 2 563 300] is 563  
Finished finding largest 
```

## 6.3 Delay method

Delay is not limited to functions. It is perfectly legal to delay a method call. Let's write a small program to test this.

Sample code:

```go
package main

import (  
    "fmt"
)


type person struct {  
    firstName string
    lastName string
}

func (p person) fullName() {  
    fmt.Printf("%s %s",p.firstName,p.lastName)
}

func main() {  
    p := person {
        firstName: "John",
        lastName: "Smith",
    }
    defer p.fullName()
    fmt.Printf("Welcome ")  
}
```

operation result:

```
Welcome John Smith 
```

## 6.4 Delay parameters

The parameters of the delayed function are executed when the delayed statement is executed, not when the actual function call is executed.

Let us understand this problem through an example.

Sample code:

```go
package main

import (  
    "fmt"
)

func printA(a int) {  
    fmt.Println("value of a in deferred function", a)
}
func main() {  
    a := 5
    defer printA(a)
    a = 10
    fmt.Println("value of a before deferred function call", a)

}
```

operation result:

```
value of a before deferred function call 10  
value of a in deferred function 5 
```

## 6.5 Postponement of the stack

When a function has multiple delayed calls, they are added to a stack and executed in Last In First Out (LIFO) order.

We will write a small program that uses a bunch of defers to print a string. Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    name := "Naveen"
    fmt.Printf("Orignal String: %s\n", string(name))
    fmt.Printf("Reversed String: ")
    for _, v := range []rune(name) {
        defer fmt.Printf("%c", v)
    }
}
```

operation result:

```
Orignal String: Naveen  
Reversed String: neevaN 
```



## 6.6 defer note

```
Defer function:
When the statement in the peripheral function is executed normally, only when all the delayed functions in it are executed, the peripheral function will truly end its execution.
When the return statement in the peripheral function is executed, the peripheral function will only return after all the delayed functions in it are executed.
When the code in the peripheral function triggers a run panic, the run-time panic will only be extended to the calling function only after all the delayed functions are executed.
```



Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation






