package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	age := m["James"]
	fmt.Println("The age of James", age)

	if v, ok := m["James"]; ok {
		fmt.Println("The age of James", v)
	} else {
		fmt.Println("James is not in the map")
	}

	age = m["Q"]
	fmt.Println("The age of Q", age)

	if v, ok := m["Q"]; ok {
		fmt.Println("The age of Q", v)
	} else {
		fmt.Println("Q is not in the map")
	}
}
