package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	cwd, _ := os.Getwd()
	path := (cwd + "/data/sections.txt")
	lines := readLines(path)

	// Part 1 - one set of sections must comepletely contain the other
	points := 0
	for _, line := range lines {
		sectionsElf1 := strings.Split(line, ",")[0]
		minSectionElf1, err := strconv.Atoi(strings.Split(sectionsElf1, "-")[0])
		maxSectionElf1, err := strconv.Atoi(strings.Split(sectionsElf1, "-")[1])
		sectionsElf2 := strings.Split(line, ",")[1]
		minSectionElf2, err := strconv.Atoi(strings.Split(sectionsElf2, "-")[0])
		maxSectionElf2, err := strconv.Atoi(strings.Split(sectionsElf2, "-")[1])

		if err != nil {
			fmt.Println("Error during conversion")
			return
		}

		if minSectionElf1 >= minSectionElf2 && maxSectionElf1 <= maxSectionElf2 {
			points++
		} else if minSectionElf1 <= minSectionElf2 && maxSectionElf1 >= maxSectionElf2 {
			points++
		}

	}
	fmt.Printf("Total Score (1st part): %d\n", points)

	// Part 2 - overlap at all
	points = 0
	for _, line := range lines {
		sectionsElf1 := strings.Split(line, ",")[0]
		minSectionElf1, err := strconv.Atoi(strings.Split(sectionsElf1, "-")[0])
		maxSectionElf1, err := strconv.Atoi(strings.Split(sectionsElf1, "-")[1])
		sectionsElf2 := strings.Split(line, ",")[1]
		minSectionElf2, err := strconv.Atoi(strings.Split(sectionsElf2, "-")[0])
		maxSectionElf2, err := strconv.Atoi(strings.Split(sectionsElf2, "-")[1])

		if err != nil {
			fmt.Println("Error during conversion")
			return
		}

		// This implementation is not the best, but it works
		if minSectionElf1 <= maxSectionElf2 && minSectionElf1 >= minSectionElf2 {
			points++
		} else if maxSectionElf1 >= minSectionElf2 && maxSectionElf1 <= maxSectionElf2 {
			points++
		} else if minSectionElf2 <= maxSectionElf1 && minSectionElf2 >= minSectionElf1 {
			points++
		} else if maxSectionElf2 >= minSectionElf1 && maxSectionElf2 <= maxSectionElf1 {
			points++
		}
	}
	fmt.Printf("Total Score (2nd part): %d", points)

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
