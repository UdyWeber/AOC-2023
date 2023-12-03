package day_three

import (
	"fmt"
	"github.com/UdyWeber/GodventOfCode/utils"
	"strconv"
	"strings"
	"unicode"
)

var lines []string
var adjacentsSum int

func init() {
	data := utils.FileReader("day_three/puzzle.txt")
	lines = strings.Split(data, "\r\n")
}

func containsSpecialCharacters(s string) bool {
	for _, char := range s {
		if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			return true
		}
	}
	return false
}

func addAdjacents(lineIdx int, charIdx int) {
	adjacentsOf := string(lines[lineIdx][charIdx])
	lastNum := ""

	for line := lineIdx - 1; line <= lineIdx+1; line++ {
		for column := charIdx - 1; column <= charIdx+1; column++ {
			runeAt := lines[line][column]
			charAt := string(runeAt)

			if charAt == "." || charAt == adjacentsOf {
				continue
			}

			// Has to compose the number
			if unicode.IsNumber(rune(runeAt)) {
				fullNum := charAt

				for i := column - 1; i >= 0; i-- {
					backChar := string(lines[line][i])

					if backChar == "." || backChar == adjacentsOf {
						break
					}

					fullNum = backChar + fullNum
				}

				for i := column + 1; i < len(lines[line]); i++ {
					frontChar := string(lines[line][i])

					if frontChar == "." || frontChar == adjacentsOf {
						break
					}

					fullNum += string(lines[line][i])
				}

				if fullNum == lastNum {
					continue
				}

				number, _ := strconv.Atoi(fullNum)
				adjacentsSum += number
				lastNum = fullNum
			}

		}
	}
}

func addAdjacentsWithGearRation(lineIdx int, charIdx int) {
	adjacentsOf := string(lines[lineIdx][charIdx])
	foundNums := map[string]int{}

	for line := lineIdx - 1; line <= lineIdx+1; line++ {
		for column := charIdx - 1; column <= charIdx+1; column++ {
			runeAt := lines[line][column]
			charAt := string(runeAt)

			if charAt == "." || charAt == adjacentsOf {
				continue
			}

			// Has to compose the number
			if unicode.IsNumber(rune(runeAt)) {
				fullNum := charAt

				for i := column - 1; i >= 0; i-- {
					backChar := string(lines[line][i])

					if backChar == "." || backChar == adjacentsOf {
						break
					}

					fullNum = backChar + fullNum
				}

				for i := column + 1; i < len(lines[line]); i++ {
					frontChar := string(lines[line][i])

					if frontChar == "." || frontChar == adjacentsOf {
						break
					}

					fullNum += string(lines[line][i])
				}

				_, exists := foundNums[fullNum]

				if exists {
					continue
				}

				number, _ := strconv.Atoi(fullNum)
				foundNums[fullNum] = number
			}
		}
	}

	ratioSum := 0
	if len(foundNums) == 2 {
		ratioSum = 1

		for _, val := range foundNums {
			ratioSum *= val
		}
	}

	adjacentsSum += ratioSum
}

func SolutionPart1() {
	for idxLine, line := range lines {
		for idxRune, intChar := range line {
			char := string(intChar)

			if char != "." && containsSpecialCharacters(char) {
				addAdjacents(idxLine, idxRune)
			}

		}

	}

	fmt.Println(adjacentsSum)
}

func SolutionPart2() {
	for idxLine, line := range lines {
		for idxRune, intChar := range line {
			char := string(intChar)

			if char != "." && containsSpecialCharacters(char) {
				addAdjacentsWithGearRation(idxLine, idxRune)
			}

		}

	}

	fmt.Println(adjacentsSum)
}
