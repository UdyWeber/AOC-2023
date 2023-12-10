package day_six

import (
	"fmt"
	"github.com/UdyWeber/GodventOfCode/utils"
	"strconv"
	"strings"
)

type race struct {
	raceDuration    int
	raceMaxDistance int
}

func formatRaces() []race {
	data := utils.FileReader("day_six/puzzle.txt")
	datas := strings.Split(data, "\r\n")

	times := strings.Fields(strings.Split(datas[0], ":")[1])
	distances := strings.Fields(strings.Split(datas[1], ":")[1])

	races := make([]race, len(times))

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		races[i] = race{
			raceDuration:    time,
			raceMaxDistance: distance,
		}
	}

	return races
}

func getTotalWins(theRace *race) int {
	totalWins := 0

	for time := theRace.raceDuration - 1; time > 0; time-- {
		leftover := theRace.raceDuration - time

		if time*leftover > theRace.raceMaxDistance {
			totalWins++
		}
	}

	return totalWins
}

func SolutionPart1() {
	races := formatRaces()

	multipliedRaces := 1

	for _, theRace := range races {
		totalWins := getTotalWins(&theRace)
		multipliedRaces *= totalWins
	}

	fmt.Println(multipliedRaces)
}

func SolutionPart2() {
	data := utils.FileReader("day_six/puzzle.txt")
	datas := strings.Split(data, "\r\n")

	times := strings.Fields(strings.Split(datas[0], ":")[1])
	distances := strings.Fields(strings.Split(datas[1], ":")[1])

	actualTime := ""
	for _, time := range times {
		actualTime += time
	}

	time, _ := strconv.Atoi(actualTime)

	actualDistance := ""
	for _, distance := range distances {
		actualDistance += distance
	}

	distance, _ := strconv.Atoi(actualDistance)

	totalWins := getTotalWins(&race{
		raceDuration:    time,
		raceMaxDistance: distance,
	})

	fmt.Println(totalWins)
}
