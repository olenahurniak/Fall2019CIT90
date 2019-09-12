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

	oh := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, oh := range oh {
		fmt.Println(oh.first, " ", oh.last)
		// fmt.Println(oh.last)
		for i, val := range oh.favFlavors {
			fmt.Println(i, "-", val)
		}
		fmt.Println("*********")
	}
}
