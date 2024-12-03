package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	values := parseInput(strings.Fields(string(input)))

	firstList, secondList := splitLists(values)

	sort.Ints(firstList)
	sort.Ints(secondList)

	distances := calculateDistances(firstList, secondList)

	frequencyMap := buildFrequencyMap(secondList)
	score := calculateScore(firstList, frequencyMap)

	totalDistance := sumSlice(distances)

	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Total Distance: %d\n", totalDistance)
}

func parseInput(stringsList []string) []int {
	values := make([]int, len(stringsList))
	for i, str := range stringsList {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Invalid number in input: %v", err)
		}
		values[i] = val
	}
	return values
}

func splitLists(values []int) (firstList, secondList []int) {
	firstList = make([]int, 0, len(values)/2)
	secondList = make([]int, 0, len(values)/2)
	for i, val := range values {
		if i%2 == 0 {
			firstList = append(firstList, val)
		} else {
			secondList = append(secondList, val)
		}
	}
	return
}

func calculateDistances(firstList, secondList []int) []int {
	distances := make([]int, len(firstList))
	for i := range firstList {
		distances[i] = int(math.Abs(float64(firstList[i] - secondList[i])))
	}
	return distances
}

func buildFrequencyMap(values []int) map[int]int {
	frequencyMap := make(map[int]int)
	for _, val := range values {
		frequencyMap[val]++
	}
	return frequencyMap
}

func calculateScore(firstList []int, frequencyMap map[int]int) int {
	score := 0
	for _, val := range firstList {
		score += val * frequencyMap[val]
	}
	return score
}

func sumSlice(values []int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
