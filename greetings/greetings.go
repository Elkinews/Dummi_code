package grettings // create package with specified name

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string { // function  Name (argument type argument) return type
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
