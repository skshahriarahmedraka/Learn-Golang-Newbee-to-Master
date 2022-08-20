

# One, collection (Map)

> @author：Han Ru
> 
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## 1.1 What is Map

map is a built-in type in Go, which associates a value with a key. The value can be retrieved using the corresponding key.

Map is an unordered collection of key-value pairs. The most important point of Map is to quickly retrieve data through key. Key is similar to index and points to the value of data.
Map is a collection, so we can iterate it like arrays and slices. However, Map is unordered, and we cannot determine its return order. This is because Map is implemented using a hash table, which is also a reference type

Points to note when using map: 

-The map is unordered, and the printed map will be different each time. It cannot be obtained by index, but must be obtained by key
-The length of the map is not fixed, that is, like slice, it is also a reference type
-The built-in len function is also applicable to maps, returning the number of keys owned by the map 
-The key of the map can be all comparable types, such as boolean, integer, floating-point, complex, string... and it can also be a key.

## 1.2 Use of Map

### 1.2.1 Use make() to create a map

You can use the built-in function make or the map keyword to define Map:

```go
/* Declare variables, the default map is nil */
var map_variable map[key_data_type]value_data_type

/* Use make function */
map_variable = make(map[key_data_type]value_data_type)
```

```go
rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2}
```

If the map is not initialized, a nil map will be created. nil map cannot be used to store key-value pairs

```go
package main

import "fmt"

func main() {
   var countryCapitalMap map[string]string
   /* Create a collection*/
   countryCapitalMap = make(map[string]string)
   
   /* map insert key-value pairs, the capital of each country*/
   countryCapitalMap["France"] = "Paris"
   countryCapitalMap["Italy"] = "Rome"
   countryCapitalMap["Japan"] = "Tokyo"
   countryCapitalMap["India"] = "New Delhi"
   
   /* Use key to output map value*/
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   /* Check whether the element exists in the collection*/
   captial, ok := countryCapitalMap["United States"]
   /* If ok is true, it exists, otherwise it does not exist*/
   if(ok){
      fmt.Println("Capital of United States is", captial)  
   }else {
      fmt.Println("Capital of United States is not present") 
   }
}
```

operation result:

```go
Capital of France is Paris
Capital of Italy is Rome
Capital of Japan is Tokyo
Capital of India is New Delhi
Capital of United States is not present
```

### 1.2.2 delete() function

The delete(map, key) function is used to delete the elements of the collection, and the parameters are map and its corresponding key. The delete function does not return any value.

```go
package main

import "fmt"

func main() {   
   /* Create map */
   countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
   
   fmt.Println("original map")   
   
   /* Print map */
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   /* Delete element */
   delete(countryCapitalMap,"France");
   fmt.Println("Entry for France is deleted")  
   
   fmt.Println("map after deleting elements")   
   
   /* Print map */
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
}
```

operation result:

```go
Original map
Capital of France is Paris
Capital of Italy is Rome
Capital of Japan is Tokyo
Capital of India is New Delhi
Entry for France is deleted
Map after deleting elements
Capital of Italy is Rome
Capital of Japan is Tokyo
Capital of India is New Delhi
```

### 1.2.3 ok-idiom

We can get the corresponding value in the map by key. The syntax is:

```go
map[key] 
```

But when the key does not exist, we will get the default value of the value type. For example, the string type will get an empty string, and the int type will get 0. But the program will not report an error.

So we can use ok-idiom to get the value and know whether the key/value exists

```go
value, ok := map[key] 
```

Sample code:

```go
package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	fmt.Println(x, ok)
	x, ok = m["a"]
	fmt.Println(x, ok)
}

```

operation result:

```go
0 false
1 true
```

### 1.2.4 The length of the map

Use the len function to determine the length of the map.

```go
len(map) // You can get the length of the map
```

### 1.2.5 map is a reference type

Similar to slices, maps are reference types. When the mapping is assigned to a new variable, they all point to the same internal data structure. Therefore, changes in one will reflect the other.

Sample code:

```go
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("Original person salary", personSalary)
    newPersonSalary := personSalary
    newPersonSalary["mike"] = 18000
    fmt.Println("Person salary changed", personSalary)

}
```

operation result:

```
Original person salary map[steve:12000 jamie:15000 mike:9000]  
Person salary changed map[steve:12000 jamie:15000 mike:18000] 
```

>map cannot be compared with the == operator. == can only be used to check whether the map is empty. Otherwise, an error will be reported: invalid operation: map1 == map2 (map can only be compared to nil)



Qianfeng Go language learning group: 784190273

Author B station:

https://space.bilibili.com/353694001

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation

