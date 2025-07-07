package main

import (
	"fmt"
	"math/rand"
)

func main() {

	count := 0

	for i := 0; i < 100; i++ {

		if x := rand.Intn(5); x == 3 {
			count++
			fmt.Printf("iteration %v\t total count %v\t x is three\n", i, count)
		}
	}

}
