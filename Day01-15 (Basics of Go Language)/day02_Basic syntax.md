

# Basic syntax-variables

> @author：Han Ru
>
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## One, the use of variables

### 1.1 What is a variable

A variable is a name provided to a memory location for storing a specific type of value. There are multiple syntaxes for declaring variables in Go.

So the essence of the variable is a small piece of memory, used to store data, the value can be changed during the running of the program



### 1.2 Declaring variables

The var name type is the syntax for declaring a single variable.

> Start with a letter or underscore and consist of one or more letters, numbers, and underscores

Declare a variable

The first one is to specify the variable type, and if no value is assigned after declaration, the default value is used

```go
var name type
name = value
```

The second is to determine the type of the variable based on the value (Type inference)

If a variable has an initial value, Go will automatically be able to use the initial value to infer the type of the variable. Therefore, if the variable has an initial value, you can omit the type in the variable declaration.

```go
var name = value
```

The third is to omit var. Note: The variable on the left side should not have been declared (when multiple variables are declared at the same time, at least one is guaranteed to be a new variable), otherwise it will cause a compilation error (short declaration)



```go
name := value

// E.g
var a int = 10
var b = 10
c: = 10
```

> In this way, it can only be used in the function body, and cannot be used for the declaration and assignment of global variables

Sample code:

```go
package main
var a = "Hello"
var b string = "World"
var c bool

func main(){
    println(a, b, c)
}
```

operation result:

```go
Hello World false
```

#### Multiple variable declaration

The first type, separated by commas, separates declaration and assignment. If there is no assignment, there is a default value

```go
var name1, name2, name3 type
name1, name2, name3 = v1, v2, v3
```

The second, direct assignment, the following variable types can be different types

```go
var name1, name2, name3 = v1, v2, v3
```

The third type, collection type

```go
var (
    name1 type1
    name2 type2
)
```

### 1.3 Matters needing attention

-Variables must be defined before they can be used
-The go language is a static language, and the type of the variable and the type of the assignment must be consistent.
-Variable names cannot conflict. (The same role cannot conflict in the domain)
-Short definition, at least one of the variable names on the left is new
-Short definition method, global variables cannot be defined.
-The zero value of the variable. Also called the default value.
-The variable must be used if it is defined, otherwise it will fail to compile.

If in the same code block, we can't use the initialization declaration again for the variable with the same name, for example: a := 20 is not allowed, the compiler will prompt the error no new variables on left side of :=, but a = 20 is ok, because this is to assign a new value to the same variable.

If you use it before defining variable a, you will get a compilation error undefined: a. If you declare a local variable but do not use it in the same code block, you will also get a compilation error, such as the variable a in the following example:

```go
func main() {
   var a string = "abc"
   fmt.Println("hello, world")
}
```

Trying to compile this code will get the error a declared and not used

In addition, simply assigning a value to a is not enough, this value must be used, so use

In the same scope, if a variable with the same name already exists, the subsequent declaration initialization will degenerate into an assignment operation. But the premise is that at least one new variable must be defined and in the same scope. For example, the following y is the newly defined variable

```go
package main

import (
	"fmt"
)

func main() {
	x := 140
	fmt.Println(&x)
	x, y := 200, "abc"
	fmt.Println(&x, x)
	fmt.Print(y)
}
```

operation result:

```go
0xc04200a2b0
0xc04200a2b0 200
abc
```





# Basic syntax-constant

## One, the use of constants

### 1.1 Constant declaration

A constant is an identifier of a simple value, an amount that will not be modified when the program is running.

```go
const identifier [type] = value
```

```go
Explicit type definition: const b string = "abc"
Implicit type definition: const b = "abc"
```
```go
package main

import "fmt"

func main() {
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int
   const a, b, c = 1, false, "str" ​​//Multiple assignment

   area = LENGTH * WIDTH
   fmt.Printf("Area is: %d", area)
   println()
   println(a, b, c)   
}
```
operation result:

```go
Area: 50
1 false str
```

Constants can be used as enumerations, constant groups

```go
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```
If the type and initialization value are not specified in the constant group, it will be the same as the right value of the non-empty constant in the previous line

```go
package main

import (
	"fmt"
)

func main() {
	const (
		x uint16 = 16
		y
		s = "abc"
		z
	)
	fmt.Printf("%T,%v\n", y, y)
	fmt.Printf("%T,%v\n", z, z)
}
```
operation result:

```go
uint16,16
string,abc
```

Notes on constants:

-The data types in constants can only be boolean, numeric (integer, floating-point, and complex) and string types

-Unused constants will not report errors when compiling

-When displaying the specified type, you must ensure that the left and right value types of the constant are consistent, and display type conversion can be done if necessary. This is different from variables. Variables can be of different types.



### 1.2 iota

iota, a special constant, can be considered a constant that can be modified by the compiler

iota can be used as an enumeration value:

```go
const (
    a = iota
    b = iota
    c = iota
)
```
The first iota is equal to 0. Whenever iota is used in a new line, its value is automatically increased by 1; so a=0, b=1, c=2 can be abbreviated as follows:

```go
const (
    a = iota
    b
    c
)
```
**iota usage**

```go
package main

import "fmt"

func main() {
    const (
            a = iota //0
            b //1
            c //2
            d = "ha" //Independent value, iota += 1
            e //"ha" iota += 1
            f = 100 //iota +=1
            g //100 iota +=1
            h = iota //7, restore count
            i //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}
```
operation result:

```
0 1 2 ha ha 100 100 7 8
```

If the self-increment of iota is interrupted, it must be restored explicitly. And the subsequent self-increasing value increases in line order

The self-increment default is int type, you can display the specified type by yourself

Digital constants will not allocate storage space, and there is no need to obtain values ​​through memory addressing like variables, so the address cannot be obtained









Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation

