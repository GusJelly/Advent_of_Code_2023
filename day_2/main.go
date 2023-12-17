package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readFile("input")
	for _, v := range lines {
		fmt.Println(v)
	}
}

// Loads a given file into a string slice
func readFile(name string) []string {
	// Load the file from the system
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create scanner to read the file
	var lines []string
	scanner := bufio.NewScanner(file)

	// Put all the lines inside the file into a slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
