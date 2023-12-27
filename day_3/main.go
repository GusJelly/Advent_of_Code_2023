package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := loadFile("example")

	/* matrix := getMatrix(file)

	symbols := checkSymbols(matrix)
	for _, el := range symbols {
		fmt.Println(el)
	} */

	numbers := getNumbers(file)
	fmt.Println(numbers)
	calcSum(file, numbers)
}

func loadFile(name string) []string {
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
