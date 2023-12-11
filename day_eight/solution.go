package day_eight

import (
	"fmt"
	"github.com/UdyWeber/GodventOfCode/utils"
	"strings"
)

func SolutionPart1() {
	data := strings.Split(utils.FileReader("day_eight/puzzle.txt"), "\r\n\r\n")

	coordinates := make(map[string]map[string]string)

	instructions := data[0]
	places := data[1]

	for _, place := range strings.Split(places, "\r\n") {
		placeData := strings.Split(place, " = ")

		leftRight := placeData[1]
		for _, char := range "()," {
			leftRight = strings.ReplaceAll(leftRight, string(char), "")
		}
		leftRightFields := strings.Fields(leftRight)

		left := leftRightFields[0]
		right := leftRightFields[1]

		coordinates[placeData[0]] = map[string]string{"L": left, "R": right}
	}

	current := "AAA"
	steps := 0

	for current != "ZZZ" {
		for _, s := range instructions {
			if current == "ZZZ" {
				break
			}
			current = coordinates[current][string(s)]
			steps++
		}
	}

	fmt.Println(current, steps)
}

func SolutionPart2() {
	data := strings.Split(utils.FileReader("day_eight/puzzle.txt"), "\r\n\r\n")

	coordinates := make(map[string]map[string]string)

	instructions := data[0]
	places := data[1]

	for _, place := range strings.Split(places, "\r\n") {
		placeData := strings.Split(place, " = ")

		leftRight := placeData[1]
		for _, char := range "()," {
			leftRight = strings.ReplaceAll(leftRight, string(char), "")
		}
		leftRightFields := strings.Fields(leftRight)

		left := leftRightFields[0]
		right := leftRightFields[1]

		coordinates[placeData[0]] = map[string]string{"L": left, "R": right}
	}

	var currents []string

	for key := range coordinates {
		if string(key[len(key)-1]) == "A" {
			currents = append(currents, key)
		}
	}

	// Storing the steps that it took for each cycle
	var nSteps []int

	for _, current := range currents {
		steps := 0

		for string(current[len(current)-1]) != "Z" {
			for _, char := range instructions {
				current = coordinates[current][string(char)]
				steps++
				if string(current[len(current)-1]) == "Z" {
					break
				}

			}

		}

		fmt.Println(current, steps)
		nSteps = append(nSteps, steps)
	}

	// Calculates the Lest Common Multiple MMC(In Portuguese) of each one and gets the resul
	fmt.Println(utils.LCM(nSteps[0], nSteps[1], nSteps...))
}
