package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(250)

	if x < 101 {
		fmt.Printf("Number x is less than 101: %v\n", x)
	} else if x > 100 && x < 201 {
		fmt.Printf("Number x is between 101 and 200: %v\n", x)
	} else {
		fmt.Printf("Number x is greater than 201: %v\n", x)
	}
}
