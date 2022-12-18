package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	name      string
	flowRate  int
	leadsTo   []*Valve
	distances map[*Valve]int
	parent    *Valve
	open      bool
}

func main() {
	lines := readLines("./data/pipesSmall.txt")
	var valves []*Valve
	for _, line := range lines {
		valves = parseLine(line, valves)
	}

	for _, valve := range valves {
		follower := []string{}
		for _, followerValve := range valve.leadsTo {
			follower = append(follower, followerValve.name)
		}
		fmt.Printf("I am valve %s and I lead to %v, my valve has a flowrate of %d\n", valve.name, follower, valve.flowRate)
	}
	// Part 1
	// Maximize the flow rate during 30 minutes
	// set distances
	for _, valve := range valves {
		valve.distances = make(map[*Valve]int)
		for _, valve2 := range valves {
			valve.distances[valve2] = 0
		}
	}

	valvesWithOutflow := []*Valve{}
	for _, valve := range valves {
		if valve.flowRate > 0 {
			valvesWithOutflow = append(valvesWithOutflow, valve)
		}
	}

	startValve := getValve(valves, "AA")

	// create pertubations of valves
	maxPressureRelease := 0
	N := len(valvesWithOutflow)
	p := make([]int, N+1)
	for i := 0; i <= N; i++ {
		p[i] = i
	}
	i := 1
	for i < N {
		p[i]--
		j := i % 2 * p[i]
		// swap pertubation[i] with pertubation[k]
		temp := valvesWithOutflow[i]
		valvesWithOutflow[i] = valvesWithOutflow[j]
		valvesWithOutflow[j] = temp
		i = 1
		for p[i] == 0 {
			p[i] = i
			i++
		}
		valvesToEvaluate := []*Valve{startValve}
		output := ""
		for _, valve := range valvesWithOutflow {
			valvesToEvaluate = append(valvesToEvaluate, valve)
			output += valve.name + " "
		}
		//fmt.Printf("Valves to evaluate: %s\n", output)

		pressureRelease := calcuateTotalScore(valves, valvesToEvaluate)
		if pressureRelease > maxPressureRelease {
			maxPressureRelease = pressureRelease
		}

	}
	fmt.Printf("Max pressure release: %d\n", maxPressureRelease)
}

func calcuateTotalScore(valves []*Valve, valvesToEvaluate []*Valve) int {

	for _, valve := range valvesToEvaluate {
		valve.open = false
	}

	pressureRelease := 0
	time := 0
	for i := 0; i < len(valvesToEvaluate)-1; i++ {
		distance := getDistance(valvesToEvaluate[i], valvesToEvaluate[i+1], valves)
		pressureRelease += calculatePressureRelease(valvesToEvaluate) * (distance + 1)
		valvesToEvaluate[i+1].open = true
		time += distance + 1
		if time > 30 {
			break
		}
	}
	pressureRelease += calculatePressureRelease(valvesToEvaluate) * (30 - time)
	return pressureRelease
}

func calculatePressureRelease(valves []*Valve) int {
	pressureRelease := 0
	var valvesOpen []string
	for _, valve := range valves {
		if valve.open {
			pressureRelease += valve.flowRate
			valvesOpen = append(valvesOpen, valve.name)
		}
	}
	//fmt.Printf("Valves open: %v\n", valvesOpen)
	return pressureRelease
}

func getDistance(valveStart *Valve, valveEnd *Valve, valves []*Valve) int {
	if valveStart.distances[valveEnd] != 0 {
		return valveStart.distances[valveEnd]
	}

	copyValves := []*Valve{}
	for _, v := range valves {
		copyValves = append(copyValves, v)
	}

	distances := make(map[*Valve]int64)
	for _, v := range copyValves {
		v.parent = nil
		distances[v] = int64(999999999999999999)
	}
	distances[valveStart] = 0

	for len(copyValves) > 0 {
		currentValve := getMinDistanceValve(copyValves, distances)
		copyValves = removeValve(copyValves, currentValve)
		for _, valve := range currentValve.leadsTo {
			if contains(copyValves, valve) {
				alt := distances[currentValve] + 1
				if alt < distances[valve] {
					distances[valve] = alt
					valve.parent = currentValve
				}
			}
		}
	}
	//fmt.Printf("Distance from %s to %s is %d\n", valveStart.name, valveEnd.name, distances[valveEnd])
	valveStart.distances[valveEnd] = int(distances[valveEnd])
	return int(distances[valveEnd])
}

func getMinDistanceValve(valves []*Valve, distances map[*Valve]int64) *Valve {
	minDistance := int64(999999999999999999)
	var minValve *Valve
	for _, v := range valves {
		if distances[v] < minDistance {
			minDistance = distances[v]
			minValve = v
		}
	}
	return minValve
}

func removeValve(valves []*Valve, valve *Valve) []*Valve {
	for i, v := range valves {
		if v == valve {
			return append(valves[:i], valves[i+1:]...)
		}
	}
	return valves
}

func contains(valves []*Valve, valve *Valve) bool {
	for _, v := range valves {
		if v == valve {
			return true
		}
	}
	return false
}

func parseLine(line string, valves []*Valve) []*Valve {

	flowRate, err := strconv.Atoi(strings.Split(strings.Split(line, ";")[0], "=")[1])
	if err != nil {
		log.Fatal(err)
	}
	valveName := line[6:8]
	valve := getValve(valves, valveName)
	if valve == nil {
		valve = &Valve{name: valveName}
		valves = append(valves, valve)
	}
	valve.flowRate = flowRate

	var valveFollowerNames []string
	if strings.Contains(line, " tunnels lead to valves ") {
		names := strings.Split(line, " tunnels lead to valves ")[1]
		valveFollowerNames = strings.Split(names, ", ")
	} else {
		valveFollowerNames = []string{strings.Split(line, " tunnel leads to valve ")[1]}
	}

	for _, nextValveName := range valveFollowerNames {
		nextValve := getValve(valves, nextValveName)
		if nextValve == nil {
			nextValve = &Valve{name: nextValveName}
			valves = append(valves, nextValve)
		}
		valve.leadsTo = append(valve.leadsTo, nextValve)
	}
	return valves
}

func getValve(valves []*Valve, name string) *Valve {
	for _, valve := range valves {
		if valve.name == name {
			return valve
		}
	}
	return nil
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
