package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	number int
	color  string
}

type Game struct {
	id int
	rounds [][]Cube
}

const (
	Red   string = "red"
	Blue  string = "blue"
	Green string = "green"
)

func calcGamesIdSum(games []Game) int {
	result := 0

	for _, game := range games {
		if isGamePossible(game) {
			result += game.id
		}
	}

	return result
}

func isGamePossible(game Game) bool {
	result := true
	max_red := 12
	max_green := 13
	max_blue := 14

	for _, rounds := range game.rounds {
		for _, cube := range rounds {
			if cube.color == Red && cube.number > max_red {
				result = false
			}
			if cube.color == Green && cube.number > max_green {
				result = false
			}
			if cube.color == Blue && cube.number > max_blue {
				result = false
			}
		}
	}

	return result
}

func makeGame(line string) Game {
	var game Game
	var cubes [][]Cube

	game.id = getGameId(line)

	str := delGameId(line)
	rounds := makeRoundsStr(str)
	for _, round := range rounds {
		cubes = append(cubes, makeCubes(round))
	}
	game.rounds = cubes

	return game
}

func getGameId(line string) int {
	lineSlice := strings.Split(line, " ")
	str := strings.ReplaceAll(lineSlice[1], ":", "")
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return result
}

// Turn string into Cubes (ONLY USE AFTER REMOVING GAME ID AND MAKING ROUNDS ARRAY)
func makeCubes(str string) []Cube {
	var cubes []Cube

	fields := strings.Split(str, ",")
	for i := 0; i < len(fields); i++ {
		fields[i] = strings.Trim(fields[i], " ")
	}

	for _, el := range fields {
		squareSlice := strings.Split(el, " ")

		cubeNumber, err := strconv.Atoi(squareSlice[0])
		if err != nil {
			panic(err)
		}

		cube := Cube{cubeNumber, squareSlice[1]}
		cubes = append(cubes, cube)
	}

	return cubes
}

// Remove game id
func delGameId(line string) string {
	_, str, _ := strings.Cut(line, ":")
	str = strings.Trim(str, " ")

	return str
}

// Make string into rounds string array
func makeRoundsStr(line string) []string {
	var rounds []string

	rounds = strings.Split(line, ";")

	for i := 0; i < len(rounds); i++ {
		rounds[i] = strings.Trim(rounds[i], " ")
	}

	return rounds
}

// Loads a given file into a string slice
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
