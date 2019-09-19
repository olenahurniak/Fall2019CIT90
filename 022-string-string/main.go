package main

import "fmt"

type hotdog int

var nyStand hotdog

func (h hotdog) String() {
	return fmt.Sprintf("There are %d stands in NY\n", h)
}

func main() {
	nyStand = 4
	fmt.Println(nyStand)
	z := nyStand.String()
	fmt.Println("the assigned z:", z)
}
