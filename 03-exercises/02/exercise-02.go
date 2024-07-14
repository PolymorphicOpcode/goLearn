package main

import (
	"fmt"
	"math"
	"math/rand"
)

// when two or more consecutive named function param share type, can omit type from all but last
func add(x, y int) int {
	return x + y
}

func sayHello() {
	fmt.Println("hello")
}

// func name(param) (return var/type) {}
func swap(x, y string) (string, string) {
	return y, x
}

// called a naked return, only to be used in short functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(1000))

	// Capital letter means it's an exported (public) function / constant
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	sayHello()
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
