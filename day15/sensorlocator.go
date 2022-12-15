package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Sensor struct {
	x           int
	y           int
	searchRange int
}

type Beacon struct {
	x int
	y int
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := (cwd + "/data/sensors.txt")

	nonPresentRow := 0
	nonPresentRow = 10      // testdata
	nonPresentRow = 2000000 // data
	lines := readLines(path)

	gridSize := getGridSize(lines)
	gridWidth := gridSize[1] - gridSize[0]

	sensors, beacons := getSensorsAndBeacons(lines)
	lineOfInterst := make([]int, gridWidth-gridSize[0])
	blocks := 0
	for i := range lineOfInterst {
		lineOfInterst[i] = 0
		x := i + gridSize[0]
		// check if a sensor blocks this field
		for _, sensor := range sensors {
			distance := vecDistManhatten([2]int{x, nonPresentRow}, [2]int{sensor.x, sensor.y})
			if distance <= sensor.searchRange {
				lineOfInterst[i] = 3
				blocks++
				break
			}
		}
		for _, beacon := range beacons {
			if x == beacon.x && nonPresentRow == beacon.y {
				lineOfInterst[i] = 2
				blocks--
				fmt.Print("beacon found\n")
				break
			}
		}

	}
	fmt.Printf("Number of Blocks: %v\n", blocks)

	// part2
	// Determine the beacon location - add diagonal elements around the sensor range
	var possiblePoints [][]int
	for _, sensor := range sensors {
		searchrange := sensor.searchRange + 1
		for i := 0; i < searchrange; i++ {
			inv := searchrange - i
			possiblePoints = append(possiblePoints, []int{sensor.x + i, sensor.y + inv})
			possiblePoints = append(possiblePoints, []int{sensor.x + i, sensor.y - inv})
			possiblePoints = append(possiblePoints, []int{sensor.x - i, sensor.y + inv})
			possiblePoints = append(possiblePoints, []int{sensor.x - i, sensor.y - inv})
		}
	}

	lower := 0
	upper := 4000000
	for i := range possiblePoints {
		blocked := false
		x := possiblePoints[i][0]
		y := possiblePoints[i][1]
		if x < lower || y < lower {
			continue
		}
		if x > upper || y > upper {
			continue
		}
		for _, sensor := range sensors {
			distance := vecDistManhatten([2]int{x, y}, [2]int{sensor.x, sensor.y})
			if distance <= sensor.searchRange {
				blocked = true
				break
			}
		}
		if !blocked {
			// Unblocked point found - should be only one so its okay to break here
			freq := possiblePoints[i][0]*4000000 + possiblePoints[i][1]
			fmt.Printf("Possible point: %v\n Frequency: %v\n", possiblePoints[i], freq)
			break
		}
	}
}

func getSensorsAndBeacons(lines []string) ([]Sensor, []Beacon) {

	var sensors []Sensor
	var beacons []Beacon

	// Plant sensors
	for _, line := range lines {
		fmt.Printf("Line: %v\n", line)
		point1x, err := strconv.Atoi(strings.Split(strings.Split(line, "x=")[1], ",")[0])
		point1y, err := strconv.Atoi(strings.Split(strings.Split(line, "y=")[1], ":")[0])
		point2x, err := strconv.Atoi(strings.Split(strings.Split(line, "x=")[2], ",")[0])
		point2y, err := strconv.Atoi(strings.TrimSpace(strings.Split(line, "y=")[2]))
		if err != nil {
			log.Fatal(err)
		}
		// Fill the grid with the sensors
		sensor := [2]int{point1x, point1y}
		beacon := [2]int{point2x, point2y}
		distance := vecDistManhatten(sensor, beacon)
		fmt.Printf("Sensor added, at %v,%v with range %v\n", point1x, point1y, distance)
		sensors = append(sensors, Sensor{point1x, point1y, distance})
		beacons = append(beacons, Beacon{point2x, point2y})
	}
	return sensors, beacons
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func vecDistManhatten(a [2]int, b [2]int) int {
	return (Abs(a[0] - b[0])) + (Abs(a[1] - b[1]))
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid[0]); i++ {
		for j := range grid {
			if grid[j][i] == 0 {
				fmt.Printf(".")
			} else if grid[j][i] == 1 {
				fmt.Printf("S")
			} else if grid[j][i] == 2 {
				fmt.Printf("B")
			} else if grid[j][i] == 3 {
				fmt.Printf("#")
			} else {
				fmt.Printf("%v", grid[j][i])
			}
		}
		fmt.Println()
	}
}

func getGridSize(lines []string) [4]int {
	maxHeight := 0
	minHeight := 0
	maxWidth := 0
	minWidth := 0
	for _, line := range lines {
		point1x, err := strconv.Atoi(strings.Split(strings.Split(line, "x=")[1], ",")[0])
		point1y, err := strconv.Atoi(strings.Split(strings.Split(line, "y=")[1], ":")[0])
		point2x, err := strconv.Atoi(strings.Split(strings.Split(line, "x=")[2], ",")[0])
		point2y, err := strconv.Atoi(strings.TrimSpace(strings.Split(line, "y=")[2]))

		sensor := [2]int{point1x, point1y}
		beacon := [2]int{point2x, point2y}
		distance := vecDistManhatten(sensor, beacon)

		if err != nil {
			log.Fatal(err)
		}
		if point1x+distance > maxWidth {
			maxWidth = point1x + distance
		}
		if point1x-distance < minWidth {
			minWidth = point1x - distance
		}
		if point1y+distance > maxHeight {
			maxHeight = point1y + distance
		}
		if point1y-distance < minHeight {
			minHeight = point1y - distance
		}

	}
	return ([4]int{minWidth, maxWidth, minHeight, maxHeight})
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
