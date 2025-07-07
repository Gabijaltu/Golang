package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(10)
	y := rand.Intn(10)

	switch {
	case x < 4 && y < 4:
		fmt.Printf("Both x and y are less than 4. X is: %v, Y is: %v", x, y)
	case x > 6 && y > 6:
		fmt.Printf("Both x and y are greater than 6. X is: %v, Y is: %v", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("X is between 4 and 6. X is: %v", x)
	case y != 5:
		fmt.Printf("Y is not equal to 5. Y is: %v", y)
	default:
		fmt.Printf("No conditions met. X is: %v, Y is: %v", x, y)
	}
}
