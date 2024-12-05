package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var directions = [8][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
	{1, 1}, {1, -1},
	{-1, 1}, {-1, -1},
}

const word = "XMAS"

func countOccurrences(grid [][]rune) int {
	wordLen := len(word)
	wordRunes := []rune(word)
	rows, cols := len(grid), len(grid[0])
	count := 0

	checkWord := func(row, col, dRow, dCol int) bool {
		for i := 0; i < wordLen; i++ {
			newRow, newCol := row+i*dRow, col+i*dCol
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || grid[newRow][newCol] != wordRunes[i] {
				return false
			}
		}
		return true
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == wordRunes[0] {
				for _, dir := range directions {
					if checkWord(row, col, dir[0], dir[1]) {
						count++
					}
				}
			}
		}
	}

	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error with file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, []rune(line))
	}

	result := countOccurrences(grid)
	fmt.Printf("Total %d\n", result)
}
