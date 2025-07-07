package main

import "fmt"

func main() {

	y := 0
	for y < 10 {

		if y == 7 {
			break
		}

		fmt.Printf("y is: %v\n", y)
		y++
	}
}
