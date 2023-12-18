package main

func calcPower(red int, green int, blue int) int {
	return red * green * blue
}

func calcMinCubes(game Game) (int, int, int) {
	var (
		red_min int = 0
		green_min int = 0
		blue_min int = 0
	)

	// Check rounds
	for _, round := range game.rounds {
		for _, cube := range round {
			if red_min < cube.number && cube.color == Red{
				red_min = cube.number
			}
			if green_min < cube.number && cube.color == Green{
				green_min = cube.number
			}
			if blue_min < cube.number && cube.color == Blue{
				blue_min = cube.number
			}
		}
	}

	return red_min, green_min, blue_min
}

func calcSumPowerSets(games []Game) int {
	sum := 0

	for _, game := range games {
		red_min, green_min, blue_min := calcMinCubes(game)
		power := calcPower(red_min, green_min, blue_min)

		sum += power
	}

	return sum
}
