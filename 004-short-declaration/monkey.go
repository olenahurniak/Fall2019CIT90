package main

import (
	"fmt"
)

var b = "pkg"
// var b string = "pkg"

func main() {

	// short declaration operator
	// strings are wrapped wit "" or ''
	
	x := "Lily"
	y := 30
	z := 42
	// a := 'here is'

	// shift + alt + <arrow key>
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// fmt.Println(a)
	fmt.Println(b)


	fmt.Printf("the VALUE stored in x is %v and is of Type %T\n", x, x)
	fmt.Printf("the VALUE stored in y is %v and is of Type %T\n", y, y)
	fmt.Printf("the VALUE stored in z is %v and is of Type %T\n", z, z)
	// fmt.Printf("the VALUE stored in x is %v and is of Type %T\n", a, a)
	fmt.Printf("the VALUE stored in b is %v and is of Type %T\n", b, b)

	fmt.Println("calling foo")
	foo()

	// go run *.go
}