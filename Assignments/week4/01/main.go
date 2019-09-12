package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "Pavlo",
		last:  "Gurnyak",
		favFlavors: []string{
			"Chocolate",
			"Pistachios",
			"Caramel",
		},
	}

	p2 := person{
		first: "Olena",
		last:  "Hurniak",
		favFlavors: []string{
			"Strawberry",
			"Tiramisu",
			"Lemon",
		},
	}

	fmt.Println(p1.first, " ", p1.last)
	// fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, "-", v)
	}

	fmt.Println(p2.first, " ", p2.last)
	// fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, "-", v)
	}
}
