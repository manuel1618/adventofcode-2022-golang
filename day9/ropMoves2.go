package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Coords struct {
	X int
	Y int
}

func (tail *Coords) MoveTo(head *Coords) {
	// If the head is right next to or directly on top of the tail then we do nothing
	diffInX := math.Abs(float64(tail.X - head.X))
	diffInY := math.Abs(float64(tail.Y - head.Y))
	if diffInY <= 1 && diffInX <= 1 {
		return
	}
	//fmt.Println(fmt.Sprintf("Tail moving (%s) to meet head (%s)", tail.String(), head.String()))
	xDirection := 1
	yDirection := 1
	if head.Y < tail.Y {
		yDirection = -1
	}
	if head.X < tail.X {
		xDirection = -1
	}
	if head.X == tail.X {
		tail.Y += 1 * yDirection
		return
	}
	if head.Y == tail.Y {
		tail.X += 1 * xDirection
		return
	}

	// Diagonal
	tail.X += 1 * xDirection
	tail.Y += 1 * yDirection
}

func (c *Coords) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func main() {
	input := readLines("./data/moves.txt")
	coords := make([]Coords, 10)
	longTailPositions := map[string]bool{}
	shortTailPositions := map[string]bool{}
	longTailPositions[coords[9].String()] = true
	shortTailPositions[coords[1].String()] = true
	for _, cmd := range input {
		var direction string
		var amount int
		fmt.Sscanf(cmd, "%s %d", &direction, &amount)
		for x := 0; x < amount; x++ {
			switch direction {
			case "R":
				coords[0].X++
			case "U":
				coords[0].Y++
			case "D":
				coords[0].Y--
			case "L":
				coords[0].X--
			}
			for i := 1; i < len(coords); i++ {
				coords[i].MoveTo(&coords[i-1])
			}
			longTailPositions[coords[9].String()] = true
			shortTailPositions[coords[1].String()] = true
		}
	}
	fmt.Println(fmt.Sprintf("Part 1: %d", len(shortTailPositions)))
	fmt.Println(fmt.Sprintf("Part 2: %d", len(longTailPositions)))
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
