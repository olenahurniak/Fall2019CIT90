package main

import (
	"fmt"
)

func main() {
	n, m:= foo()

	fmt.Println("The product of two numbers is ", n * m)
	fmt.Println("The sum of two numbers is ", n + m)
}

func foo() (int, int) {
	return 42, 34
}