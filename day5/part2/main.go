package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error with file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules [][]int
	var instructions [][]int
	var ordersToFix [][]int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")
			rules = append(rules, convertToIntSlice(rule))
		} else if line == "" {
			continue
		} else {
			instruction := strings.Split(line, ",")
			instructions = append(instructions, convertToIntSlice(instruction))
		}

	}

	beforeRuleMap := make(map[int][][]int)
	afterRuleMap := make(map[int][][]int)

	for _, rulesArr := range rules {
		beforeRuleMap[rulesArr[0]] = append(beforeRuleMap[rulesArr[0]], rulesArr)
		afterRuleMap[rulesArr[0]] = append(afterRuleMap[rulesArr[0]], rulesArr)
	}

	filterRule := func(instruction int) (filteredRules [][]int) {
		return append(beforeRuleMap[instruction], afterRuleMap[instruction]...)
	}

instructionsLoop:
	for _, instructionArr := range instructions {
		for instructionIndex, instruction := range instructionArr {
			currentInstructionRules := filterRule(instruction)

			for _, rule := range currentInstructionRules {
				if violatesRule(instructionArr, instructionIndex, rule) {
					ordersToFix = append(ordersToFix, instructionArr)
					continue instructionsLoop
				}
			}
		}
	}

	precedence := make(map[int][]int)
	for _, rule := range rules {
		precedence[rule[1]] = append(precedence[rule[1]], rule[0])
	}

	comesBefore := func(a, b int) bool {
		for _, p := range precedence[b] {
			if p == a {
				return true
			}
		}
		return false
	}

	sortSubSlice := func(subSlice []int) {
		sort.SliceStable(subSlice, func(i, j int) bool {
			a, b := subSlice[i], subSlice[j]

			if comesBefore(a, b) {
				return true
			}

			if comesBefore(b, a) {
				return false
			}

			return false
		})
	}

	for _, order := range ordersToFix {
		sortSubSlice(order)
	}

	total := 0

	for _, order := range ordersToFix {
		middle := order[len(order)/2]

		total += middle
	}

	fmt.Println("Total", total)

}

func violatesRule(instructions []int, currentIndex int, rule []int) bool {
	before, after := rule[0], rule[1]
	instruction := instructions[currentIndex]

	if instruction == before {
		afterIndex := slices.Index(instructions, after)
		return afterIndex != -1 && currentIndex > afterIndex
	}

	return false
}

func convertToIntSlice(strSlice []string) (intSlice []int) {
	for _, str := range strSlice {
		val, _ := strconv.Atoi(str)
		intSlice = append(intSlice, val)
	}
	return

}
