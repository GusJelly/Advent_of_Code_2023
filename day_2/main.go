package main

import "fmt"

func main() {
	lines := loadFile("input")
	var games []Game

	for _, line := range lines {
		games = append(games, makeGame(line))
	}

	// Part one
	fmt.Println(calcGamesIdSum(games))

	// Part two
	fmt.Println(calcSumPowerSets(games))
}
