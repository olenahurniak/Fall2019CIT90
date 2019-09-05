package main

import "fmt"

func main() {
	foo()
	bar()
	monkey()
}

func foo() {
	x := 0

	if x == 4 {
		fmt.Println(4)
	} else if x == 5 {
		fmt.Println(5)
	} else if x == 6 {
		fmt.Println(6)
	} else {
		fmt.Println("x was non of those")
	}
}

func bar() {
	x := 7

	switch x {
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)

	default:
		fmt.Println("non of thoose")
	}

}

func monkey() {
	x := false

	switch {
	case (3 == 2):
		fmt.Println("not true")
	case (3 == 3):
		fmt.Println("3=3 true")
		// fallthrough
	case x:
		fmt.Printf("x is %v \n", x)
		// fallthrough
	default:
		fmt.Println("non of thoose were true")
	}
}
