package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		switch {
		case x < 4 && y < 4:
			fmt.Printf("Both x and y are less than 4. X is: %v, Y is: %v\n", x, y)
		case x > 6 && y > 6:
			fmt.Printf("Both x and y are greater than 6. X is: %v, Y is: %v\n", x, y)
		case x >= 4 && x <= 6:
			fmt.Printf("X is between 4 and 6. X is: %v\n", x)
		case y != 5:
			fmt.Printf("Y is not equal to 5. Y is: %v\n", y)
		default:
			fmt.Printf("No conditions met. X is: %v, Y is: %v\n", x, y)
		}
	}

}
