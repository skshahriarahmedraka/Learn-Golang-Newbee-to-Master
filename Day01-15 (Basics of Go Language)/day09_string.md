

# One, string (string)

> @author：Han Ru
> 
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## 1.1 What is a string

A string in Go is a slice of bytes. You can create a string by encapsulating its content in "". Strings in Go are Unicode compatible and UTF-8 encoded.

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    name := "Hello World"
    fmt.Println(name)
}
```



## 1.2 Use of string

### 1.2.1 Access a single byte in a string

```go
package main

import (  
    "fmt"
)

func main() {  
    name := "Hello World"
    for i:= 0; i <len(s); i++ {
        fmt.Printf("%d ", s[i])
    }
    fmt.Printf("\n")
    for i:= 0; i <len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}
```

operation result:

72 101 108 108 111 32 87 111 114 108 100 
H ello W orld 

## 1.3 strings package

There are many functions for manipulating strings when accessing the strings package.



## 1.4 strconv package

Visit the strconv package, you can convert between string and other numeric types.



Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation


