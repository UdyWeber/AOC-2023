package day_four

import (
	"fmt"
	"github.com/UdyWeber/GodventOfCode/utils"
	"strconv"
	"strings"
)

func SolutionPart1() {
	data := utils.FileReader("day_four/puzzle.txt")
	games := strings.Split(data, "\r\n")

	pointsSum := 0

	for _, game := range games {
		gamePoints := 0
		scores := strings.Split(game, ":")[1]

		splitedScore := strings.Split(scores, "|")

		winningNumbers := strings.Fields(splitedScore[0])
		myNumbers := strings.Fields(splitedScore[1])

		for _, winningNumber := range winningNumbers {
			for _, myNumber := range myNumbers {
				if winningNumber == myNumber {
					if gamePoints == 0 {
						gamePoints++
					} else {
						gamePoints *= 2
					}

					break
				}
			}
		}

		pointsSum += gamePoints
	}

	fmt.Println(pointsSum)
}

func SolutionPart2() {
	data := utils.FileReader("day_four/puzzle.txt")
	games := strings.Split(data, "\r\n")

	scratchCardsGained := map[int]int{}

	for _, game := range games {
		countedNums := 0

		separatedData := strings.Split(game, ":")
		cardData := separatedData[0]
		cardId, _ := strconv.Atoi(strings.Fields(cardData)[1])

		scores := separatedData[1]

		splitedScore := strings.Split(scores, "|")

		winningNumbers := strings.Fields(splitedScore[0])
		myNumbers := strings.Fields(splitedScore[1])

		for _, winningNumber := range winningNumbers {
			for _, myNumber := range myNumbers {
				if winningNumber == myNumber {
					countedNums++
				}
			}
		}

		value, exists := scratchCardsGained[cardId]
		value++

		if exists {
			scratchCardsGained[cardId] += 1
		} else {
			scratchCardsGained[cardId] = 1
		}

		for i := 1; i <= countedNums; i++ {
			scratchCardsGained[i+cardId] += value
		}
	}

	totalSum := 0

	for _, val := range scratchCardsGained {
		totalSum += val
	}

	fmt.Println(totalSum)
}
