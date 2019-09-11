package main

import "fmt"

func main() {
	// composite literal
	oh := []string{"Antoine de Saint-Exupéry,", "The Little Prince"}
	fmt.Println(oh)
	fmt.Println(oh[0])
	fmt.Println(oh[1])

	for i, v := range oh {
		fmt.Println(i, "-", v)	
	}
}