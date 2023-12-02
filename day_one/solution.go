package day_one

import (
	"fmt"
	"github.com/UdyWeber/GodventOfCode/utils"
	"strconv"
	"strings"
	"unicode"
)

var names map[string]string

func init() {
	names = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
}

func getValNumeric(keyWord string) string {
	value, exists := names[keyWord]

	if !exists {
		return ""
	}

	return value
}

func Solution() {
	data := utils.FileReader("day_one/puzzle.txt")
	sum := 0

	lines := strings.Split(data, "\r\n")

	for _, word := range lines {

		var firstNum string
		for i := 0; i < len(word); i++ {
			aRune := word[i]

			if unicode.IsNumber(rune(aRune)) {
				firstNum = string(aRune)
				break
			}

			numWord := string(aRune)

			for j := i + 1; j < len(word)-1; j++ {
				newRune := word[j]
				numWord += string(newRune)

				if value := getValNumeric(numWord); value != "" {
					firstNum = value
					break
				}
			}

			if firstNum != "" {
				break
			}

		}

		var secondNum string
		for i := len(word) - 1; i >= 0; i-- {
			aRune := word[i]

			if unicode.IsNumber(rune(aRune)) {
				secondNum = string(aRune)
				break
			}

			numWord := string(aRune)

			for j := i - 1; j >= 0; j-- {
				newRune := word[j]

				numWord = string(newRune) + numWord

				if value := getValNumeric(numWord); value != "" {
					secondNum = value
					break
				}
			}

			if secondNum != "" {
				break
			}

		}

		fmt.Printf("%s %s from (%s)\n", firstNum, secondNum, word)
		conv, _ := strconv.Atoi(firstNum + secondNum)
		sum += conv
	}

	fmt.Println(sum)
}
