package greetings

import "fmt"

// Hello returns a greeting for the named persoon.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message

}
