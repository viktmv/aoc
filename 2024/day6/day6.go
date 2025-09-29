package main

import (
	"bufio"
	"os"

	"log"
)

type Point struct {
	x       int
	y       int
	val     string
	visited bool
}

type Grid = [][]*Point

func newGrid(lines []string) Grid {
	var grid [][]*Point

	for y, line := range lines {
		var row []*Point
		for x, char := range line {
			point := &Point{x: x, y: y, val: string(char), visited: false}
			row = append(row, point)
		}
		grid = append(grid, row)
	}

	return grid
}

func main() {
	file, err := os.Open("input_test.txt")

	if err != nil {
		panic("file not found?")
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

    grid := newGrid(lines)
    
    for y := range(grid) {
        for x := range(grid[y]) {
            point := grid[y][x]
        }
    }
}
