package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	fmt.Printf("%v is %d years old. \t and the type is %T and %T", name, age, name, age)
	fmt.Println("👹")
	fmt.Println(`
	
		You're so skibidi
				
	`)
}
