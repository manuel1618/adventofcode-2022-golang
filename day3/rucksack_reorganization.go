package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	cwd, _ := os.Getwd()
	path := (cwd +"/data/rucksacks.txt")
	lines := readLines(path)
	
	// Part 1
	points := 0
    for _,line := range lines {
		items1 := line[0:len(line)/2]
		items2 := line[len(line)/2:]

		if len(items1)!=len(items2) {
			fmt.Println("Rucksack not properly split")
		}
		
		char := ""
		for i := 0; i < len(items1); i++ {
			for j := 0; j < len(items2); j++ {
				if items1[i] == items2[j] {
					char = string(items1[i])
				}
			}
		 }
		 if char !=  "" {
			points += getPointsForLetter(char)
		 }
    }
	fmt.Printf("Total Score (1st part): %d\n",points)

	// Part 2
	points =0
	for i := 0; i< len(lines); i++ {
		if len(lines)-1 < i+2 {
			break
		}
		rucksack1 := lines[i]
		rucksack2 := lines[i+1]
		rucksack3 := lines[i+2]

		// the laziest impl of the planet, man
		char:=""
		for j := 0; j < len(rucksack1); j++ {
			for k:= 0; k < len(rucksack2); k++ {
				for l := 0; l < len(rucksack3); l++ {
					if  rucksack1[j] == rucksack2[k] && rucksack2[k]== rucksack3[l]{
						char = string(rucksack1[j])
						break
					}
				}
			}
		}
		
		if char !=  "" {
			points += getPointsForLetter(char)
		} else{
			fmt.Println("Did not find a badge")
		}
		i+=2
	}

	fmt.Printf("Total Score (2nd part): %d",points)
	
}


func getPointsForLetter(char string) int {
	pointsData := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	points := indexOf(pointsData,strings.ToLower(char))+1

	if strings.ToUpper(char)==char {
		points += 26
	}

	return points
}

// read a file and return the lines as an array
func readLines(path string) []string{
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


func indexOf(arr []string, val string) int {
    for pos, v := range arr {
        if v == val {
            return pos
        }
    }
    return -1
}