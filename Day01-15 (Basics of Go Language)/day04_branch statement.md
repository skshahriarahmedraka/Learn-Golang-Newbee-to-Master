

# One, the process structure of the program

> @author：Han Ru
>
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

There are three types of program flow control structures: sequence structure, selection structure, and loop structure.

Sequence structure: execute from top to bottom, line by line.

Select the structure: certain codes will only be executed when the conditions are met. 0-1 times

​ Branch statement: if, switch, select

Loop structure: When the conditions are met, some code will be executed repeatedly. 0-N times

​ Loop statement: for



# Two, conditional statement

## 2.1 if statement

Syntax format:

```go
if Boolean expression {
   /* Execute when the boolean expression is true */
}
```

```go
if Boolean expression {
   /* Execute when the boolean expression is true */
} else {
  /* Execute when the boolean expression is false */
}
```

```go
if Boolean expression 1 {
   /* Execute when the boolean expression 1 is true */
} else if Boolean expression 2{
   /* Execute when Boolean expression 1 is false and Boolean expression 2 is true */
} else{
   /* When the above two Boolean expressions are both false, execute */
}
```



Sample code:

```go
package main

import "fmt"

func main() {
   /* Define local variables*/
   var a int = 10
 
   /* Use if statement to judge Boolean expression*/
   if a <20 {
       /* If the condition is true, execute the following statement*/
       fmt.Printf("a is less than 20\n")
   }
   fmt.Printf("The value of a is: %d\n", a)
}
```

## 2.2 if variant



If it contains an optional statement component (executed before the condition is evaluated), there is also a variant. Its syntax is

```go
if statement; condition {  
}

if condition{
    
    
}
```

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    if num := 10; num% 2 == 0 {//checks if number is even
        fmt.Println(num,"is even") 
    } else {
        fmt.Println(num,"is odd")
    }
}
```

>It should be noted that the definition of num is in if, so it can only be used in the if..else statement block, otherwise the compiler will report an error.





## 2.3 switch statement: "switch"

A switch is a conditional statement that evaluates an expression and compares it with a list of possible matches, and executes the code block based on the match. It can be considered an idiomatic way to write multiple if else clauses.

The switch statement is used to perform different actions based on different conditions. Each case branch is unique, and it is tested one by one from top to bottom until it matches.
The switch statement executes from top to bottom until a match is found. There is no need to add break after the match.

And if switch has no expression, it will match true

By default, switch in Go is equivalent to having a break at the end of each case. After a successful match, it will not automatically execute other cases, but jump out of the entire switch, but fallthrough can be used to force the subsequent case code to be executed.

The variable var1 can be of any type, while val1 and val2 can be any values ​​of the same type. The types are not limited to constants or integers, but must be the same type; or the final result is an expression of the same type.
You can **test multiple possible values ​​at the same time, and separate them with commas**, for example: case val1, val2, val3.

```go
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```


Sample code:

```go
package main

import "fmt"

func main() {
   /* Define local variables*/
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70: grade = "C" //case can have multiple values
      default: grade = "D"  
   }

   switch {
      case grade == "A":
         fmt.Printf("Excellent!\n")     
      case grade == "B", grade == "C":
         fmt.Printf("Good\n")      
      case grade == "D":
         fmt.Printf("Pass\n")      
      case grade == "F":
         fmt.Printf("Failed\n")
      default:
         fmt.Printf("Different\n" );
   }
   fmt.Printf("Your grade is %s\n", grade );      
}
```



## 2.4 fallthrough

If you need to go through the subsequent case, add fallthrough

```go
package main

import (
	"fmt"
)

type data [2]int

func main() {
	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)
		fallthrough
	case 6:
		x += 20
		fmt.Println(x)

	}

}

```

operation result:

```go
15
35
```



The expression in case is optional and can be omitted. If the expression is omitted, it is considered switch true, and each case expression is evaluated as true, and the corresponding code block is executed.

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    num := 75
    switch {// expression is omitted
    case num >= 0 && num <= 50:
        fmt.Println("num is greater than 0 and less than 50")
    case num >= 51 && num <= 100:
        fmt.Println("num is greater than 51 and less than 100")
    case num >= 101:
        fmt.Println("num is greater than 100")
    }

}
```



> Precautions for switch
>
> 1. The constant value after the case cannot be repeated
> 2. There can be multiple constant values ​​after the case
> 3. fallthrough should be the last line of a case. If it appears somewhere in the middle, the compiler will throw an error.

## 2.5 Type Switch

The switch statement can also be used in type-switch to determine the type of variable actually stored in an interface variable.

```go
switch x.(type){
    case type:
       statement(s);      
    case type:
       statement(s); 
    /* You can define any number of cases */
    default: /* optional*/
       statement(s);
}
```

```go
package main

import "fmt"

func main() {
   var x interface{}
     
   switch i := x.(type) {
      case nil:	  
         fmt.Printf("Type of x:%T",i)                
      case int:	  
         fmt.Printf("x is int type")                       
      case float64:
         fmt.Printf("x is float64 type")           
      case func(int) float64:
         fmt.Printf("x is of type func(int)")                      
      case bool, string:
         fmt.Printf("x is bool or string type")       
      default:
         fmt.Printf("Unknown Type")     
   }   
}
```

operation result:

```go
Type of x:
   
    
```



Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation



   

