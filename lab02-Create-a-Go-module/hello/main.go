package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("Greetings")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"HP", "Await", "Darrin", "VanHieu", "vandao"}

	// Request a greeting message.
	message, err := greetings.HelloMultiplePeople(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
