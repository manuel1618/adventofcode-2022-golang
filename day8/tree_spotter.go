package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	cwd,err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := (cwd + "/data/trees.txt")
	lines := readLines(path)

	visibleTrees := 0
	maxScenicScore := 0
	for i := 1; i < len(lines)-1; i++ {	
		// iterate rows
		for j := 1; j < len(lines[0])-1; j++ {
			// iterate columns
			tree := lines[i][j]
			
			// Part 1
			visibleUp := true
			visibleDown := true
			visibleLeft := true
			visibleRight := true

			// Part 2
			viewingDistanceUp := 0
			viewingDistanceDown := 0
			viewingDistanceLeft := 0
			viewingDistanceRight := 0

			// look up 
			for k := i-1; k >= 0; k-- { 
				viewingDistanceUp++
				if lines[k][j] >= tree  {
					visibleUp = false
					break
				}
			}
			// look down
			for k := i+1; k < len(lines); k++ {
				viewingDistanceDown++
				if lines[k][j] >= tree  {
					visibleDown = false
					break
				}
			}
			// look left
			for k := j-1; k >= 0; k-- {
				viewingDistanceLeft++
				if lines[i][k] >= tree  {
					visibleLeft = false
					break
				}
			}
			// look right
			for k := j+1; k < len(lines[0]); k++ { 
				viewingDistanceRight++
				if lines[i][k] >= tree  {
					visibleRight = false
					break
				}
			}
			visible := visibleUp || visibleDown || visibleLeft || visibleRight 
			if visible {
				visibleTrees++
			}
			// calculate scenic score
			scenicScore := viewingDistanceUp * viewingDistanceDown * viewingDistanceLeft * viewingDistanceRight
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	} 
		
	// Part 1
	edgeTrees :=2*len(lines[0]) + 2*len(lines) - 4
	visibleTrees += edgeTrees
	fmt.Printf("Number of visible trees: %d\n", visibleTrees)
		
	// Part 2
	fmt.Printf("Max Scenic Score: %d\n", maxScenicScore)
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