package main

import "fmt"

func main() {

	a := "Hello World!!!" // string
	b := 2019           // int
	c := true          // bool


	x := []int{10, 20, 30, 40, 50}
	fmt.Println(a, b, c)
	fmt.Println(x)
	fmt.Printf("the type of x is %T\n", x)
	
	for i, v := range x {
		fmt.Println("index", i, "we have value -", v)
	}
}