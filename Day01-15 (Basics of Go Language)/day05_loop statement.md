

# One, the process structure of the program

> @author：Han Ru
>
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

There are three types of program flow control structures: sequence structure, selection structure, and loop structure.

Sequence structure: code execution line by line from top to bottom

Select the structure: certain codes will be executed only if the conditions are met. 0-1 times

​ if statement, switch statement

Loop structure: When the conditions are met, some code will be executed repeatedly. 0-n times

​ for statement

# Two, loop statement

The loop statement indicates that the condition is met, and a certain piece of code can be executed repeatedly.

for is the only loop statement. (Go does not have a while loop)

##2.1 for statement

Grammatical structures:

```
for init; condition; post {}
```

> The initialization statement is executed only once. After the initialization loop, the condition will be checked. If the condition is evaluated as true, then the body of the loop in {} will be executed, followed by the post statement. The post statement will be executed after each successful iteration of the loop. After executing the post statement, the condition will be rechecked. If it is correct, the loop will continue to execute, otherwise the loop will terminate.

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        fmt.Printf(" %d",i)
    }
}
```

>Variables declared in a for loop are only available within the scope of the loop. Therefore, i cannot access the loop outside.



## 2.2 for loop variants

**All three components, namely initialization, conditions and post are optional. **

```
for condition {}
```

The effect is similar to while

```
for {}
```

The effect is the same as for(;;)

The range format of the for loop can iteratively loop over slices, maps, arrays, strings, etc.

```
for key, value := range oldMap {
    newMap[key] = value
}
```

```go
package main

import "fmt"

func main() {

   var b int = 15
   var a int

   numbers := [6]int{1, 2, 3, 5} 

   /* for loop*/
   for a := 0; a <10; a++ {
      fmt.Printf("The value of a is: %d\n", a)
   }

   for a <b {
      a++
      fmt.Printf("The value of a is: %d\n", a)
      }

   for i,x:= range numbers {
      fmt.Printf("The value of position %d x = %d\n", i,x)
   }   
}
```

operation result:

```
The value of a: 0
The value of a: 1
The value of a: 2
The value of a: 3
The value of a: 4
The value of a: 5
The value of a: 6
The value of a: 7
The value of a: 8
The value of a: 9
The value of a: 1
The value of a: 2
The value of a: 3
The value of a: 4
The value of a: 5
The value of a: 6
The value of a: 7
The value of a: 8
The value of a: 9
The value of a: 10
The value of a: 11
The value of a: 12
The value of a: 13
The value of a: 14
The value of a: 15
The value of bit 0 x = 1
The value of the 1st bit x = 2
The value of x at bit 2 = 3
The value of x at bit 3 = 5
The value of the 4th bit x = 0
The value of the 5th bit x = 0
```



## 2.2 Multi-layer for loop

There are loop nesting in the for loop, which means multiple loops.



# Three, the statement that jumps out of the loop

> @author：Han Ru
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## 1, break statement

break: Jump out of the loop body. The break statement is used to abruptly terminate the for loop before ending its normal execution

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i> 5 {
            break //loop is terminated if i> 5
        }
        fmt.Printf("%d ", i)
    }
    fmt.Printf("\nline after for loop")
}
```

## 2, continue statement

continue: Jump out of a loop. The continue statement is used to skip the current iteration of the for loop. All the code in the for loop after the continue statement will not be executed in the current iteration. The loop will continue to the next iteration.

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Printf("%d ", i)
    }
}
```





# Four, goto statement

goto: You can unconditionally move to the specified line in the process.

Grammatical structures:

```go
goto label;
..
..
label: statement;
```



![goto1](img/goto1.jpg)

```go
package main

import "fmt"

func main() {
   /* Define local variables*/
   var a int = 10

   /* Loop*/
   LOOP: for a <20 {
      if a == 15 {
         /* Skip iteration*/
         a = a + 1
         goto LOOP
      }
      fmt.Printf("The value of a: %d\n", a)
      a++     
   }  
}
```



Unified error handling
Multiple error handling is very tricky when there is code duplication, for example:

```go
		err := firstCheckError()
    if err != nil {
        goto onExit
    }
    err = secondCheckError()
    if err != nil {
        goto onExit
    }
    fmt.Println("done")
    return
onExit:
    fmt.Println(err)
    exitProcess()



```



The picture in this article comes from the Internet





Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation



