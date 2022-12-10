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
	path := ("./data/instructions.txt")
	lines := readLines(path)

	// Part 1
	startAndDeltaCycle := []int{20,40}
	score := 0
	cycle := 1
	value := 1
	for i := 0; i < len(lines); i++ {
		//fmt.Printf("Cycle: %v, Value: %v, command: %v\n", cycle, value, lines[i])
		score = updateScore(startAndDeltaCycle, cycle, value, score)
		op := lines[i]
		if op == "noop" {
			cycle++	
			continue
		} else {
			if strings.Split(lines[i], " ")[0] == "addx" {
				v,err := strconv.Atoi(strings.Split(lines[i], " ")[1])
				if err != nil {
					log.Fatal(err)
				}
				cycle++
				score = updateScore(startAndDeltaCycle, cycle, value, score)
				value += v
				cycle++
				continue
			}
		}
	}

	// Part 2
	startAndDeltaCycle = []int{20,40}
	cycle = 1
	value = 1
	sprite := []int{1,2,3}
	msg := ""
	screenWidth := 40
	for i := 0; i < len(lines); i++ {
		//fmt.Printf("Cycle: %v, Value: %v, command: %v\n", cycle, value, lines[i])
		msg = writeMsg(cycle,sprite,screenWidth,msg)
		op := lines[i]
		if op == "noop" {
			cycle++	
			continue
		} else {
			if strings.Split(lines[i], " ")[0] == "addx" {
				v,err := strconv.Atoi(strings.Split(lines[i], " ")[1])
				if err != nil {
					log.Fatal(err)
				}
				cycle++
				msg = writeMsg(cycle,sprite,screenWidth,msg)
				value += v
				sprite = []int{value,value+1,value+2}
				cycle++
				continue
			}
		}
	}
	deserializeAndPrintMsg(msg, screenWidth)

}

func deserializeAndPrintMsg(msg string, width int) {

	output := ""
	for i := 1; i <= len(msg); i++ {
		output += string(msg[i-1])
		if i%width == 0 {
			fmt.Printf("Output: %v\n", output)
			output =""
		}
	}
}

func writeMsg(cycle int, stripe []int, screenWidth int, msg string) string {

	if cycle > screenWidth {
		cycle -= cycle / screenWidth * screenWidth
	}

	
	for _,v := range stripe {
		if v == cycle {
			msg += "#"
			return msg
		}
	}
	msg+= "."
	return msg
}

func updateScore(startAndDeltaCycle []int,  cycle int, value int,score int) int {
	if outputTest(startAndDeltaCycle, cycle) {
		score += value*cycle
		fmt.Printf("Score update: Cycle:%v, Value: %v, Delta:%v, Score:%v\n", cycle,value, value*cycle,score)

	}
	return score
}


func outputTest(startAndDeltaCycle []int, cycle int) bool  {

	startCycle := startAndDeltaCycle[0]
	deltaCycle := startAndDeltaCycle[1]

	for j := 0; j < cycle; j++ {
		targetCycle := startCycle + j *deltaCycle
		if targetCycle > cycle {
			return false
		}
		if  targetCycle == cycle {
			return true
		}
	}
	return false
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
