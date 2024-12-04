package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open the file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		matches := parseMatches(line)

		stringSlices := convertToStringSlices(matches)

		for _, val := range stringSlices {
			numRe := regexp.MustCompile(`[0-9]+`)

			matches := numRe.FindAll([]byte(val), -1)
			intSlices := convertToIntSlices(matches)

			result := 1
			for _, val := range intSlices {
				result *= val
			}

			total += result
		}
	}

	fmt.Printf("Total: %d", total)
}

func convertToStringSlices(matches [][]byte) []string {
	stringSlices := make([]string, len(matches))

	for i, b := range matches {
		stringSlices[i] = string(b)
	}

	return stringSlices
}

func convertToIntSlices(matches [][]byte) []int {
	intSlices := make([]int, len(matches))

	for i, b := range matches {
		val, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatalf("Error with int: %v", err)
		}
		intSlices[i] = val
	}

	return intSlices
}

func parseMatches(line string) (matches [][]byte) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matches = re.FindAll([]byte(line), -1)

	return
}
