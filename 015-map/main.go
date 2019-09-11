package main

import "fmt"

func main() {
	m := map[string]int{
		"james": 22,
		"jenny": 24,
	
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["jenny"])

	v, ok := m["james"]
	fmt.Println(v)
	fmt.Println(ok)

	v, ok = m["jack"]
	fmt.Println(v)
	fmt.Println(ok)

	xi := []int{42, 43, 44, 45}
	for i, v := range xi {
		fmt.Println("slice", i, v)
	}

	for k, v := range m {
		fmt.Println("map", k, v)
	}

	type flavor string
	xf := []flavor{"chocolate", "vanila", "strawbary"}
	for i, v := range xf {
		fmt.Printf("flavor %d %s -", i, v)
		fmt.Printf("type %T - converted %T\n", v, string(v))
	}

	type orders map[int][]flavor
	
	o := orders{
		1: []flavor{"chocolate"},
		2: []flavor{"vanila"},
		3: []flavor{"strawbary"},
		4: []flavor{"peanut butter"},
		5: []flavor{"mint chocolate"},
	}

	for k, v := range o {
		fmt.Printf("order number %d had ", k)
		for _, vv := range v {
			fmt.Printf("%s ", vv)
		}
		fmt.Printf("\n")
	}

	// innit; condition; post
	for i := 0; i < 10; i++ {
		fmt.Printf("map loop number %d - ", i)
		for _, v := range o {
			fmt.Printf("%s ", v)
		}

		fmt.Println()
	}

	xf = []flavor{"chocolate", "strawberry", "vanilla", "peant butter", "mint chocolote"}
	for i := 0; i < 10; i++ {
		fmt.Printf("loop number %d -", i)
		for _, v := range xf {
			fmt.Printf("%s ", v)
		}

		fmt.Println()
	}
}