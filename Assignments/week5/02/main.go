// Hands-on exercise #7
// Assign a func to a variable, then call that func

package main


import (
	"fmt"
)

var x int
var g func()

func main() {

	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i) // 0 1 2 
		}
	}
	
	f()
	fmt.Printf("%T\n", f) // func()

	fmt.Println(x) // 0
	fmt.Printf("%T\n", x) // int

	g = f
	g()	// 0 1 2 
	fmt.Printf("this is g %T\n", g) // this is g func()
	
}