package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, World.")

	/* Building using `go build -ldflags "-w -s"` will remove symbols and debug info*/
}
