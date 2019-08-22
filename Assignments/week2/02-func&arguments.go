package main

import (
	"fmt"
)

func main() {
	n, m := foo()
	x, z := bar()
	c, d := far()

	fmt.Println(n, m)
	fmt.Println(x, z)
	fmt.Println(c, d)
	fmt.Println(n, m, x, z, c, d)
}

func foo() (int, string) {
	return 1447, "Herndon Ave"
}

func bar() (int, string) {
	return 32, "Elen"
}

func far() (int, string) {
	return 37, "Juliia"
}