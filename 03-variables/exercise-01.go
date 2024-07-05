package main

import "fmt"

func main() {

	var h int = 44
	const i string = "hello"
	fmt.Printf("%s, %d\n", i, h)

	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	j := 42
	fmt.Printf("42 as binary, %b \n", j)
	fmt.Printf("42 as hexadecimal, %x \n", j)

	g, h := 0, 1
	fmt.Printf("%v \t %b \t %#x\n", g, g, g)
	fmt.Printf("%v \t %b \t %#x\n", h, h, h)

}
