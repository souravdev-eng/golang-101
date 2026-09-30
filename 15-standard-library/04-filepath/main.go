package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("reports", "2024", "sales.csv") // Use the OS path separator.

	fmt.Println("path:", filepath.ToSlash(path)) // Display the same way on every OS.
	fmt.Println("folder:", filepath.ToSlash(filepath.Dir(path)))
	fmt.Println("file:", filepath.Base(path))
	fmt.Println("extension:", filepath.Ext(path))
}
