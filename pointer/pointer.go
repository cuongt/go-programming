package main

import "fmt"

func main() {
	// s := "new string"
	// sptr := &s
	// fmt.Println(s)
	// fmt.Println(*sptr)

	// := is a declaration, whereas = is an assignment

	s := "string"
	var sptr *string
	sptr = &s
	fmt.Println(*sptr)

	i := new(int)
	fmt.Println(*i)
}
