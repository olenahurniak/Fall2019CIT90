// Hands-on exercise #3
package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
	fuel string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type electro struct {
	vehicle
	fourWheel bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
			fuel: "gas",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
			fuel: "gas",
		},
		luxury: false,
	}

	e := electro{
		vehicle: vehicle{
			doors: 4,
			color: "red",
			fuel: "electro",
		},
		fourWheel: true,
		
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(e)
	fmt.Println(e.fuel, t.fuel, s.fuel)
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}
