package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) String() string {
	return fmt.Sprintf("I am a person name %s\n", p.first)
}

func (p secretAgent) String() string {
	return fmt.Sprintf("I am also a secret agent name %s\n and it is %t that i have a licence to kill", p.first, p.ltk)
}

func main() {
	jm := person {
		first: "jenny",
	}
	fmt.Println(jm)
	a := jm.String()
	fmt.Println(a)

	jb := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	fmt.Println(jb)
	b := jb.person.String()
	c := jb.String()
	fmt.Println(b)
	fmt.Println(c)
}