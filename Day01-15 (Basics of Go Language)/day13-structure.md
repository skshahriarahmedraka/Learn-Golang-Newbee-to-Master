

# One, structure

> @author：Han Ru
>
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## 1.1 What is a structure

Arrays in Go language can store the same type of data, but in the structure we can define different data types for different items.
A structure is a data collection composed of a series of data of the same type or different types.

## 1.2 Definition and initialization of structure

```go
type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}
```

Once the structure type is defined, it can be used for variable declarations

```go
variable_name := structure_variable_type {value1, value2...valuen}
```

**Initialization structure**

```go
// 1. Provide initialization values ​​in order
P := person{"Tom", 25}
// 2. Initialize by way of field: value, so it can be in any order
P := person{age:24, name:"Tom"}
// 3.new method, if the initial value is not set, the default initial value of the type will be given
p := new(person)
p.age=24
```

## 1.3 Structure access

Access to structure members (access to various fields of the structure)

The dot. operator is used to access the various fields of the structure.

```go
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books /* declare Book1 as Books type*/
   var Book2 Books /* declare Book2 as Books type*/

   /* book 1 description*/
   Book1.title = "Go Language"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go Language Tutorial"
   Book1.book_id = 6495407

   /* book 2 description*/
   Book2.title = "Python Tutorial"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python language tutorial"
   Book2.book_id = 6495700

   /* Print Book1 information*/
   fmt.Printf( "Book 1 title: %s\n", Book1.title)
   fmt.Printf( "Book 1 author: %s\n", Book1.author)
   fmt.Printf( "Book 1 subject: %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id: %d\n", Book1.book_id)

   /* Print Book2 information*/
   fmt.Printf( "Book 2 title: %s\n", Book2.title)
   fmt.Printf( "Book 2 author: %s\n", Book2.author)
   fmt.Printf( "Book 2 subject: %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id: %d\n", Book2.book_id)
}
```

operation result:

```go
Book 1 title: Go language
Book 1 author: www.runoob.com
Book 1 subject: Go language tutorial
Book 1 book_id: 6495407
Book 2 title: Python tutorial
Book 2 author: www.runoob.com
Book 2 subject: Python language tutorial
Book 2 book_id: 6495700
```

## 1.4 Structure pointer

Pointer to a structure
You can also create pointers to structures.

**Structure pointer**

```go
var struct_pointer *Books
```

The pointer variable defined above can store the address of the structure variable. View the address of the structure variable, you can place the & symbol in front of the structure variable

```go
struct_pointer = &Book1;
```

Use structure pointers to access structure members, use the "." operator

```go
struct_pointer.title;
```

```go
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books /* Declare Book1 of type Book */
   var Book2 Books /* Declare Book2 of type Book */

   /* book 1 description*/
   Book1.title = "Go Language"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go Language Tutorial"
   Book1.book_id = 6495407

   /* book 2 description*/
   Book2.title = "Python Tutorial"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python language tutorial"
   Book2.book_id = 6495700

   /* Print Book1 information*/
   printBook(&Book1)

   /* Print Book2 information*/
   printBook(&Book2)
}
func printBook( book *Books) {
   fmt.Printf( "Book title: %s\n", book.title);
   fmt.Printf( "Book author: %s\n", book.author);
   fmt.Printf( "Book subject: %s\n", book.subject);
   fmt.Printf( "Book book_id: %d\n", book.book_id);
}
```

Structure instantiation can also be like this

```go
package main

import "fmt"

type Books struct {
}

func (s Books) String() string {
	return "data"
}
func main() {
	fmt.Printf("%v\n", Books{})
}
```



## 1.5 Anonymous field of structure

**Anonymous field of structure**

You can use fields to create structures. These fields only contain a type without a field name. These fields are called anonymous fields.

In the type, use the method of not writing the field name, and use another type

```go
type Human struct {
    name string
    age int
    weight int
} 
type Student struct {
    Human // Anonymous field, then the default Student contains all fields of Human
    speciality string
} 
func main() {
    // we initialize a student
    mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
    // we visit the corresponding field
    fmt.Println("His name is ", mark.name)
    fmt.Println("His age is ", mark.age)
    fmt.Println("His weight is ", mark.weight)
    fmt.Println("His speciality is ", mark.speciality)
    // Modify the corresponding remarks
    mark.speciality = "AI"
    fmt.Println("Mark changed his speciality")
    fmt.Println("His speciality is ", mark.speciality)
    // modify his age information
    fmt.Println("Mark become old")
    mark.age = 46
    fmt.Println("His age is", mark.age)
    // modify his weight information
    fmt.Println("Mark is not an athlet anymore")
    mark.weight += 60
    fmt.Println("His weight is", mark.weight)
}
```

