// Hands-on exercise #8
// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

package main

import "fmt"

func main() {
	oh := foo()
	pg := bar()
	fmt.Println(oh()) /// 34
	fmt.Println(pg()) /// My name is...
}
// int
func foo() func() int {
	return func() int {
		return 32+2
	}
}
// string
func bar() func() string {
	return func() string {
		return "My name is..."

	}
}