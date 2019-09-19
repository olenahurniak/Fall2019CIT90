				// Hands-on exercise #6
				// - Build and use an anonymous func 
package main

import (
	"fmt"
)

func main() {

	func() {
		for i := 10; i <= 25; i++ {
			fmt.Println(i)
		}

	}()

	fmt.Println("done. Note i from 10 to 25")
}
