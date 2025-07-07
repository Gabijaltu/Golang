package main

import "fmt"

func main() {

	james := []string{"James", "Bond", "Shaken, not stirred"}
	moneypenny := []string{"Miss", "Moneypenny", "I'm 008."}

	inception := [][]string{james, moneypenny}

	for i, outer := range inception {
		fmt.Println("Outer slice index:", i)
		for j, inner := range outer {
			fmt.Printf("Inner slice index: %d, Value: %s\n", j, inner)
		}
	}
}
