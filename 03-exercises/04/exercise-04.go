package main

import "fmt"

func main() {
	// Zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	t := string(i)
	// J inherits "int" type from i
	j := i
	fmt.Printf("%v\t%v\n", t, j)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)

}
