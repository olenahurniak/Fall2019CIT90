package main

import "fmt"

type grade int
type student struct {
	first, last string
	score       grade
}
type result map[string]int

func main() {
	mike := student{
		first: "Mike",
		last:  "Filipo",
		score: 42,
	}
	fmt.Println(mike)

	r := result{
		"taqueria 2": 5,
		
	}
}
