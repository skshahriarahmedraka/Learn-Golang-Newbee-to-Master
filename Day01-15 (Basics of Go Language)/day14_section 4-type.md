

# type keyword

> @author：Han Ru
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.



Type is an important and commonly used keyword in Go syntax, and type does not just correspond to the typedef in C/C++. Knowing the use of type makes it easy to understand the use of core concepts such as struct, interface, and functions in the go language.



## One, type definition

### 1.1 Define structure

Use type to define the structure type:

```go
//1, define the structure
//Structure definition
type person struct {
   name string //Be careful not to have a comma after it
   age int
}
```



### 1.2 Define the interface

Use type to define the interface type:

```go
type USB interface {
	start()
	end()
}
```



### 1.3 Define other new types

Using type, you can also define new types.

grammar:

```go
type Type name Type
```



Sample code:

```go
package main

import "fmt"

type myint int
type mystr string

func main() {

	 var i1 myint
	 var i2 = 100
	 i1 = 100
	 fmt.Println(i1)
	 //i1 = i2 //cannot use i2 (type int) as type myint in assignment
	 fmt.Println(i1,i2)
	 
	 var name mystr
	 name = "Wang Ergou"
	 var s1 string
	 s1 = "Li Xiaohua"
	 fmt.Println(name)
	 fmt.Println(s1)
	 name = s1 //cannot use s1 (type string) as type mystr in assignment
}

```

### 1.4 Defining the type of function

Go language supports functional programming and can use high-level programming syntax. A function can be used as a parameter of another function, or as a return value of another function. When defining this higher-order function, if the type of the function is more complicated, we can use type to define the type of this function:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	 res1 := fun1()
	 fmt.Println(res1(10,20))
}


type my_fun func (int,int)(string)

//The return value of the fun1() function is of type my_func
func fun1 () my_fun{
	fun := func(a,b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

```



## Two, type alias

The type alias is written as:

```go
type alias = Type
```

Type alias regulations: TypeAlias ​​is just an alias of Type. In essence, TypeAlias ​​and Type are the same type. Just like a child who had a nickname and a baby name when he was a child, he would use his scientific name after school, and the English teacher would give him an English name, but these names all refer to him.



Type alias is a new feature added in Go 1.9 version. Mainly used for type compatibility issues in code upgrades and migrations. In C/C++ language, code reconstruction and upgrading can use macros to quickly define a new piece of code. There is no option to add macros in the Go language, but will solve the most troublesome type name change problem in refactoring.

The code for the built-in type definition before Go 1.9 was written like this:

```go
type byte uint8
type rune int32
```

And after Go 1.9 version becomes:

```go

type byte = uint8
type rune = int32
```

This modification is the modification made in conjunction with the type alias.

Sample code:

```go
package main

import (
	"fmt"
)

func main() {

	var i1 myint
	var i2 = 100
	i1 = 100
	fmt.Println(i1)
	//i1 = i2 //cannot use i2 (type int) as type myint in assignment
	fmt.Println(i1,i2)
	var i3 myint2
	i3 = i2
	fmt.Println(i1,i2,i3)

}

type myint int
type myint2 = int //Not redefining the type, just aliasing int


```



## Three, non-local types cannot define methods

Being able to name various types at will, does it mean that you can add methods for these types in your own package?

```go
package main
import (
    "time"
)
// Define the alias of time.Duration as MyDuration
type MyDuration = time.Duration
// add a function to MyDuration
func (m MyDuration) EasySet(a string) {//cannot define new methods on non-local type time.Duration
}
func main() {
}
```

The above code reports an error. Error message: cannot define new methods on non-local type time.Duration

Compiler hint: You cannot define a new method on a non-local type time.Duration. The non-native method refers to the package where the code that uses time.Duration is located, that is, the main package. Because time.Duration is defined in the time package, it is used in the main package. The time.Duration package is not in the same package as the main package, so you cannot define methods for types that are not in the same package.

There are two ways to solve this problem:

-Change the type alias to type definition: type MyDuration time.Duration, that is, change MyDuration from alias to type.
-Put the alias definition of MyDuration in the time package.

## Fourth, use aliases when embedding structure members

What happens when a type alias is used as an embedded member of a structure?

```go
package main

import (
	"fmt"
)

type Person struct {
	name string
}

func (p Person) Show() {
	fmt.Println("Person-->",p.name)
}

//Type alias
type People = Person

type Student struct {
	// Embed two structures
	Person
	People
}

func (p People) Show2(){
	fmt.Println("People------>",p.name)
}

func main() {
	//
	var s Student

	//s.name = "王二狗" //ambiguous selector s.name
	s.People.name = "Li Xiaohua"
	s.Person.name = "Wang Ergou"
	//s.Show() //ambiguous selector s.Show
	s.Person.Show()
	s.People.Show2()
	fmt.Printf("%T,%T\n",s.Person,s.People) //main.Person,main.Person

}

```

When directly accessing the name through s, or s directly calls the Show() method, because both types have a name field and a Show() method, ambiguity will occur, which proves that the essence of People is indeed the Person type.



Part of the content is quoted from: http://c.biancheng.net/view/25.html







Qianfeng Go language learning group: 784190273

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation

