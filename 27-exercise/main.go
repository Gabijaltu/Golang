package main

import "fmt"

func main() {

	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for index, value := range a {
		fmt.Printf("Type: %T, Value: %v, Index: %v\n", value, value, index)
	}

}
