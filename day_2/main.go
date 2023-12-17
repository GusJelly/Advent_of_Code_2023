package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	number int
	color  string
}

type Round struct {
	id          int
	numberCubes []Cube
}

type Game struct {
	id     int
	rounds []Round
}

func main() {
	lines := loadFile("input")

	getCubeNumber(lines[0])
}

// Get cubeNumber
func getCubeNumber(line string) {
	_, line, _ = strings.Cut(line, ":")
	line = strings.Trim(line, " ")
	rounds := strings.Split(line, ";")

	for i := 0; i < len(rounds); i++ {
		rounds[i] = strings.Trim(rounds[i], " ")
		fmt.Println(rounds[i])
	}
}

// Gets the game's id from the game's tokens slice
func getGameId(tokens []string) int {
	id, err := strconv.Atoi(strings.Trim(tokens[1], ":"))
	if err != nil {
		panic(err)
	}

	return id
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
