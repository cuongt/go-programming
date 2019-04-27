package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int
	for _, a := range os.Args[1:] {

		fmt.Print(os.Args[1:])
		fmt.Println("\n")
		i, err := strconv.Atoi(a)
		fmt.Print(i)
		fmt.Println("\n")
		if err != nil {
			panic(fmt.Sprintf("Invalid value: %v", err))
		}
		sum += i
	}
	fmt.Printf("Sum = %v \n", sum)
}
