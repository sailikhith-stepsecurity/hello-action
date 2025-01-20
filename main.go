package main

import (
	"fmt"
	"os"
)

func main() {
	// Get input from environment variable
	input := os.Getenv("INPUT_MESSAGE")
	if input == "" {
		input = "Hello, World!"
	}

	// Print the input
	fmt.Printf("Message Updated from v1.0.2: %s\n", input)
}
