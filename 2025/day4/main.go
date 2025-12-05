package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Grid = [][]string

func main() {
	var grid Grid

	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("could not open file for reading %s", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid = append(grid, line)
	}

	accessible := 0
	// pt1 
	counter := countAccessible(grid)

	// p2 
	for counter != 0 {
		accessible += counter
		counter = countAccessible(grid)
	}

	log.Printf("accessible %d", accessible)
}

func prettyPrintGrid(grid Grid) {
	for row := range grid {
		fmt.Printf("|")
		for col := range grid[row] {
			fmt.Printf(" %s ", grid[row][col])
		}
		fmt.Printf("|")
		fmt.Printf("\n")
	}
}

func countAccessible(grid Grid) int {
	accessible := 0
	for row := range grid {
		for col := range grid[row] {
			roll := grid[row][col]
			if roll != "@" {
				continue
			}

			occupied := 0
			// top-left
			if row > 0 && col > 0 {
				if grid[row-1][col-1] == "@" {
					occupied++
				}
			}
			// top center
			if row > 0 {
				if grid[row-1][col] == "@" {
					occupied++
				}
			}
			// top right
			if row > 0 && col < len(grid[row])-1 {
				if grid[row-1][col+1] == "@" {
					occupied++
				}
			}
			// right
			if col < len(grid[row])-1 {
				if grid[row][col+1] == "@" {
					occupied++
				}
			}
			// bottom right
			if row < len(grid)-1 && col < len(grid[row])-1 {
				if grid[row+1][col+1] == "@" {
					occupied++
				}
			}
			// bottom
			if row < len(grid)-1 {
				if grid[row+1][col] == "@" {
					occupied++
				}
			}
			// bottom-left
			if row < len(grid)-1 && col > 0 {
				if grid[row+1][col-1] == "@" {
					occupied++
				}
			}
			// left
			if col > 0 {
				if grid[row][col-1] == "@" {
					occupied++
				}
			}

			if occupied < 4 {
				grid[row][col] = "x"
				accessible++
			}
		}
	}

	return accessible
}
