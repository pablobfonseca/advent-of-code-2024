package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open the file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mulEnabled := true

	var matchesToProcess []string
	for scanner.Scan() {
		line := scanner.Text()

		parsedMatches := parseInstructions(line)

		for _, match := range parsedMatches {
			if match == "don't()" {
				mulEnabled = false
			}
			if match == "do()" {
				mulEnabled = true
			}

			if mulEnabled && strings.HasPrefix(match, "mul(") {
				matchesToProcess = append(matchesToProcess, match)
			}
		}
	}

	processValues(matchesToProcess)
}

func processValues(matches []string) int {
	total := 0

	for _, val := range matches {
		numRe := regexp.MustCompile(`[0-9]+`)

		matches := numRe.FindAllString(val, -1)
		intSlices := convertToIntSlices(matches)

		result := 1
		for _, val := range intSlices {
			result *= val
		}

		total += result
	}

	return total
}

func convertToIntSlices(matches []string) []int {
	intSlices := make([]int, len(matches))

	for i, b := range matches {
		val, err := strconv.Atoi(b)
		if err != nil {
			log.Fatalf("Error with int: %v", err)
		}
		intSlices[i] = val
	}

	return intSlices
}

func parseInstructions(line string) []string {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	matches := re.FindAllString(line, -1)

	var result []string

	for _, match := range matches {
		result = append(result, match)
	}

	return result
}

func parseMulMatches(str string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matches := re.FindAllString(str, -1)

	return matches
}
