package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}
