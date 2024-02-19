package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

// Accepts string default "Hello";
// Returns string;
// Greets the provided string, defaulting to standard "Hello, World" when not provided
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
