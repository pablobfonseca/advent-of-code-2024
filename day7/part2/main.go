package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type memoKey struct {
	index        int
	currentValue int
}

func perform(slice []int, target int) bool {
	memo := make(map[memoKey]bool)
	return check(slice, 1, slice[0], target, memo)
}

func check(slice []int, index, currentValue, target int, memo map[memoKey]bool) bool {
	if index == len(slice) {
		return currentValue == target
	}

	key := memoKey{index, currentValue}
	if result, exists := memo[key]; exists {
		return result
	}

	nextValue := slice[index]

	add := check(slice, index+1, currentValue+nextValue, target, memo)
	multiply := check(slice, index+1, currentValue*nextValue, target, memo)

	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentValue, nextValue))
	concat := check(slice, index+1, concatenated, target, memo)

	memo[key] = add || multiply || concat
	return memo[key]
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	total := 0
	var results []int

	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int
		arr := strings.Split(line, ":")
		result, _ := strconv.Atoi(arr[0])

		nums := strings.Split(strings.TrimPrefix(arr[1], " "), " ")

		for _, n := range nums {
			parsed, _ := strconv.Atoi(n)
			numbers = append(numbers, parsed)
		}

		if perform(numbers, result) {
			total += result
			results = append(results, result)
		}

	}
	fmt.Printf("Total: %d\n", total)
}
