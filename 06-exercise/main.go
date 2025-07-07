package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Puppy says:", puppy.Bark())
	fmt.Println(puppy.BigBark())
	puppy.From11()
	puppy.From12()
}
