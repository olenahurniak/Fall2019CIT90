package main

import "fmt"

func main() {
	x := 42
	fmt.Println("x\t", x)
	fmt.Println("&x\t", &x)
	fmt.Printf("type of x\t %T\n", x)
	fmt.Printf("type of &x\t %T\n", &x)
	
}
