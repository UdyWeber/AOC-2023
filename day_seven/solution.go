package day_seven

import (
	"fmt"
	"github.com/UdyWeber/GodventOfCode/utils"
	"sort"
	"strconv"
	"strings"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_KIND
	FULL_HOUSE
	FOUR_OF_KIND
	FIVE_OF_KIND
)

type hand struct {
	kind  int
	cards string
}

type bet struct {
	bid  int
	hand hand
}

func getAlphabetEquivalent(k string, withJokers bool) string {
	var keys string

	if withJokers {
		keys = "AKQT98765432J"
	} else {
		keys = "AKQJT98765432"
	}

	alphabet := "ABCDEFGHIJKLM"

	newKey := ""

	for _, s := range k {
		for j := 0; j < len(keys); j++ {
			if string(s) == string(keys[j]) {
				newKey += string(alphabet[j])
				break
			}
		}
	}

	return newKey
}

func sortRangesAlphabetically(arr *[]bet, start, end int) {
	if start == end || end-start <= 1 {
		return
	}

	for i := start; i < end; i++ {
		for j := i; j < end; j++ {
			if getAlphabetEquivalent((*arr)[i].hand.cards, false) > getAlphabetEquivalent((*arr)[j].hand.cards, false) {
				temp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}
}

func getHandKind(cards *string, withJokers bool) int {
	countings := make(map[rune]int)

	jokers := 0

	for _, val := range *cards {
		if withJokers {
			if string(val) == "J" {
				jokers++
			} else {
				countings[val]++
			}
		} else {
			countings[val]++
		}
	}

	var highest int
	for _, val := range countings {
		if val > highest {
			highest = val
		}
	}

	switch highest {
	case 5:
		return FIVE_OF_KIND
	case 4:
		return FOUR_OF_KIND
	case 3:
		if len(countings) > 2 {
			return THREE_OF_KIND
		}

		return FULL_HOUSE
	case 2:
		if len(countings) == 3 {
			return TWO_PAIR
		}

		return ONE_PAIR
	case 1:
		return HIGH_CARD
	default:
		fmt.Println("not implemented")
	}

	return -1
}

func (h *hand) isStrongerThan(other hand) bool {
	if h.kind != other.kind {
		return h.kind > other.kind
	}
	for i := range h.cards {
		if h.cards[i] != other.cards[i] {
			value := getAlphabetEquivalent(string(h.cards[i]), true)
			otherValue := getAlphabetEquivalent(string(other.cards[i]), true)
			return value < otherValue
		}
	}
	return false
}

func formatBets() []bet {
	lines := strings.Split(utils.FileReader("day_seven/puzzle.txt"), "\r\n")

	bets := make([]bet, len(lines))

	for i, line := range lines {
		data := strings.Fields(line)

		cardString := data[0]

		bid, _ := strconv.Atoi(data[1])

		bets[i] = bet{
			bid: bid,
			hand: hand{
				kind:  getHandKind(&cardString, false),
				cards: cardString,
			},
		}
	}

	sort.Slice(bets, func(i, j int) bool {
		return bets[i].hand.kind < bets[j].hand.kind
	})

	// Sort slices of array that exists in the same category
	currentIdx := 0
	for typeIdx := 0; typeIdx < 7; typeIdx++ {
		for i := currentIdx; i < len(bets); i++ {
			if bets[i].hand.kind != typeIdx {
				sortRangesAlphabetically(&bets, currentIdx, i)
				currentIdx = i
				break
			}
		}

	}

	return bets
}

func formatBetsWithJokers() []bet {
	lines := strings.Split(utils.FileReader("day_seven/puzzle.txt"), "\r\n")

	bets := make([]bet, len(lines))

	// categories := make(map[int][]bet)

	for i, line := range lines {
		data := strings.Fields(line)

		cardString := data[0]

		bid, _ := strconv.Atoi(data[1])
		kind := getHandKind(&cardString, true)

		bets[i] = bet{
			bid: bid,
			hand: hand{
				kind:  kind,
				cards: cardString,
			},
		}

	}

	sort.Slice(bets, func(i, j int) bool { return !(bets[i].hand.isStrongerThan(bets[j].hand)) })

	return bets
}

func SolutionPart1() {
	bets := formatBets()
	total := 0

	for i, bet := range bets {
		fmt.Println(i, bet.hand.cards, bet.hand.kind)
		total += bet.bid * (i + 1)
	}

	fmt.Println(total)
}

func SolutionPart2() {
	bets := formatBetsWithJokers()
	total := 0

	for i, bet := range bets {
		fmt.Println(i, bet.hand.cards, bet.hand.kind)
		total += bet.bid * (i + 1)
	}

	fmt.Println(total)
}
