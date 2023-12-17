package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	red int
	green int
	blue int
}

type Game struct {
	id int
	rounds []Round
}

func main() {
	lines := readFile("input")
	for _, line := range lines {
		fmt.Println()
		fmt.Println("getGameId: ", getGameId(tokenize(line)))
		fmt.Println("getRounds: ", getRounds(tokenize(line)))
	}
}

// Gets the different rounds in a game
func getRounds(tokens []string) []Round {
	var rounds []Round

	for i := 2; i < len(tokens); i++ {
		fmt.Printf("%s ", tokens[i])
	}

	return rounds
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

func tokenize(line string) []string {
	tokens := strings.Split(line, " ")
	return tokens
}
