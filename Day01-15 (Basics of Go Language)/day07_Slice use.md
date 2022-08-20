

# One, slice (Slice)

> @author：Han Ru
> 
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## 1.1 What is a slice

Go slice is an abstraction of array.
The length of the Go array cannot be changed. Such a collection is not suitable in certain scenarios. Go provides a flexible and powerful built-in type slice ("dynamic array"). Compared with the array, the length of the slice is not fixed. Yes, you can append elements, which may increase the capacity of the slice when appending

Slicing is a convenient, flexible and powerful wrapper. The slice itself has no data. They are just references to existing arrays.

Compared with the array, the slice does not need to set the length, and there is no need to set the value in [], which is relatively free

From the conceptual point of view, slice is like a structure, this structure contains three elements: 

1. Pointer to the starting position specified by slice in the array
2. Length, which is the length of the slice
3. The maximum length, which is the length from the beginning of the slice to the last position of the array

## 1.2 The slicing syntax

**Define slice**

```go
var identifier []type
```

The slice does not need to specify the length.
Or use the make() function to create slices:

```go
var slice1 []type = make([]type, len)
Can also be abbreviated as
slice1 := make([]type, len)
```

```go
make([]T, length, capacity)
```

**initialization**

```go
s[0] = 1
s[1] = 2
s[2] = 3
```

```go
s :=[] int {1,2,3} 
```

```go
s := arr[startIndex:endIndex] 
```

Create a new slice from the elements under the subscript startIndex to endIndex-1 in arr (** before closing and opening after **), the length is endIndex-startIndex

```go
s := arr[startIndex:] 
```

The default endIndex will indicate the last element up to arr

```go
s := arr[:endIndex] 
```

The default startIndex will indicate starting from the first element of arr

```go
package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4] //creates a slice from a[1] to a[3]
    fmt.Println(b)
}
```

## 1.3 Modify slice

Slice does not have any data of its own. It is just a representation of the underlying array. Any changes made to the slice will be reflected in the underlying array.

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
    fmt.Println("array after",darr) 
}
```

operation result:

```
array before [57 89 90 82 100 78 67 69 59]  
array after [57 89 91 83 101 78 67 69 59]  
```

When multiple slices share the same underlying array, the changes made by each element will be reflected in the array.

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
    fmt.Println("array after modification to slice nums2", numa)
}
```

operation result:

```
array before change 1 [78 79 80]  
array after modification to slice nums1 [100 79 80]  
array after modification to slice nums2 [100 101 80]  
```



## 1.4 len() and cap() functions

The length of the slice is the number of elements in the slice. The capacity of a slice is the number of elements in the underlying array starting from the index where the slice was created.

The slice is indexable, and the length can be obtained by the len() method
Slicing provides a method to calculate the capacity cap() can measure how long the slice can reach

```go
package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)

   printSlice(numbers)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

operation result

```go
len=3 cap=5 slice=[0 0 0]
```

**Empty Slice**

A slice is nil by default before being initialized, and its length is 0

```go
package main

import "fmt"

func main() {
   var numbers []int

   printSlice(numbers)

   if(numbers == nil){
      fmt.Printf("The slice is empty")
   }
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

operation result

```go
len=0 cap=0 slice=[]
Slice is empty
```

```go
package main

import "fmt"

func main() {
   /* Create slice*/
   numbers := []int{0,1,2,3,4,5,6,7,8}   
   printSlice(numbers)

   /* Print the original slice*/
   fmt.Println("numbers ==", numbers)

   /* Print sub-slices from index 1 (inclusive) to index 4 (not included) */
   fmt.Println("numbers[1:4] ==", numbers[1:4])

   /* The default lower limit is 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])

   /* The default upper limit is len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])

   numbers1 := make([]int,0,5)
   printSlice(numbers1)

   /* Print sub-slices from index 0 (inclusive) to index 2 (not included) */
   number2 := numbers[:2]
   printSlice(number2)

   /* Print sub-slices from index 2 (inclusive) to index 5 (not included) */
   number3 := numbers[2:5]
   printSlice(number3)

}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

operation result

```go
len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
numbers == [0 1 2 3 4 5 6 7 8]
numbers[1:4] == [1 2 3]
numbers[:3] == [0 1 2]
numbers[4:] == [4 5 6 7 8]
len=0 cap=5 slice=[]
len=2 cap=9 slice=[0 1]
len=3 cap=7 slice=[2 3 4]
```

## 1.5 append() and copy() functions 

append appends one or more elements to the slice, and then returns a slice of the same type as the slice
The copy function copy copies elements from the src of the source slice to the target dst, and returns the number of copied elements

The append function will change the contents of the array referenced by the slice, thereby affecting other slices that reference the same array. But when there is nothing left in the slice
When the remaining space (ie (cap-len) == 0), new array space will be dynamically allocated at this time. The returned slice array pointer will point to this space, and the original
The contents of the array will remain unchanged; other slices that reference this array will not be affected

The following code describes the copy method from copying a slice and the append method to append new elements to the slice

```go
package main

import "fmt"

func main() {
   var numbers []int
   printSlice(numbers)

   /* Allow empty slices to be appended*/
   numbers = append(numbers, 0)
   printSlice(numbers)

   /* Add an element to the slice*/
   numbers = append(numbers, 1)
   printSlice(numbers)

   /* Add multiple elements at the same time*/
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)

   /* Create slice numbers1 is twice the capacity of the previous slice*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)

   /* Copy the content of numbers to numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)   
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

operation result

```go
len=0 cap=0 slice=[]
len=1 cap=2 slice=[0]
len=2 cap=2 slice=[0 1]
len=5 cap=8 slice=[0 1 2 3 4]
len=5 cap=12 slice=[0 1 2 3 4]
```

> There is no connection between numbers1 and numbers. When numbers changes, numbers1 will not change. In other words, the copy method will not establish a connection between the two slices



Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation

