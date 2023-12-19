package main

import "fmt"

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

func checkSymbols(slice []string) {
    for i := 0; i < len(slice); i++ {
	fmt.Println(slice[i])
    }
}
