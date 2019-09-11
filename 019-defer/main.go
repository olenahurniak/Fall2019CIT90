package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("1 foo")
}

func bar() {
	fmt.Println("2 bar")
}

// examples
// https://godoc.org/os#pkg-index
// https://godoc.org/os#File.Read
