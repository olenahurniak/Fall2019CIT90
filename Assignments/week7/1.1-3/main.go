package main

import (
	"fmt"
	"os"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " \n"
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
