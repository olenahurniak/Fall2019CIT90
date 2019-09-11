package main

import ("fmt")

func main() {
	n:= foo()

	fmt.Println("Hello, ", n, "!!!")
}

func foo() (string) {
	return "this is my class CIT90 - Go programming"
}