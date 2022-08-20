

# Object Oriented (OOP)

> @author：Han Ru
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.



Go is not a pure object-oriented programming language. Object-oriented in go, the structure replaces the class.

Go does not provide a class, but it provides a structure struct, method method, which can be added to the structure. Provides the behavior of binding data and methods, which are similar to classes.

## 1.1 Define structure and method

To understand better through the following code, first create a package in the src directory and name it oop, in the oop directory, create a subdirectory named employee, and create a go file in the directory named employee.go.

Directory structure: oop -> employee -> employee.go

Save the following code in the employee.go file:

```go
package employee

import (  
    "fmt"
)

type Employee struct {  
    FirstName string
    LastName string
    TotalLeaves int
    LeavesTaken int
}

func (e Employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves-e.LeavesTaken))
}
```

Then in the oop directory, create a file and name it main.go, and save the following content

```go
package main

import "oop/employee"

func main() {  
    e := employee.Employee {
        FirstName: "Sam",
        LastName: "Adolf",
        TotalLeaves: 30,
        LeavesTaken: 20,
    }
    e.LeavesRemaining()
}
```

operation result:

```
Sam Adolf has 10 leaves remaining 
```

## 1.2 The New() function replaces the constructor

The program we wrote above looks good, but there is a subtle problem in it. Let's see what happens when we define employee struct with a value of 0. Change the content of main. Go to the code below,

```go
package main

import "oop/employee"

func main() {  
    var e employee.Employee
    e.LeavesRemaining()
}
```

operation result:

```
has 0 leaves remaining
```

It can be known from the running result that the variable created with the zero value of Employee is not available. It does not have a valid first name or last name, nor does it effectively retain details. In other OOP languages, such as java, this problem can be solved by using a constructor. Use the parameterized constructor to create a valid object.



Go does not support constructors. If the zero value of a certain type is not available, the programmer's task is not to export the type to prevent access to other packages, and provide a function named NewT (parameters), which initializes the type T and the required value. In go, it is a convention for naming a function, which creates a value of type T to NewT(parameters). This is like a constructor. If the package defines only one type, then one of its conventions is to name the function New(parameters) instead of NewT(parameters).



Change the code of employee.go:

First modify the employee structure to be non-exported, and create a function New(), which will create a new Employee. code show as below:

```go
package employee

import (  
    "fmt"
)

type employee struct {  
    firstName string
    lastName string
    totalLeaves int
    leavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {  
    e := employee {firstName, lastName, totalLeave, leavesTaken}
    return e
}

func (e employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves-e.leavesTaken))
}
```

We have made some important changes here. We have set the initial letter e of Employee struct to lowercase, that is, we have changed the type Employee struct to type Employee struct. By doing this, we successfully exported the employee structure and prevented access to other packages. It is a good practice to export all the fields of the unexported structure as unexported, unless there is a specific need to export them. Since we don't need to use the employee struct fields anywhere outside the package, we did not export all the fields either.

Since employee is not exported, it is impossible to create a value of type employee from other packages. Therefore, we provide a new output function. Take the required parameters as input and return the newly created employee.

This program still needs to be modified to make it work, but let's run this program to understand the effect of the changes so far. If this program runs, it will fail with the following compilation error,

```
go/src/constructor/main.go:6: undefined: employee.Employee  
```

This is because we have unexported Employee, so the compiler throws an error, the type is not defined in main. perfect. Exactly what we want. No other package can create a zero-valued employee. We successfully prevented an unusable employee structure value from being created. The only way to create employees now is to use the new features.

Modify main.go code

```go
package main  

import "oop/employee"

func main() {  
    e := employee.New("Sam", "Adolf", 30, 20)
    e.LeavesRemaining()
}
```

operation result:

```
Sam Adolf has 10 leaves remaining 
```

Therefore, we can understand that although Go does not support classes, the structure can be used effectively. In the place where the constructor is used, the New(parameters) method can be used.

## 1.3 Composition (Composition) replaces inheritance (Inheritance)

Go does not support inheritance, but it does support composition. The general definition of combination is "put together". An example of composition is a car. A car is made up of wheels, engines, and various other components.

A blog post is an example of a perfect combination. Every blog has title, content and author information. This can be perfectly expressed in combination.

### 1.3.1 Realize composition by embedding structure

This can be achieved by embedding a struct type in another structure.

Sample code:

```go
package main

import (  
    "fmt"
)

/*
We created an author struct, which contains the field name, lastName and bio. We also added a method fullName() with author as the receiver type, which will return the author's full name.
*/
type author struct {  
    firstName string
    lastName string
    bio string
}

func (a author) fullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}
/*
The post struct has field titles and contents. It also has an embedded anonymous field author. This field indicates that the post struct is composed of the author. Now post struct can access all the fields and methods of the author structure. We also added the details() method to the post struct, which prints out the author's title, content, full name and bio.
*/
type post struct {  
    title string
    content string
    author
}

func (p post) details() {  
    fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    fmt.Println("Author: ", p.author.fullName())
    fmt.Println("Bio: ", p.author.bio)
}

func main() {  
    author1 := author{
        "Naveen",
        "Ramanathan",
        "Golang Enthusiast",
    }
    post1 := post{
        "Inheritance in Go",
        "Go supports composition instead of inheritance",
        author1,
    }
    post1.details()
}

```

operation result:

```
Title: Inheritance in Go  
Content: Go supports composition instead of inheritance  
Author: Naveen Ramanathan  
Bio: Golang Enthusiast  
```

**Slices embedded in the structure**

Add the following code under the main function of the above program and run

```go
type website struct {  
        []post
}
func (w website) contents() {  
    fmt.Println("Contents of Website\n")
    for _, v := range w.posts {
        v.details()
        fmt.Println()
    }
}
```

