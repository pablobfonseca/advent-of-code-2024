package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isXPattern(grid [][]rune, row, col int) bool {
	rowsLen, colsLen := len(grid), len(grid[0])

	if row-1 < 0 || row+1 >= rowsLen || col-1 < 0 || col+1 >= colsLen {
		return false
	}
	diag1Top := grid[row-1][col-1]
	diag1Bottom := grid[row+1][col+1]
	diag1Match := (diag1Top == 'M' && diag1Bottom == 'S') || (diag1Top == 'S' && diag1Bottom == 'M')

	diag2Top := grid[row-1][col+1]
	diag2Bottom := grid[row+1][col-1]
	diag2Match := (diag2Top == 'M' && diag2Bottom == 'S') || (diag2Top == 'S' && diag2Bottom == 'M')

	return diag1Match && diag2Match
}

func countOccurrences(grid [][]rune) int {
	rowsLen, colsLen := len(grid), len(grid[0])
	count := 0

	for row := 0; row < rowsLen; row++ {
		for col := 0; col < colsLen; col++ {
			if grid[row][col] == 'A' && isXPattern(grid, row, col) {
				count++
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
