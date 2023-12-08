package day_five

import (
	"fmt"
	"github.com/UdyWeber/GodventOfCode/utils"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func b_search(val int, vals []int) int {
	low := 0
	high := len(vals) - 1

	for low <= high {
		mid := int(math.Floor(float64((low + high) / 2)))
		result := vals[mid]

		if result == val {
			return mid
		} else if result > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func StringArrToIntArr(stringArray []string) []int {
	intArr := make([]int, len(stringArray))

	for i, numString := range stringArray {
		number, _ := strconv.Atoi(numString)
		intArr[i] = number
	}

	return intArr
}

func getEquivalent(initial int, mapper *[]string) int {
	newVal := initial

	for _, mapperInfo := range *mapper {
		infos := StringArrToIntArr(strings.Fields(mapperInfo))
		destinationEq, sourceEq, rangeEq := infos[0], infos[1], infos[2]

		if newVal <= (sourceEq+rangeEq) && newVal >= sourceEq {
			for j := 0; j < rangeEq; j++ {
				if initial == sourceEq+j {
					newVal = destinationEq + j
				}
			}
		}

		if newVal != initial {
			break
		}
	}

	return newVal
}

func SolutionPart1() {
	data := utils.FileReader("day_five/puzzle.txt")
	lines := strings.Split(data, "\r\n\r\n")

	seeds := StringArrToIntArr(strings.Fields(strings.Split(lines[0], ":")[1]))

	lowest := 0

	for i, seed := range seeds {
		soilMap := strings.Split(lines[1], "\r\n")[1:]
		soilEq := getEquivalent(seed, &soilMap)

		fertilizerMap := strings.Split(lines[2], "\r\n")[1:]
		fertilizerEq := getEquivalent(soilEq, &fertilizerMap)

		waterMap := strings.Split(lines[3], "\r\n")[1:]
		waterEq := getEquivalent(fertilizerEq, &waterMap)

		lightMap := strings.Split(lines[4], "\r\n")[1:]
		lightEq := getEquivalent(waterEq, &lightMap)

		temperatureMap := strings.Split(lines[5], "\r\n")[1:]
		temperatureEq := getEquivalent(lightEq, &temperatureMap)

		humidityMap := strings.Split(lines[6], "\r\n")[1:]
		humidityEq := getEquivalent(temperatureEq, &humidityMap)

		locationMap := strings.Split(lines[7], "\r\n")[1:]
		locationEq := getEquivalent(humidityEq, &locationMap)

		if lowest == 0 || locationEq < lowest {
			lowest = locationEq
		}

		fmt.Println(i+1, "/", len(seeds), seed, soilEq, fertilizerEq, waterEq, lightEq, temperatureEq, humidityEq, locationEq)
	}

	fmt.Println("Lowest val was: ", lowest)
}

type mapEntry struct {
	destinationEq, sourceEq, rangeEq int
}

type entireMap []mapEntry

var mappers []entireMap

func processSeed(seedIndex, seedSource, seedRange int, results *chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for seed := seedSource; seed < seedSource+seedRange; seed++ {
		locationId := seed

		for _, someMap := range mappers {
			for _, r := range someMap {
				if locationId >= r.sourceEq && locationId < r.sourceEq+r.rangeEq {
					locationId = r.destinationEq + (locationId - r.sourceEq)
					break
				}
			}
		}

		fmt.Println("Thread", seedIndex, seed, locationId)

		*results <- locationId
	}
}

func SolutionPart2() {
	data := utils.FileReader("day_five/puzzle.txt")
	lines := strings.Split(data, "\r\n\r\n")

	seeds := StringArrToIntArr(strings.Fields(strings.Split(lines[0], ":")[1]))

	mappers = make([]entireMap, len(lines[1:]))

	for i := 0; i < len(mappers); i++ {
		linesData := strings.Split(lines[i+1], "\r\n")[1:]
		mappedLines := make(entireMap, len(linesData))

		for j := 0; j < len(linesData); j++ {
			infos := StringArrToIntArr(strings.Fields(linesData[j]))

			mappedLines[j] = mapEntry{
				destinationEq: infos[0],
				sourceEq:      infos[1],
				rangeEq:       infos[2],
			}
		}

		sort.Slice(mappedLines, func(i, j int) bool {
			return mappedLines[i].sourceEq < mappedLines[j].sourceEq
		})

		mappers[i] = mappedLines
	}

	var wg sync.WaitGroup

	results := make(chan int, len(seeds))

	for i := 0; i < len(seeds); i += 2 {
		seedSource, seedRange := seeds[i], seeds[i+1]

		wg.Add(1)
		go processSeed(i+2, seedSource, seedRange, &results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	lowest := math.MaxInt64

	for result := range results {
		if result < lowest {
			lowest = result
		}
	}

	fmt.Println("Lowest val was: ", lowest)
}