Run error:

```
main.go:31:9: syntax error: unexpected [, expecting field name or embedded type 
```

This error points to the embedded part of structs[]post. The reason is that it is impossible to embed a piece anonymously. A field name is required. Let's fix this error and let the compiler pass.

```go
type website struct {  
        posts []post
}
```

Now let us modify the main function to create several posts for our new website. After modifying the complete code as follows:

```go
package main

import (  
    "fmt"
)

type author struct {  
    firstName string
    lastName string
    bio string
}

func (a author) fullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {  
    title string
    content string
    author
}
func (p post) details() {  
    fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    fmt.Println("Author: ", p.fullName())
    fmt.Println("Bio: ", p.bio)
}

type website struct {  
 posts []post
}
func (w website) contents() {  
    fmt.Println("Contents of Website\n")
    for _, v := range w.posts {
        v.details()
        fmt.Println()
    }
}
func main() {  
    author1 := author{
        "Naveen",
        "Ramanathan",
        "Golang Enthusiast",
    }
    post1 := post{
        "Inheritance in Go",
        "Go supports composition instead of inheritance",
        author1,
    }
    post2 := post{
        "Struct instead of Classes in Go",
        "Go does not support classes but methods can be added to structs",
        author1,
    }
    post3 := post{
        "Concurrency",
        "Go is a concurrent language and not a parallel one",
        author1,
    }
    w := website{
        posts: []post{post1, post2, post3},
    }
    w.contents()
}   
```

operation result:

```
Contents of Website

Title: Inheritance in Go  
Content: Go supports composition instead of inheritance  
Author: Naveen Ramanathan  
Bio: Golang Enthusiast

Title: Struct instead of Classes in Go  
Content: Go does not support classes but methods can be added to structs  
Author: Naveen Ramanathan  
Bio: Golang Enthusiast

Title: Concurrency  
Content: Go is a concurrent language and not a parallel one  
Author: Naveen Ramanathan  
Bio: Golang Enthusiast  
```

## 1.4 Polymorphism

Polymorphism in Go is achieved with the help of interfaces. As we have already discussed, interfaces can be implemented implicitly in Go. If the type provides definitions for all methods declared in the interface, then an interface is implemented. Let's see how to implement polymorphism with the help of interfaces.

Any type that defines all methods of an interface is said to implicitly implement the interface.

Variables of a type interface can hold any value that implements the interface. This attribute of the interface is used to implement polymorphism in Go.

For example, a fictitious organization has revenue from two items: fixed bills and time and materials. The net income of the organization is calculated from the sum of the income of these projects. To keep this tutorial simple, we assume that the currency is US dollars and we will not deal with cents. It will be represented by integers.

First we define an interface: Income

```go
type Income interface {  
    calculate() int
    source() string
}
```

Next, define two structures: FixedBilling and TimeAndMaterial

```go
type FixedBilling struct {  
    projectName string
    biddedAmount int
}
```

```go
type TimeAndMaterial struct {  
    projectName string
    noOfHours int
    hourlyRate int
}
```

The next step is to define the method of these structure types, calculate and return the actual income and source of income.

```go
func (fb FixedBilling) calculate() int {  
    return fb.biddedAmount
}

func (fb FixedBilling) source() string {  
    return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {  
    return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {  
    return tm.projectName
}
```

Next, let's declare the calculateNetIncome function that calculates and prints the total income.

```go
func calculateNetIncome(ic []Income) {  
    var netincome int = 0
    for _, income := range ic {
        fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
        netincome += income.calculate()
    }
    fmt.Printf("Net income of organisation = $%d", netincome)
}
```

The above calculateNetIncome function accepts part of the Income interface as a parameter. It calculates the total income by traversing the slices and calling the calculate() method. It also displays the source of income by calling the source() method. According to the specific type of the income interface, different calculate() and source() methods will be called. Therefore, we have implemented polymorphism in the calculateNetIncome function.

In the future, if the organization adds a new source of revenue, this function can still correctly calculate the total revenue without a line of code change.

Finally we write the following main function:

```go
func main() {  
    project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
    project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
    project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
    incomeStreams := []Income{project1, project2, project3}
    calculateNetIncome(incomeStreams)
}
```

operation result:

```
Income From Project 1 = $5000  
Income From Project 2 = $10000  
Income From Project 3 = $4000  
Net income of organisation = $19000  
```



Suppose the organization has found a new source of income through advertising. Let us see how to simply add a new income method and calculate the total income without making any changes to the calculateNetIncome function. Due to polymorphism, this is feasible.

First, let us define the Advertisement type and calculate() and source() methods.

```go
type Advertisement struct {  
    adName string
    CPC int
    noOfClicks int
}

func (a Advertisement) calculate() int {  
    return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {  
    return a.adName
}
```

The ad type has three fields: adName, CPC (cost per click) and noof clicks (cost per click). The total revenue from advertising is the products of CPC and noOfClicks.



Modify the main function:

```go
func main() {  
    project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
    project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
    project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
    bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
    popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
    incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
    calculateNetIncome(incomeStreams)
}
```

operation result:

```
Income From Project 1 = $5000  
Income From Project 2 = $10000  
Income From Project 3 = $4000  
Income From Banner Ad = $1000  
Income From Popup Ad = $3750  
Net income of organisation = $23750 
```

In summary, we have not made any changes to the calculateNetIncome function, although we have added a new income method. It only works because of polymorphism. Since the new Advertisement type also implements the Income interface, we can add it to the incomeStreams slice. The calculateNetIncome function also works without any changes, because it can call the calculate() and source() methods of the Advertisement type.



Qianfeng Go language learning group: 784190273

Corresponding video address:

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation

