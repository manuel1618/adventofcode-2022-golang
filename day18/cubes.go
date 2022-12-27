package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := readLines("./data/cubes.txt")

	maxX, maxY, maxZ := getGridSize(lines)

	grid3D := make([][][]int, maxX+1)
	for i := range grid3D {
		grid3D[i] = make([][]int, maxY+1)
		for j := range grid3D[i] {
			grid3D[i][j] = make([]int, maxZ+1)
		}
	}
	grid3D = fillGrid3D(grid3D, lines)
	printGrid(grid3D)
	surfaceArea := getSurfaceArea(grid3D)
	fmt.Println(surfaceArea)
}

func printGrid(grid3D [][][]int) {
	for x := range grid3D {
		for y := range grid3D[x] {
			for z := range grid3D[x][y] {
				fmt.Print(grid3D[x][y][z])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func getSurfaceArea(grid3D [][][]int) int {
	surfaceArea := 0
	for x := range grid3D {
		for y := range grid3D[x] {
			for z := range grid3D[x][y] {
				if grid3D[x][y][z] == 1 {
					surfaceArea += getSurfaceAreaOfCube(grid3D, x, y, z)
				}
			}
		}
	}
	return surfaceArea
}

func getSurfaceAreaOfCube(grid3D [][][]int, x, y, z int) int {
	surfaceArea := 0
	if x == 0 || grid3D[x-1][y][z] == 0 {
		surfaceArea++
	}
	if x == len(grid3D)-1 || grid3D[x+1][y][z] == 0 {
		surfaceArea++
	}
	if y == 0 || grid3D[x][y-1][z] == 0 {
		surfaceArea++
	}
	if y == len(grid3D[x])-1 || grid3D[x][y+1][z] == 0 {
		surfaceArea++
	}
	if z == 0 || grid3D[x][y][z-1] == 0 {
		surfaceArea++
	}
	if z == len(grid3D[x][y])-1 || grid3D[x][y][z+1] == 0 {
		surfaceArea++
	}
	return surfaceArea
}

func fillGrid3D(grid3D [][][]int, lines []string) [][][]int {
	for _, line := range lines {
		x, y, z := getCoords(line)
		grid3D[x][y][z] = 1
	}
	return grid3D
}

func getGridSize(lines []string) (int, int, int) {
	maxX := 0
	maxY := 0
	maxZ := 0
	for _, line := range lines {
		x, y, z := getCoords(line)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if z > maxZ {
			maxZ = z
		}
	}
	return maxX, maxY, maxZ
}

func getCoords(line string) (int, int, int) {
	var x, y, z int
	_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
	if err != nil {
		log.Fatal(err)
	}
	return x, y, z
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
