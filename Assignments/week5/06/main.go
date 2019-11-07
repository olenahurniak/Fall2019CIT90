// HANDS-on EXERCISE #6 week 5
package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("type of x\t %T\n", x)
	fmt.Printf("type of &x\t %T\n ", &x)
	foo(&x)
}

func foo(y *int) {
	*y++
	fmt.Println(y)
	fmt.Println(*y)
	
	fmt.Println(&y)
	fmt.Printf("type of y\t %T\n", y)
	fmt.Printf("type of &y\t %T\n ", &y)
}