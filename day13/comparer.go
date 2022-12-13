package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	path := ("./data/values.txt")
	lines := readLines(path)

	// Part 1
	// Find shortest path from start to the end point
	sumIndices := 0
	pair := 1
	for i := 0; i < len(lines)-1; i++ {
		fmt.Printf("----->    Comparing new Lists \nleft: %v, \nright:%v\n", lines[i], lines[i+1])
		fmt.Println("")
		if compareTwoStrings(lines[i], lines[i+1]) == 1 {
			sumIndices += pair
		}
		pair += 1
		i += 2
	}
	fmt.Println("Part 1: ", sumIndices)

	// Part 2
	var linesToSort []string
	for i := 0; i < len(lines); i++ {
		// ignore empty lines
		if lines[i] != "" {
			linesToSort = append(linesToSort, lines[i])
		}
	}
	linesToSort = append(linesToSort, "[[2]]")
	linesToSort = append(linesToSort, "[[6]]")
	sort.Slice(linesToSort, func(i, j int) bool {
		return compareTwoItems(linesToSort[i], linesToSort[j]) > 0
	})
	a := 0
	b := 0
	for i := 0; i < len(linesToSort); i++ {
		if linesToSort[i] == "[[2]]" {
			a = i + 1
		}
		if linesToSort[i] == "[[6]]" {
			b = i + 1
		}
	}
	fmt.Println("Part 2: ", a*b)

}

func compareTwoStrings(left string, right string) int {

	leftItems := parseLine(left)
	rightItems := parseLine(right)
	compareResult := 0

	for i := 0; i < len(leftItems); i++ {
		if i >= len(rightItems) {
			fmt.Println("Right side run out of items")
			return -1
		}
		compareResult = compareTwoItems(leftItems[i], rightItems[i])
		if compareResult != 0 {
			return compareResult
		}
		if compareResult == 1 {
			fmt.Println("Left side is smaller")
			return 1
		}
		if compareResult == -1 {
			fmt.Println("Right side is smaller")
			return -1
		}
	}
	if len(leftItems) < len(rightItems) {
		fmt.Println("Left side is out of items")
		return 1
	}
	return 0
}

func compareTwoItems(left string, right string) int {
	fmt.Printf("Item: %s and %s\n", left, right)
	if left == "" && right != "" {
		return 1
	}
	if right == "" && left != "" {
		return -1
	}
	if left == "" && right == "" {
		return 0
	}

	if left[0] == '[' {
		if right[0] == '[' {
			// both are arrays
			return compareTwoStrings(left, right)
		} else {
			// left is array, right is not
			return compareTwoStrings(left, "["+right+"]")
		}
	} else {
		if right[0] == '[' {
			// left is not array, right is
			return compareTwoStrings("["+left+"]", right)
		} else {
			// both are not arrays
			leftItems := strings.Split(left, ",")
			rightItems := strings.Split(right, ",")
			for i := 0; i < len(leftItems); i++ {
				if i >= len(rightItems) {
					fmt.Println("Right side run out of items")
					return -1
				}
				l, err := strconv.Atoi(leftItems[i])
				r, err := strconv.Atoi(rightItems[i])
				if err != nil {
					log.Fatal(err)
				}
				if l < r {
					return 1
				} else if l > r {
					return -1
				}
			}
			if len(leftItems) < len(rightItems) {
				fmt.Println("Left side is out of items")
				return 1
			}
		}
	}

	return 0
}

func parseLine(line string) []string {
	// remove outer brackets
	line = line[1 : len(line)-1]
	newLine := ""
	// split at , if brackets are closed
	bracketsCounter := 0
	for i := 0; i < len(line); i++ {
		if line[i] == '[' {
			bracketsCounter++
		} else if line[i] == ']' {
			bracketsCounter--
		}
		if line[i] == ',' && bracketsCounter == 0 {
			newLine += ";"
		} else {
			newLine += string(line[i])
		}
	}
	return strings.Split(newLine, ";")
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
