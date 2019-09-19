package main

import "fmt"

type person struct {
	first, last string
}
type dbPG map[int]person

// var mPG dbPG
type dbMongo map[int]person

func (pg dbPG) access(x int) person {
	return pg[x]
}

func (pg dbPG) set(x int, p person) {
	pg[x] = p
}

func init() {

}
func main() {
	p1 := person{
		first: "lily",
		last:  "frangyan",
	}
	mpg := dbPG{}
	mpg.set(1, p1)
	fmt.Println(mpg)	/// map[1:{lily frangyan}]
	fmt.Println(mpg.access(1))	// {lily frangyan}
}
