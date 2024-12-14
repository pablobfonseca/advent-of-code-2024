package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

type Grid struct {
	rowLen    int
	colLen    int
	data      [][]rune
	positions map[rune][]Point
}

func NewGrid(data [][]rune) *Grid {
	positions := make(map[rune][]Point)

	for row := range data {
		for col, char := range data[row] {
			if char == '.' {
				continue
			}
			positions[char] = append(positions[char], Point{x: col, y: row})
		}
	}

	return &Grid{
		rowLen:    len(data),
		colLen:    len(data[0]),
		data:      data,
		positions: positions,
	}
}
func (g *Grid) isInBounds(p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.y < g.rowLen && p.x < g.colLen
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	var gridData [][]rune

	for scanner.Scan() {
		line := scanner.Text()

		gridData = append(gridData, []rune(line))
	}

	grid := NewGrid(gridData)

	result := findAntinodes(grid)

	fmt.Printf("Total: %d", result)

}

func findAntinodes(grid *Grid) int {
	antinodeSet := make(map[Point]struct{})

	for _, points := range grid.positions {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1, p2 := points[i], points[j]

				dx, dy := p2.x-p1.x, p2.y-p1.y

				antinode1 := Point{x: p1.x - dx, y: p1.y - dy}
				antinode2 := Point{x: p2.x + dx, y: p2.y + dy}

				if grid.isInBounds(antinode1) {
					antinodeSet[antinode1] = struct{}{}
				}
				if grid.isInBounds(antinode2) {
					antinodeSet[antinode2] = struct{}{}
				}
			}
		}
	}

	return len(antinodeSet)
}
