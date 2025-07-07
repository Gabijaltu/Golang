package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is the init function. It runs before main.")
}

func main() {

	x := rand.Intn(250)

	switch {
	case x < 101:
		fmt.Printf("Number x is less than 101: %v\n", x)
	case x > 100 && x < 201:
		fmt.Printf("Number x is between 101 and 200: %v\n", x)
	default:
		fmt.Printf("Number x is greater than 201: %v\n", x)
	}
}
