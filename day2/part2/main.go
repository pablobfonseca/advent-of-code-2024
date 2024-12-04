package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		nums := convertLineToArray(line)

		reports = append(reports, nums)
	}

	safeCount := 0

	for _, report := range reports {
		if isSafeWithDampener(report) {
			safeCount++
		}
	}

	fmt.Printf("Safe count: %d\n", safeCount)
}

func isSafe(report []int) bool {
	length := len(report)
	if length < 2 {
		return false
	}

	increasing := true
	decreasing := true

	for i := 1; i < length; i++ {
		diff := report[i] - report[i-1]
		absDiff := math.Abs(float64(diff))

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if diff > 0 {
			decreasing = false
		}
		if diff < 0 {
			increasing = false
		}
	}

	return increasing || decreasing
}

func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		modifiedReport := make([]int, 0, len(report)-1)
		modifiedReport = append(modifiedReport, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)

		if isSafe(modifiedReport) {
			return true
		}
	}

	return false
}

func convertLineToArray(line string) []int {
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, field := range fields {
		if field == "" {
			continue
		}
		num, err := strconv.Atoi(field)
		if err != nil {
			log.Fatalf("Invalid number: %q: %v", field, err)
		}
		nums[i] = num
	}

	return nums
}
