package main

import "fmt"

func main() {

	a := [5]int{}

	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	for index, value := range a {
		fmt.Printf("Type: %T, Value: %v, Index: %v\n", value, value, index)
	}

}
