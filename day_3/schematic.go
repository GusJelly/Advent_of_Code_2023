package main

import (
	"errors"
	"fmt"
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

func findNumbers(schematic [][]string) []int {
	var result []int

	return result
}

// Check the matrix for symbols near the numbers
// Returns the coordinates of the symbols in an array
func checkSymbols(schematic [][]string) {
	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			for _, dir := range directions {
				coord, err := checkDirection(i, j, dir, schematic)
				if err != nil {
					continue
				}
				fmt.Println(schematic[coord[0]][coord[1]])
			}
		}
	}
}

// Check if the direction we are trying to look into is a valid location
// in the given matrix
func checkDirection(y int, x int, dir []int, slice [][]string) ([]int, error) {
	newDir := []int{y + dir[0], x + dir[1]}
	height := len(slice)
	length := len(slice[0])

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
