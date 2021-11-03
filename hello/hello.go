package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//Set properties of predefined logger, including
	//the log entry prefix and a flag to disable printing
	//the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Georgina", "el Bicho", "Juanchi"}

	//Request a greeting message
	message, err := greetings.Hellos(names)
	//if an error was returned, print it to console and
	//exit the program.
	if err != nil {
		log.Fatal(err)
	}
	//If no errors was returned, print it to the console and
	//exit the program
	fmt.Println(message)
}
