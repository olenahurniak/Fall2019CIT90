package main

import "fmt"

type person struct {
	first, last string
}

type dbPG map[int]person

type dbMongo map[int]person

type storage interface {
	access(int) person
	set(int, person)
}

func gettr(s storage, x int) {
	s.access(x)
}

func setter(s storage, x int, p person) {
	s.set(x, p)
}

// func (pg dbPG) access(x int) person {
// 	return pg[x]
// }

// func (pg dbPG) set(x int, p person) {
// 	pg[x] = p
// }

// func (pg dbMongo) access(x int) person {
// 	fmt.Println("")
// 	return pg[x]
// }

// func (pg dbMongo) set(x int, p person) {
// 	fmt.Println("")
// 	pg[x] = p
// }

// func init() {

// }
func main() {
	p1 := person{
		first: "lily",
		last:  "frangyan",
	}
	mpg := dbPG{}
	setter(mpg, 1, p1)
	fmt.Println(mpg)
}
