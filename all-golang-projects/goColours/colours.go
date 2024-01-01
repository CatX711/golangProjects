package main

import (
	"fmt"
)

const (
	// ANSI color code escape sequences
	red   = "\033[31m"
	green = "\033[32m"
	blue  = "\033[34m"
	reset = "\033[0m"
)

func colorize(color, text string) {
	fmt.Printf("%s%s%s", color, text, reset)
}

func main() {
	// Example usage
	colorize(red, "This is red text\n")
	colorize(green, "This is green text\n")
	colorize(blue, "This is blue text\n")

	// Reset to default color (optional)
	fmt.Println("Back to default color")
}
