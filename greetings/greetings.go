package greetings

import "fmt"

//Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return  a greeting that embeds the name person.
	message := fmt.Sprintln("Hi, %v. Welcome", name)
	return message
}
