

# Keyboard input and print output

>Author: Han Ru
>
>Source: Qianfeng Education

## One, print out

### 1.1 fmt package

The fmt package implements formatted I/O similar to C language printf and scanf. The formatting verb ('verb') is derived from the C language but is simpler.

For details, please refer to the API of the official website fmt: https://golang.google.cn/pkg/fmt/

![fmtpackage](img/fmtpackage.png)



### 1.2 Import package

```go
import "fmt"
```



### 1.3 Commonly used printing functions

**Print:**

[func Print(a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Print)

**Formatted printing:**

[func Printf(format string, a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Printf)

**Wrap after printing**

[func Println(a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Println)



Common placeholders in formatted printing:

```
Format and print placeholders:
			%v, output as is
			%T, print type
			%t, bool type
			%s, string
			%f, floating point
			%d, decimal integer
			%b, a binary integer
			%o, octal
			%x, %X, hexadecimal
				%x: 0-9, af
				%X: 0-9, AF
			%c, print characters
			%p, print address
			. . .
```

Sample code:

```go
package main

import (
	"fmt"
)

func main() {
	a := 100 //int
	b := 3.14 //float64
	c := true // bool
	d := "Hello World" //string
	e := `Ruby` //string
	f :='A'
	fmt.Printf("%T,%b\n", a, a)
	fmt.Printf("%T,%f\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e)
	fmt.Printf("%T,%d,%c\n", f, f, f)
	fmt.Println("-----------------------")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

}

```



operation result:

![yunxing1](img/yunxing1.png)



## Two, keyboard input

### 2.1 fmt package read keyboard input

Common methods:

[func Scan(a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Scan)

[func Scanf(format string, a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Scanf)

[func Scanln(a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Scanln)



```go
package main

import (
	"fmt"
)

func main() {
	var x int
	var y float64
	fmt.Println("Please enter an integer, a floating point type:")
	fmt.Scanln(&x,&y)//Read the keyboard input, and assign values ​​to x and y by operating the address. Blocking
	fmt.Printf("The value of x: %d, the value of y: %f\n",x,y)

	fmt.Scanf("%d,%f",&x,&y)
	fmt.Printf("x:%d,y:%f\n",x,y)
}
```

operation result:

![yunxing2](img/yunxing2.png)



### 2.2 bufio package reading

https://golang.google.cn/pkg/bufio/



The bufio package is the method of IO operation:

First create the Reader object:

![bufio1](img/bufio1.png)



Then you can read in various ways:

![bufio2](img/bufio2.png)



Sample code:

```go
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Please enter a string:")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("Data read:", s1)

}
```

running result:

![yunxing3](img/yunxing3.png)



Qianfeng Go language learning group: 784190273

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation

