package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olenahurniak/Fall2019CIT90/Assignments/week8/convert"
)

func main() {
	for _, arg := range os.Args[1:] {
		w, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		p := weightconv.Pound(w)
		k := weightconv.Kilogram(w)
		fmt.Printf("%s = %s, %s = %s\n",
			p, weightconv.PToKg(p), k, weightconv.KgToP(k))
	}
}
