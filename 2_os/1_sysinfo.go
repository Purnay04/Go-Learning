package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get the system's hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		return
	}
	fmt.Printf("Hostname: %s\n", hostname)

	// Get the value of the PATH environment variable
	envVar := "PATH"
	envValue := os.Getenv(envVar)
	if envValue == "" {
		fmt.Printf("Environment variable %s is not set.\n", envVar)
	} else {
		fmt.Printf("%s entries:\n", envVar)
		paths := strings.Split(envValue, ":")
		for _, p := range paths {
			fmt.Println(p)
		}
	}
}
