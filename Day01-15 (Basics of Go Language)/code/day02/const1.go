package main

import "fmt"

func main() {
	
	// Print a Constant Value
	fmt.Println(100)
	fmt.Println("hello")

	// Declear a Constant string and int
	const PATH string = "http://walcart.com"
	const PI = 3.14
	fmt.Println(PATH)
	
	const C1, C2, C3 = 100, 3.14, "haha"
	const (
		MALE   = 0
		FEMALE = 1
		UNKNOW = 3
	)
	const (
		a int = 100
		b
		c string = "ruby"
		d
		e
	)
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%d\n", b, b)
	fmt.Printf("%T,%s\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e)

	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)

}
