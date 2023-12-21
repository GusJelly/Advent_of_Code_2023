package main

import (
    "strings"
)

var directions = map[string][]int {
    "up":        {0, 1},
    "down":      {0, -1},
    "center":    {0, 0},
    "left":      {-1, 0},
    "right":     {0, 1},
    "downRight": {1, 1},
    "upRight":   {-1, 1},
    "upLeft":    {-1, -1},
    "downLeft":  {1, -1},
}

func getMatrix(slice []string) [][]string {
    var grid [][]string

    for i := 0; i < len(slice); i++ {
	grid = append(grid, strings.Split(slice[i], ""))
    }

    return grid
}

func checkDirections(slice [][]string) {

}
