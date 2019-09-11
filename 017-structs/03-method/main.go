package main

import "fmt"

type person struct {
	first  string
	last   string
	favnum int
	drinks bool
}

// func (r Receiver )
func (p person) speak() {
	fmt.Printf("I am a person an my name is %s %s \n", p.first, p.last)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Printf("I am secret agent and my name is %s %s \n", sa.person.first, sa.person.last)
}

func main() {
	p1 := person{
		first:  "john",
		last:   "doe",
		favnum: 12,
		drinks: false,
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person: person{
			first:  "lily",
			last:   "jun",
			favnum: 10,
			drinks: false,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa1.person.speak()
}
