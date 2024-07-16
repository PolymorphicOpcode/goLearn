package main

import (
	"fmt"
)

var c, python, java bool
var i, j int = 1, 2.0

func main() {
	var i int
	fmt.Printf("%T\t%T\t%T\t%T\t\n", i, c, python, java)
}
