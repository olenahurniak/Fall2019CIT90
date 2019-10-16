// Hands-on exercise #2
package main

import (
	"fmt"
)

func main() {
	ii := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	n := foo(ii...)
	fmt.Println(n)
	
	ii2 := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	n2 := bar(ii2)
	fmt.Println(n2)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
