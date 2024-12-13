package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type VisitedKey struct {
	row, col int
}

type LoopVisitedKey struct {
	row, col, dirRow, dirCol int
}

type Position struct {
	row int
	col int
}

func (p *Position) add(newPos *Position) Position {
	return Position{
		row: p.row + newPos.row,
		col: p.col + newPos.col,
	}
}

func (p *Position) turnRight() Position {
	pos := Position{
		row: p.col,
		col: -p.row,
	}

	return pos
}

type Grid struct {
	rowLen  int
	colLen  int
	data    [][]rune
	charPos map[rune]Position
}

func NewGrid(data [][]rune) *Grid {
	charPos := make(map[rune]Position)
	for row := range data {
		for col, char := range data[row] {
			charPos[char] = Position{row, col}
		}
	}
	return &Grid{
		rowLen:  len(data),
		colLen:  len(data[0]),
		data:    data,
		charPos: charPos,
	}
}

func (g *Grid) get(pos Position) (bool, string) {
	if pos.row < 0 || pos.col < 0 || pos.row >= g.rowLen || pos.col >= g.colLen {
		return false, ""
	} else {
		return true, string(g.data[pos.row][pos.col])
	}
}

func (g *Grid) find(char rune) (pos Position) {
	return g.charPos[char]
}

func (g *Grid) set(pos Position, char rune) {
	g.data[pos.row][pos.col] = char
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var gridData [][]rune

	for scanner.Scan() {
		lines := scanner.Text()

		gridData = append(gridData, []rune(lines))
	}

	grid := NewGrid(gridData)
	startPos := grid.find('^')
	visited := getVisited(grid, startPos)

	loopCount := 0

	for _, pos := range visited {
		originalChar := grid.data[pos.row][pos.col]

		grid.data[pos.row][pos.col] = '#'

		isLooping := isALoop(grid, startPos)

		if isLooping {
			loopCount++
		}

		grid.data[pos.row][pos.col] = originalChar
	}

	fmt.Printf("Loops: %v\n", loopCount)
	fmt.Printf("Total: %d\n", len(visited))
}

func isALoop(grid *Grid, startPos Position) bool {
	visited := make(map[LoopVisitedKey]int)

	pos := startPos
	dir := Position{-1, 0}

	for {
		visitedKey := LoopVisitedKey{row: pos.row, col: pos.col, dirRow: dir.row, dirCol: dir.col}
		if visited[visitedKey] > 1 {
			return true
		}

		visited[visitedKey]++

		next := pos.add(&dir)

		found, nextChar := grid.get(next)

		if found {
			if nextChar == "#" {
				dir = dir.turnRight()
			} else {
				pos = next
			}
		} else {
			break
		}
	}

	return false
}

func getVisited(grid *Grid, startPos Position) map[VisitedKey]Position {
	visited := make(map[VisitedKey]Position)

	pos := startPos
	dir := Position{-1, 0}

	for {
		visitedKey := VisitedKey{row: pos.row, col: pos.col}
		visited[visitedKey] = pos

		next := pos.add(&dir)

		found, nextChar := grid.get(next)

		if found {
			if nextChar == "#" {
				dir = dir.turnRight()
			} else {
				pos = next
			}
		} else {
			break
		}
	}

	return visited
}
