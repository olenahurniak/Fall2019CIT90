package main

import "fmt"

func main() {
	xi := []int{4, 3, 8, 9, 19}
	s := foo(xi)
	fmt.Println(s)

	
	t := bar(42, 43, 44)
	fmt.Println(t)

	
	u := bar(xi...)
	fmt.Println(u)
}

func foo(ii []int) int {
	var sum int
	for _, v := range ii {
		fmt.Println("adding to total:", v)
		sum += v
	}
	return sum
}

func bar(ii ...int) int {
	var sum int

	fmt.Println("bar running",ii)
	fmt.Printf("%T\n", ii)
	for _, v := range ii {
		fmt.Println("adding to total:", v)
		sum += v
	}
	return sum
}