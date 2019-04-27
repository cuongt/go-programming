package main

import (
	"fmt"
	"strings"
)

// main method
func main() {

	// boolean
	a := true
	b := false

	fmt.Printf("a = %v, b = %v, a == b = %v\n", a, b, a == b)

	// Number
	i := 100
	j := 101

	fmt.Printf("%v + %v = %v\n", i, j, i+j)

	// String

	s := "This is my string"
	fmt.Printf("End with string? %v\n", strings.HasSuffix(s, "string"))

}
