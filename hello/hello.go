package main

import (
	"fmt"

	"example.com/grettings"
)

func main() {
	// Get a gretting message and print it.
	message := grettings.Hello("Eugenia")
	fmt.Println(message)
}
