package main

import (
	"fmt"
)

func main() {
	oh := []string{"Lily,", "Armine,", "Hovo,", "David,", "Gagik,"}
	for i, v := range oh {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", oh)
}