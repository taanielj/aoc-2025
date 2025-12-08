package day04

import (
	"aoc2025/internal/utils"
	"fmt"
)

type shelfGrid struct {
	height int
	width  int
	cells  [][]bool
}

func SolveDay04() {
	allowedNeighbors := 4
	lines, err := utils.ReadLinesFromFile("inputs/day04.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		panic(err)
	}

	grid := getGridFromLines(lines)
	fmt.Println("Day 04 Solution:")

	rollsOfPaper := 0
	for {
		collected := gatherRolls(grid, allowedNeighbors)
		if collected == 0 {
			break
		}
		rollsOfPaper += collected
	}
	fmt.Printf("Number of rolls of paper needed: %d\n", rollsOfPaper)
}

func getGridFromLines(lines []string) shelfGrid {
	height := len(lines)
	width := len(lines[0])
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
		for j, char := range lines[i] {
			if char == '@' {
				cells[i][j] = true
			} else {
				cells[i][j] = false
			}
		}
	}
	return shelfGrid{
		height: height,
		width:  width,
		cells:  cells,
	}
}

func gatherRolls(grid shelfGrid, allowedNeighbors int) int {
	rollsOfPaper := 0

	for i := 0; i < grid.height; i++ {
		for j := 0; j < grid.width; j++ {
			if grid.cells[i][j] {
				neighbors := countNeighbors(grid, i, j)
				if neighbors < allowedNeighbors {
					rollsOfPaper++
					grid.cells[i][j] = false
				}

			}
		}
	}
	fmt.Printf(" Rolls collected this round: %d\n", rollsOfPaper)
	return rollsOfPaper
}

func countNeighbors(grid shelfGrid, row, col int) int {
	neighbors := 0
	directions := []struct{ dRow, dCol int }{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}

	for _, dir := range directions {
		neighborRow := row + dir.dRow
		neighborCol := col + dir.dCol
		if neighborRow >= 0 && neighborRow < grid.height &&
			neighborCol >= 0 && neighborCol < grid.width &&
			grid.cells[neighborRow][neighborCol] {
			neighbors++
		}
	}

	return neighbors
}
