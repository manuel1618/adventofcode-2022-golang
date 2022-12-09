package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := (cwd + "/data/moves.txt")
	lines := readLines(path)

	// Part 1
	// Determine the grid size
	gridSize := getGridSize(lines)
	start := [2]int{Abs(gridSize[0]), Abs(gridSize[2])}
	gridWidth := gridSize[1] - gridSize[0]
	gridHeight := gridSize[3] - gridSize[2]

	// Create a grid of the size determined above
	grid := make([][]int, gridWidth)
	// fmt.Printf("Grid size: %v\n", gridSize)
	for i := range grid {
		grid[i] = make([]int, gridHeight)
	}

	// Rope
	ropeLength := 2 // part1
	//ropeLength := 10 // part2
	rope := make([][2]int, ropeLength)
	for i := range rope {
		rope[i] = start
	}

	// Fill the grid with the moves
	for _, line := range lines {
		direction := string(line[0])
		distance, err := strconv.Atoi(line[2:])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < distance; i++ {
			// Determine the new head position
			rope = updateRopePosition(rope, direction, grid)
			// Fill the grid
			tailGrid := rope[ropeLength-1]
			grid[tailGrid[0]][tailGrid[1]] = 1
			//printRope(rope, grid)
		}
	}
	// get the number of 1s in the grid
	numberOfOnes := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				numberOfOnes++
			}
		}
	}
	// fmt.Println("Final grid:")
	// printGrid(grid)
	fmt.Printf("Number of 1s: %v\n", numberOfOnes)
}

// Helper Functions
func printRope(rope [][2]int, grid [][]int) {
	// make a local copy of the grid
	localGrid := make([][]int, len(grid))
	for i := range grid {
		localGrid[i] = make([]int, len(grid[i]))
	}

	for i := range rope {
		localGrid[rope[i][0]][rope[i][1]] = i + 1
	}
	printGrid(localGrid)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func vecDist(a [2]int, b [2]int) float64 {
	x := float64(Abs(a[0] - b[0]))
	y := float64(Abs(a[1] - b[1]))
	distance := math.Sqrt(x*x + y*y)
	return distance
}

func printGrid(grid [][]int) {
	for i := len(grid[0]) - 1; i >= 0; i-- {
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

func determineNewHeadPosition(headPosition [2]int, direction string) [2]int {
	switch direction {
	case "U":
		return [2]int{headPosition[0], headPosition[1] + 1}
	case "D":
		return [2]int{headPosition[0], headPosition[1] - 1}
	case "R":
		return [2]int{headPosition[0] + 1, headPosition[1]}
	case "L":
		return [2]int{headPosition[0] - 1, headPosition[1]}
	}
	return [2]int{0, 0}
}

func updateRopePosition(oldRope [][2]int, direction string, grid [][]int) [][2]int {

	// Create a new rope
	newRope := make([][2]int, len(oldRope))

	// Determine the new head position
	newHead := determineNewHeadPosition(oldRope[0], direction)
	newRope[0] = newHead
	for i := 1; i < len(oldRope); i++ {
		newRope[i] = oldRope[i]
	}

	for j := 0; j < len(newRope)-1; j++ {
		distance := vecDist(newRope[j], newRope[j+1])
		if distance >= 2 {
			widthMovement := newRope[j][0] - newRope[j+1][0]
			heightMovement := newRope[j][1] - newRope[j+1][1]

			// diagonal movement
			if Abs(widthMovement) == Abs(heightMovement) {
				if widthMovement > 0 && heightMovement > 0 {
					// up right
					newRope[j+1] = [2]int{newRope[j][0] - 1, newRope[j][1] - 1}
				} else if widthMovement > 0 && heightMovement < 0 {
					// down right
					newRope[j+1] = [2]int{newRope[j][0] - 1, newRope[j][1] + 1}
				} else if widthMovement < 0 && heightMovement > 0 {
					// up left
					newRope[j+1] = [2]int{newRope[j][0] + 1, newRope[j][1] - 1}
				} else if widthMovement < 0 && heightMovement < 0 {
					// down left
					newRope[j+1] = [2]int{newRope[j][0] + 1, newRope[j][1] + 1}
				}

			} else {
				// horizontal or vertical movement
				if widthMovement == 0 {
					if heightMovement > 0 {
						// up
						newRope[j+1] = [2]int{newRope[j][0], newRope[j][1] - 1}
					} else {
						// down
						newRope[j+1] = [2]int{newRope[j][0], newRope[j][1] + 1}
					}
				} else if heightMovement == 0 {
					if widthMovement > 0 {
						// right
						newRope[j+1] = [2]int{newRope[j][0] - 1, newRope[j][1]}
					} else {
						// left
						newRope[j+1] = [2]int{newRope[j][0] + 1, newRope[j][1]}
					}
				} else if Abs(widthMovement) > Abs(heightMovement) {
					// special movements
					if widthMovement > 0 {
						// right
						newRope[j+1] = [2]int{newRope[j][0] - 1, newRope[j][1]}
					} else {
						// left
						newRope[j+1] = [2]int{newRope[j][0] + 1, newRope[j][1]}
					}
				} else if Abs(widthMovement) < Abs(heightMovement) {
					if heightMovement > 0 {
						// up
						newRope[j+1] = [2]int{newRope[j][0], newRope[j][1] - 1}
					} else {
						// down
						newRope[j+1] = [2]int{newRope[j][0], newRope[j][1] + 1}
					}
				} else {
					fmt.Println("Error")
				}
			}
		}

		// fmt.Println("")
		// fmt.Println("Zwischenschritt")
		// printRope(newRope, grid)
	}
	if vecDist(newRope[len(newRope)-1], oldRope[len(newRope)-1]) > 2 {
		// fmt.Printf("Old rope: %v\n", oldRope)
		// fmt.Printf("New rope: %v\n", newRope)
		// printRope(oldRope, grid)
		// fmt.Println("New Rope")
		// printRope(newRope, grid)
		log.Fatal("Error")
	}
	return newRope
}

func getGridSize(lines []string) [4]int {

	maxHeight := 0
	minHeight := 0
	maxWidth := 0
	minWidth := 0
	width := 0
	height := 0
	for _, line := range lines {
		direciton := line[0]
		distance, err := strconv.Atoi(line[2:])
		if err != nil {
			log.Fatal(err)
		}
		switch direciton {
		case 'U':
			height += distance
			if height > maxHeight {
				maxHeight = height
			}
		case 'D':
			height -= distance
			if height < minHeight {
				minHeight = height
			}
		case 'R':
			width += distance
			if width > maxWidth {
				maxWidth = width
			}
		case 'L':
			width -= distance
			if width < minWidth {
				minWidth = width
			}
		}
	}
	return ([4]int{minWidth, maxWidth + 1, minHeight, maxHeight + 1})
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
