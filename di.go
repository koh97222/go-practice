package main

import (
	"fmt"
	"io"
)

// Greet 挨拶
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
