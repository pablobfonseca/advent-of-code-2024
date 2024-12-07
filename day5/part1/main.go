package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	var ordersToPrint [][]int

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
					continue instructionsLoop
				}
			}
		}
		ordersToPrint = append(ordersToPrint, instructionArr)
	}

	total := 0

	for _, order := range ordersToPrint {
		middle := order[len(order)/2]

		total += middle
	}

	fmt.Println("Total", total)
}

func convertToIntSlice(strSlice []string) (intSlice []int) {
	for _, str := range strSlice {
		val, _ := strconv.Atoi(str)
		intSlice = append(intSlice, val)
	}
	return

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
