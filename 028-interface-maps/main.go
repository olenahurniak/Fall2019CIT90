package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Printf("I am a person name %s\n", p.first)
}

func (p secretAgent) speak()  {
	fmt.Printf("I am also a secret agent name %s\n and it is %t that i have a licence to kill", p.first, p.ltk)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	jm := person {
		first: "jenny",
	}
	saySomething(jm)

	jb := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	saySomething(jb)
}