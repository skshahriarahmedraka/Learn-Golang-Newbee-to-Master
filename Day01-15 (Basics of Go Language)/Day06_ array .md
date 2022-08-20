

# One, array (Array)

> @author：Han Ru
> 
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.



## 1.1 What is an array

Go language provides data structure of array type.
An array is a sequence of numbered and fixed-length data items with the same unique type. This type can be any primitive type such as integer, string, or custom type.

Array elements can be read (or modified) by index (position), the index starts from 0, the first element is indexed 0, the second index is 1, and so on. The subscript value range of the array is from 0 to the length minus 1.

Once the array is defined, the size cannot be changed.

## 1.2 Array syntax

**Declare and initialize an array**

Need to specify the size of the array and the type of data stored.

```go
var variable_name [SIZE] variable_type
```

Sample code:

```go
var balance [10] float32
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

The number of elements in {} in the initialization array cannot be greater than the number in [].
If you ignore the number in [] and do not set the size of the array, the Go language will set the size of the array according to the number of elements:

```go
var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

```go
balance[4] = 50.0
```

Other ways to create an array:

```go
  var a [4] float32 // Equivalent to: var arr2 = [4]float32{}
  fmt.Println(a) // [0 0 0 0]
  var b = [5] string{"ruby", "王二狗", "rose"}
  fmt.Println(b) // [ruby Wang Ergou rose]
  var c = [5] int{'A','B','C','D','E'} // byte
  fmt.Println(c) // [65 66 67 68 69]
  d := [...] int{1,2,3,4,5}// Set the size of the array according to the number of elements
  fmt.Println(d)//[1 2 3 4 5]
  e := [5] int{4: 100} // [0 0 0 0 100]
  fmt.Println(e)
  f := [...] int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
  fmt.Println(f)
```





**Access array elements**

```go
float32 salary = balance[9]
```

Sample code:

```go
package main

import "fmt"

func main() {
   var n [10]int /* n is an array of length 10*/
   var i,j int

   /* Initialize elements for array n */         
   for i = 0; i <10; i++ {
      n[i] = i + 100 /* Set the element to i + 100 */
   }

   /* Output the value of each array element*/
   for j = 0; j <10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j])
   }
}
```

operation result:

```go
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109
```

**Length of array**

By passing the array as a parameter to the len function, the length of the array can be obtained.

Sample code:

```go
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    fmt.Println("length of a is",len(a))

}
```

operation result:

```
length of a is 4
```

You can even ignore the length of the array in the declaration and replace it with... let the compiler find the length for you. This is done in the program below.

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    a := [...]int{12, 78, 50} // ... makes the compiler determine the length
    fmt.Println(a)
}
```



Traverse the array:

```go
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i <len(a); i++ {//looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}
```



Use range to traverse the array:

```go
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    sum := float64(0)
    for i, v := range a {//range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a",sum)
}
```

If you only need the value and want to ignore the index, you can achieve this by replacing the index with the _blank identifier.

```go
for _, v := range a {//ignores index  
}
```



## 1.3 Multidimensional array

The Go language supports multi-dimensional arrays. The following is the common syntax for multi-dimensional array declarations:

```go
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
```

```go
var threedim [5][10][4]int
```

Three-dimensional array

```go
a = [3][4]int{  
 {0, 1, 2, 3}, /* The index of the first line is 0 */
 {4, 5, 6, 7}, /* The index of the second line is 1 */
 {8, 9, 10, 11} /* The index of the third line is 2 */
}
```















## 1.4 Arrays are value types

Array is value type
Arrays in Go are value types, not reference types. This means that when they are assigned to a new variable, a copy of the original array will be assigned to the new variable. If the new variable is changed, it will not be reflected in the original array.

```go
package main

import "fmt"

func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b) 
}
```

operation result:

```
a is [USA China India Germany France]  
b is [Singapore China India Germany France] 
```

The size of the array is part of the type. Therefore [5]int and [25]int are different types. Therefore, the array cannot be resized. Don't worry about this limitation, because slices exist to solve this problem.

```go
package main

func main() {  
    a := [3]int{5, 78, 8}
    var b [5]int
    b = a //not possible since [3]int and [5]int are distinct types
}
```



Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation

