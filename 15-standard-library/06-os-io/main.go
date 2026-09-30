package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	message := strings.NewReader("receipt: tea, milk\n")   // An io.Reader.
	if _, err := io.Copy(os.Stdout, message); err != nil { // Stdout is an io.Writer.
		fmt.Fprintln(os.Stderr, "copy failed:", err)
	}
}
