package main

import "fmt"

type person struct {
	first  string
	last   string
	favnum int
	drinks bool
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first:  "john",
		last:   "doe",
		favnum: 12,
		drinks: false,
	}
	fmt.Println(p1)

	sal := secretAgent{
		person: person{
			first:  "lily",
			last:   "jun",
			favnum: 10,
			drinks: false,
		},
		ltk: true,
	}
	fmt.Println(sal)
}
