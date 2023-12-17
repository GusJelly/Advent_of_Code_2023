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
	color string
}

const (
	Red   string = "red"
	Blue  string = "blue"
	Green string = "green"
)

func main() {
	lines := loadFile("input")
	gameStr := delGameId(lines[0])
	rounds := makeRoundsStr(gameStr)

	cubes := makeCubes(rounds[0])
	for _, i := range cubes {
		fmt.Println(i)
	}
}

// Turn string into Cubes
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

		cube := Cube {cubeNumber, squareSlice[1]}
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
