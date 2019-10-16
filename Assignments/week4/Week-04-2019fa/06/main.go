// Hands-on exercise #4
package main
import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
	occupation string
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old. I am ", p.occupation)
}

func main() {
	p1 := person{
		first: "Lily",
		last:  "Frangian",
		age:   12,
		occupation: "Student",
	}

	p1.speak()

	p2 := person{
		first: "Armine",
		last:  "Frangian",
		age:   33,
		occupation: "Mother",
	}
	p2.speak()
}
