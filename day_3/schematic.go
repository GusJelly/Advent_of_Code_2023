package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var directions map[string][]int = map[string][]int{
	"up":    {-1, 0},
	"down":  {1, 0},
	"left":  {0, -1},
	"right": {0, 1},

	"upLeft":    {-1, -1},
	"downLeft":  {1, -1},
	"upRight":   {-1, -1},
	"downRight": {1, 1},
}

func getIndexes(file []string, numbers []string) ([][]int, [][]int) {
	var numbersIndexes [][]int
	var symbolsIndexes [][]int

	// Get the indexes of all numbers in the file
	for _, line := range file {
		for _, num := range numbers {
			r, _ := regexp.Compile(num)
			index := r.FindStringIndex(line)
			if index != nil {
				numbersIndexes = append(numbersIndexes, index)
			}
		}
	}
	// testing:
	for _, coord := range numbersIndexes {
		fmt.Println(coord)
	}

	// Get indexes of symbols around the numbers Symbols are anything that
	// isn't a number or a dot '.'
	for _, line := range file {
		r, _ := regexp.Compile(`[^0-9.]`)
		matches := r.FindAllStringIndex(line, -1)
		symbolsIndexes = append(symbolsIndexes, matches...)
	}

	return numbersIndexes, symbolsIndexes
}

func getNumbers(matrix []string) []string {
	var numbers []string
	r, _ := regexp.Compile(`\d+`)

	for _, el := range matrix {
		matches := r.FindAllString(el, -1)
		numbers = append(numbers, matches...)
	}

	return numbers
}

// Check the matrix for symbols near the numbers
// Returns the coordinates of the symbols in an array
func checkSymbols(schematic [][]string) [][]int {
	var symbols [][]int

	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			for _, dir := range directions {
				coord, err := checkDirection(i, j, dir, schematic)
				if err != nil {
					continue
				}
				if isSymbol(schematic, coord) {
					symbols = append(symbols, coord)
				}
			}
		}
	}

	return symbols
}

func isSymbol(slice [][]string, coords []int) bool {
	switch slice[coords[0]][coords[1]] {
	case "1":
		return false
	case "2":
		return false
	case "3":
		return false
	case "4":
		return false
	case "5":
		return false
	case "6":
		return false
	case "7":
		return false
	case "8":
		return false
	case "9":
		return false
	case ".":
		return false
	default:
		return true
	}
}

// Check if the direction we are trying to look into is a valid location
// in the given matrix
func checkDirection(y int, x int, dir []int, slice [][]string) ([]int, error) {
	newDir := []int{y + dir[0], x + dir[1]}
	height := len(slice) - 1    // These values need a -1 to
	length := len(slice[0]) - 1 // account for index value being out of range

	if newDir[0] > height || newDir[0] < 0 {
		return []int{0, 0}, errors.New("Out of bounds vertically")
	}
	if newDir[1] > length || newDir[1] < 0 {
		return []int{0, 0}, errors.New("Out of bounds horizontally")
	}

	return newDir, nil
}

// Turn the file string array into a matrix
func getMatrix(slice []string) [][]string {
	var grid [][]string

	for _, el := range slice {
		grid = append(grid, strings.Split(el, ""))
	}

	return grid
}
