package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("first")
	fmt.Println("second")
	x := 40 // third
	y := 5
	fmt.Printf("x=%v \n y=%v\n", x, y)

	if 1 == 1 {
		fmt.Println("Hello!")
	}

	if (x < y || x > y) && !(x == y) {
		fmt.Println("X and Y are different.")
	}

	// small variable scope, minimal - best practice.
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}

	switch {
	case x < 42:
		fmt.Println("x is less than 42")
		// assumed break upon execution of 1 case
	case x == 42:
		fmt.Println("x is 42")
	default:
		fmt.Println("default case")
	}
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough // will continue executing through the diff case statements, execute default as well
	case 41:
		fmt.Println("x is 41")
		fallthrough
	default:
		fmt.Println("default case")

	}

}
