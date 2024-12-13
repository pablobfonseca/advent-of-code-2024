package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	rowLen int
	colLen int
	data   [][]rune
}

func NewGrid(data [][]rune) *Grid {
	return &Grid{
		rowLen: len(data),
		colLen: len(data[0]),
		data:   data,
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
	for row := 0; row < g.rowLen; row++ {
		for col := 0; col < len(g.data[0]); col++ {
			if char == g.data[row][col] {
				pos = Position{
					row: row,
					col: col,
				}

				break
			}
		}
	}

	return pos
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

	visited := make(map[string]int)

	pos := grid.find('^')
	dir := Position{-1, 0}

	for {
		visitedKey := fmt.Sprintf("%d|%d", pos.row, pos.col)
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

	fmt.Printf("Total: %d\n", len(visited))
}
