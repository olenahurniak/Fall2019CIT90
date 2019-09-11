package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "martini", "shaken, not stirred"}
	jm := []string{"jenny", "moneypay", "manhattan", "go for it"}
	xs :=[][]string{jb, jm}

	fmt.Println(xs)	// [[james bond martini shaken, not stirred] [jenny moneypay manhattan go for it]]
	fmt.Println(xs[0])	// [james bond martini shaken, not stirred]
	fmt.Println(xs[0][0])	// james
	fmt.Println(xs[0][1])	// bond
	fmt.Println(xs[0][2])	// martini
	fmt.Println(xs[0][3])	// shaken, not stirred
	fmt.Println(xs[1])		// [jenny moneypay manhattan go for it]
	fmt.Println(xs[1][0])
	fmt.Println(xs[1][1])
	fmt.Println(xs[1][2])
	fmt.Println(xs[1][3])

}