package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := (cwd + "/data/rocks.txt")
	lines := readLines(path)

	// Part 1
	// Determine the grid size
	gridSize := getGridSize(lines)

	// Create a grid of the size determined above
	grid := make([][]int, gridSize[0])
	// fmt.Printf("Grid size: %v\n", gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize[1])
	}
	buildRocks(grid, lines)

	sandCounter := 1
	newLocation := sandDropRestLocation(grid, [2]int{500, 0})

	for newLocation[1] < len(grid[0])-1 {
		sandCounter++
		grid[newLocation[0]][newLocation[1]] = 2
		newLocation = sandDropRestLocation(grid, [2]int{500, 0})
	}
	fmt.Printf("Sand counter at which it falls into the void (Pt1): %v\n", sandCounter-1)
	// printGrid(grid)

	// Part 2
	removeSand(grid)
	// Create a grid of the size determined above
	grid = make([][]int, gridSize[0]*2)
	// fmt.Printf("Grid size: %v\n", gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize[1]+2)
	}
	buildRocks(grid, lines)
	for i := 0; i < len(grid); i++ {
		grid[i][len(grid[0])-1] = 1
	}
	sandCounter = 1

	newLocation = sandDropRestLocation(grid, [2]int{500, 0})
	for true {
		sandCounter++
		grid[newLocation[0]][newLocation[1]] = 2
		newLocation = sandDropRestLocation(grid, [2]int{500, 0})
		if newLocation[0] == 500 && newLocation[1] == 0 {
			break
		}
	}
	fmt.Printf("Sand counter at which the entrance is blocked (Pt2): %v\n", sandCounter)

}

func removeSand(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				grid[i][j] = 0
			}
		}
	}
}

func buildRocks(grid [][]int, lines []string) {
	for _, line := range lines {
		points := strings.Split(line, "->")
		for i := 0; i < len(points)-1; i++ {
			point := points[i]
			nextPoint := points[i+1]
			x1, err := strconv.Atoi(strings.TrimSpace(strings.Split(point, ",")[0]))
			y1, err := strconv.Atoi(strings.TrimSpace(strings.Split(point, ",")[1]))
			x2, err := strconv.Atoi(strings.TrimSpace(strings.Split(nextPoint, ",")[0]))
			y2, err := strconv.Atoi(strings.TrimSpace(strings.Split(nextPoint, ",")[1]))

			if err != nil {
				log.Fatal(err)
			}
			for i := min(x1, x2); i <= max(x1, x2); i++ {
				for j := min(y1, y2); j <= max(y1, y2); j++ {
					grid[i][j] = 1
				}
			}
		}
	}
}

func sandDropRestLocation(grid [][]int, start [2]int) [2]int {
	// limit check
	if start[1] >= len(grid[0])-1 {
		return start
	}
	// fall down
	if grid[start[0]][start[1]+1] == 0 {
		return sandDropRestLocation(grid, [2]int{start[0], start[1] + 1})
	} else {
		// fall diag left
		if grid[start[0]-1][start[1]+1] == 0 {
			return sandDropRestLocation(grid, [2]int{start[0] - 1, start[1] + 1})
		} else {
			// fall diag right
			if grid[start[0]+1][start[1]+1] == 0 {
				return sandDropRestLocation(grid, [2]int{start[0] + 1, start[1] + 1})
			} else {
				return start
			}
		}
	}
}

func checkIfSandRundThrough(grid [][]int, x, y int) bool {
	// lastline must contain a 2
	lastline := grid[len(grid)-1]
	for i := range lastline {
		if lastline[i] == 2 {
			return true
		}
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getGridSize(lines []string) [2]int {
	maxHeight := -math.MaxInt32
	maxWidth := -math.MaxInt32
	for _, line := range lines {
		points := strings.Split(line, "->")
		for _, point := range points {
			x, err := strconv.Atoi(strings.TrimSpace(strings.Split(point, ",")[0]))
			y, err := strconv.Atoi(strings.TrimSpace(strings.Split(point, ",")[1]))
			if err != nil {
				log.Fatal(err)
			}
			if x > maxWidth {
				maxWidth = x
			}
			if y > maxHeight {
				maxHeight = y
			}
		}
	}
	return ([2]int{maxWidth + 1, maxHeight + 1})
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid[0]); i++ {
		for j := range grid {
			if grid[j][i] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%v", grid[j][i])
			}
		}
		fmt.Println()
	}
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