> You can use the "." method to call the attribute value in the anonymous field
>
> Actually is the inheritance of the field
>
> Among them, anonymous fields can be understood as the same field name and field type
>
> Based on the above understanding, it is possible to `mark.Human = Human{"Marcus", 55, 220}` and `mark.Human.age -= 1`
>
> If there is a field in an anonymous field with the same name as a non-anonymous field, the outermost layer has priority access, the principle of proximity

Accessing and modifying fields anonymously is quite useful, but not only struct fields, all built-in types and custom types can be used as anonymous fields.



## 1.6 Structure nesting

Nested structure
A structure may contain a field, and this field in turn is a structure. These structures are called nested structures.

Sample code:

```go
package main

import (  
    "fmt"
)

type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age int
    address Address
}

func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.address = Address {
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:",p.age)
    fmt.Println("City:",p.address.city)
    fmt.Println("State:",p.address.state)
}
```



## 1.7 Promotion field

The fields in the structure that belong to the anonymous structure are called promoted fields, because they can be accessed as if they belong to a structure with anonymous structure fields. Understanding this definition is quite complicated.

Sample code:

```go
package main

import (  
    "fmt"
)

type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age int
    Address
}

func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.Address = Address{
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.city) //city is promoted field
    fmt.Println("State:", p.state) //state is promoted field
}
```

operation result

```
Name: Naveen  
Age: 50  
City: Chicago  
State: Illinois
```



## 1.8 Export structure and fields

If the structure type starts with a capital letter, then it is an exported type and can be accessed from other packages. Similarly, if the fields of the structure begin with uppercase, they can be accessed from other packages.

Sample code:

1. In the computer directory, create the file spec.go

```go
package computer

type Spec struct {//exported struct  
    Maker string //exported field
    model string //unexported field
    Price int //exported field
}
```

2. Create the main.go file

```go
package main

import "structs/computer"  
import "fmt"

func main() {  
    var spec computer.Spec
    spec.Maker = "apple"
    spec.Price = 50000
    fmt.Println("Spec:", spec)
}
```

> The directory structure is as follows:
>
> src  
> structs
> computer
> spec.go
> main.go



## 1.9 Structure comparison

The structure is a value type, if each field is comparable, it is comparable. If their corresponding fields are equal, then the two structure variables are considered equal.

Sample code:

```go
package main

import (  
    "fmt"
)

type name struct {  
    firstName string
    lastName string
}


func main() {  
    name1 := name{"Steve", "Jobs"}
    name2 := name{"Steve", "Jobs"}
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }

    name3 := name{firstName:"Steve", lastName:"Jobs"}
    name4 := name{}
    name4.firstName = "Steve"
    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
    }
}
```

operation result

```
name1 and name2 are equal  
name3 and name4 are not equal  
```

**If the field contained in the structure variable is not comparable, then the structure variable is not comparable**

Sample code:

```go
package main

import (  
    "fmt"
)

type image struct {  
    data map[int]int
}

func main() {  
    image1 := image{data: map[int]int{
        0: 155,
    }}
    image2 := image{data: map[int]int{
        0: 155,
    }}
    if image1 == image2 {
        fmt.Println("image1 and image2 are equal")
    }
}
```



## 2.0 Structure as a function parameter

Structure as a function parameter

```go
ackage main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books /* declare Book1 as Books type*/
   var Book2 Books /* declare Book2 as Books type*/

   /* book 1 description*/
   Book1.title = "Go Language"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go Language Tutorial"
   Book1.book_id = 6495407

   /* book 2 description*/
   Book2.title = "Python Tutorial"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python language tutorial"
   Book2.book_id = 6495700

   /* Print Book1 information*/
   printBook(Book1)

   /* Print Book2 information*/
   printBook(Book2)
}

func printBook( book Books) {
   fmt.Printf( "Book title: %s\n", book.title);
   fmt.Printf( "Book author: %s\n", book.author);
   fmt.Printf( "Book subject: %s\n", book.subject);
   fmt.Printf( "Book book_id: %d\n", book.book_id);
}
```

**make, new operation**

make is used for memory allocation of built-in types (map, slice, and channel). new is used for various types of memory allocation
The built-in function new is essentially the same as the functions of the same name in other languages: new(T) allocates a zero-filled memory space of type T, and returns its address, which is a value of type *T. In Go terms, it returns a pointer to the newly allocated zero value of type T. One thing is very important: new returns a pointer

The built-in function make(T, args) has different functions from new(T). Make can only create slices, maps, and channels, and returns a T type with an initial value (non-zero) instead of *T. Essentially, the reason why these three types are different is that references to data structures must be initialized before they are used. For example, a slice is a three-item descriptor containing pointers to data (internal array), length, and capacity; before these items are initialized, the slice is nil. For slice, map, and channel, make initializes the internal data structure and fills in appropriate values.

make returns the initialized (non-zero) value.





Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_advanced

