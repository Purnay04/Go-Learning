package main

import (
	"fmt"
	"os"
)

func main() {
	// --- Writing to a file ---
	writeData := "Hello, Go!"
	err := os.WriteFile("/Users/administrator/Desktop/output.txt", []byte(writeData), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Println("Data written to output.txt.")

	// --- Reading from a file ---
	data, err := os.ReadFile("/Users/administrator/Desktop/output.txt") // You can use any .txt file here
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println("File content:")
	fmt.Println(string(data))
}
