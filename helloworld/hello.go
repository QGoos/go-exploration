package helloworld

import (
	"fmt"
)

const (
	german = "German"
	french = "French"

	englishHelloPrefix = "Hello, "
	germanHelloPrefix  = "Hallo, "
	frenchHelloPrefix  = "Bonjour, "
)

// Accepts: string, string default "Hello", default "English"
// Returns: string;
// Greets the provided string, defaulting to standard "Hello, World" when not provided
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Accepts: language string; default "English"
// Returns: prefix string
func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func helloworld() {
	fmt.Println(Hello("world", ""))
}
