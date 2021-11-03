package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name wa given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	//if a name was received, return a value that embeds the name
	// in a greeting message
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat())  // purpose produce a error test
	return message, nil
}

//Hellos returns a map that associates each of the named people
//with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)
	// Loop through the recieved slice names, calling
	// the hello function to get a message fot each name
	for _, name := range names {
		// call previus function to assing a message depends on key maps (name)
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name
		messages[name] = message
	}
	return messages, nil
}

//init sets initial values for variables used in this function
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomformat returns one of a set of greeting messages. The returned
//message is selected at random.
func randomFormat() string {
	// a slice of message formats.
	formats := []string{
		"Hi uncle, %v.Welcome!",
		"Great to see you ,%v!",
		"Hail, %v well met!",
	}

	//Return a randomly selected message format by specifying
	//a random index for a slice of formats
	return formats[rand.Intn(len(formats))]
}
