package main

import (
	"io"
	"os"
)

func main() {
	file, _ := os.Create("Hello.txt")
	_, _ = io.WriteString(file, "Hello World")
	file.Close()
}
