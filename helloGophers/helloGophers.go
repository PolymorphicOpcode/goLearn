package main

import (
	"fmt"
	"os"
)

// this is a comment
/*
To Learn: Python Poetry, Go, Cobra
https://www.udemy.com/course/learn-how-to-code/?couponCode=ST16MT70224
*/

func main() {
	var name string
	var age int

	fmt.Print("Name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Age: ")
	fmt.Scanf("%d", &age)
	_, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
}
