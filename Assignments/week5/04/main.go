// Hands-on exercise #9
// A “callback” is when we pass a func into a func as an argument. For this exercise, 
// pass a func into a func as an argument 

package main

import (
        "fmt"
)

type funcDef func(string) string

func foo(s string) string {
        return fmt.Sprintf("from foo: %s", s)
}

func test(someFunc funcDef, s string) string {
        return someFunc(s)
}

func main() {
        output := test(foo, "some string")
        fmt.Println(output)
}