package main

import "fmt"

func main() {
	compLit()
	makeWay()
}

func compLit() {
	ii := []int{}
	l := len(ii)
	c := cap(ii)
	fmt.Println(l)
	fmt.Println(c)
	for i := 0; i < 100000; i++ {
		tempC := c
		ii = append(ii, i)
		if tempC != cap(ii) {
			c = cap(ii)
			fmt.Println("new cap", c)
		}
	}
}

func makeWay() {
	ii := make([]int, 0, 11366400)
	l := len(ii)
	c := cap(ii)
	fmt.Println(l)
	fmt.Println(c)
	for i := 0; i < 100000; i++ {
		tempC := c
		ii = append(ii, i)
		if tempC != cap(ii) {
			c = cap(ii)
			fmt.Println("Make new cap", c)
		}
	}
}