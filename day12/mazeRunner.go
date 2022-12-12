package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x       int
	y       int
	height  string
	parrent *Point // For BFS
	visited bool   // For BFS
}

type Queue struct {
	items []*Point
}

func (q *Queue) enqueue(i *Point) {
	q.items = append(q.items, i)
}

func (q *Queue) dequeue() *Point {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func main() {
	path := ("./data/heightMapSmall.txt")
	lines := readLines(path)
	grid, start, end := buildGrid(lines)

	// Part 1
	// Find shortest path from start to the end point
	resetPoints(grid)
	steps := findShortestPath(grid, start, end)
	fmt.Println("Part 1: ", steps)

	// Part 2
	// Find shortest path from any point with height a to the end point
	startPoints := findAllPointsInGrid(grid, "a")
	minNumberOfStepsPt2 := steps
	for _, start := range startPoints {
		resetPoints(grid)
		steps := findShortestPath(grid, start, end)
		if steps < minNumberOfStepsPt2 && steps != -1 {
			minNumberOfStepsPt2 = steps
		}
	}
	fmt.Println("Part 2: ", minNumberOfStepsPt2)

}

func resetPoints(grid [][]*Point) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j].visited = false
			grid[i][j].parrent = nil
		}
	}
}

func findAllPointsInGrid(grid [][]*Point, height string) []*Point {
	var points []*Point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].height == height {
				points = append(points, grid[i][j])
			}
		}
	}
	return points
}

func findShortestPath(grid [][]*Point, start *Point, end *Point) int {
	// Use BFS to find the shortest path
	q := Queue{}
	q.enqueue(start)
	start.visited = true
	for !q.isEmpty() {
		mp := q.dequeue()
		if mp.x == end.x && mp.y == end.y {
			steps := 0
			var point *Point = mp
			for point.parrent != nil {
				steps++
				point = point.parrent
			}
			return steps
		}
		for _, p := range getPossibleMoves(grid, mp) {
			if !p.visited {
				p.visited = true
				p.parrent = mp
				q.enqueue(p)
			}
		}
	}
	return -1
}

func getPossibleMoves(grid [][]*Point, point *Point) []*Point {
	neigbours := [][]int{{point.x - 1, point.y}, {point.x + 1, point.y}, {point.x, point.y - 1}, {point.x, point.y + 1}}
	var possibleMoves []*Point
	for _, cood := range neigbours {
		x := cood[0]
		y := cood[1]
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			// boundary check
			continue
		} else {
			if calculateAscencense(point, grid[x][y]) > 1 {
				// Not possible to move more than one up in height (a is lowest height, z is highest)
				continue
			}
			possibleMoves = append(possibleMoves, grid[x][y])
		}
	}
	return possibleMoves
}

func buildGrid(lines []string) ([][]*Point, *Point, *Point) {
	grid := make([][]*Point, len(lines))
	for i := range grid {
		grid[i] = make([]*Point, len(lines[i]))
	}
	var start *Point
	var end *Point
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			height := string(lines[i][j])
			grid[i][j] = &Point{i, j, height, nil, false}
			if height == "S" {
				start = grid[i][j]
				start.height = "a"
			}
			if height == "E" {
				end = grid[i][j]
				end.height = "z"
			}
		}
	}
	return grid, start, end
}

// read a file and return the lines as an array
func readLines(path string) []string {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initializing
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Scon the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func calculateAscencense(a *Point, b *Point) int {
	heights := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	return indexOf(heights, b.height) - indexOf(heights, a.height)
}

func indexOf(arr []string, val string) int {
	for pos, v := range arr {
		if v == val {
			return pos
		}
	}
	return -1
}
