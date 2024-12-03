package main

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input1.txt")

	if err != nil {
		panic(err)
	}

	list := strings.Fields(string(input))

	var firstList, secondList []int

	for i, val := range list {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			firstList = append(firstList, intVal)
		} else {
			secondList = append(secondList, intVal)
		}
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	var distances []int

	for i := 0; i < len(firstList); i++ {
		distance := math.Abs(float64(firstList[i] - secondList[i]))
		distances = append(distances, int(distance))
	}

	frequencyMap := make(map[int]int)

	for _, val := range secondList {
		frequencyMap[val]++
	}

	score := 0
	for _, val := range firstList {
		count := frequencyMap[val]
		score += val * count
	}

	total := 0

	for _, val := range distances {
		total += val
	}
}
