package main

import "fmt"

var x = 40

const y = 41

func main() {
	z := 42
	fmt.Printf("The value of z is %v and its type is %T\n", z, z)
	fmt.Printf("The value of x is %v and its type is %T\n", x, x)
	fmt.Printf("The value of y is %v and its type is %T\n", y, y)
}
