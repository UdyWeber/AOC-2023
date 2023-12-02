package day_two

import (
	"fmt"
	"github.com/UdyWeber/GodventOfCode/utils"
	"strconv"
	"strings"
)

func part1() {

}

func SolutionPart1() {
	data := utils.FileReader("day_two/puzzle.txt")
	lines := strings.Split(data, "\r\n")

	gamesSum := 0

	for _, line := range lines {
		game := strings.Split(line, ":")
		rounds := strings.Split(game[1], ";")

		possible := true

		for _, round := range rounds {
			counting := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}

			moves := strings.Split(round, ",")

			for _, move := range moves {
				parts := strings.Split(move, " ")
				number, _ := strconv.Atoi(parts[1])

				counting[parts[2]] += number
			}

			if counting["red"] > 12 || counting["green"] > 13 || counting["blue"] > 14 {
				possible = false
				break
			}
		}

		if possible {
			gameId, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
			gamesSum += gameId
		}
	}

	fmt.Println(gamesSum)
}

func SolutionPart2() {
	data := utils.FileReader("day_two/puzzle.txt")
	lines := strings.Split(data, "\r\n")

	gamesSum := 0

	for _, line := range lines {
		game := strings.Split(line, ":")
		rounds := strings.Split(game[1], ";")

		minimum := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range rounds {

			moves := strings.Split(round, ",")

			for _, move := range moves {
				parts := strings.Split(move, " ")
				number, _ := strconv.Atoi(parts[1])

				minimumForValue := minimum[parts[2]]

				if minimumForValue == 0 || number > minimumForValue {
					minimum[parts[2]] = number
				}
			}
		}

		gamesSum += minimum["red"] * minimum["green"] * minimum["blue"]
	}

	fmt.Println(gamesSum)
}
