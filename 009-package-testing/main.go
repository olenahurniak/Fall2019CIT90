package main

import (
	"fmt"

	"github.com/olenahurniak/Fall2019CIT90/009-package-testing/helpers"
)

func main() {
	x := foo()
	fmt.Println(x)

	y := helpers.Bar()
	fmt.Println(y)
	fmt.Println("exit")
}
