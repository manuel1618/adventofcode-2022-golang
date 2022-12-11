package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	id                  int
	items               []int
	operation           string
	testDenominator     int
	targetMonkeySuccess int
	targetMonkeyFailure int
	inspectionCounter   int
}

func main() {
	lines := readLines("./data/monkeys.txt")
	var monkeys []*Monkey
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fmt.Println(line)
		if strings.HasPrefix(line, "Monkey") {
			id, err := strconv.Atoi(strings.ReplaceAll(strings.Split(line, " ")[1], ":", ""))
			var monkeyItems []int
			itemsLine := strings.Split(lines[i+1][18:], ",")
			for _, itemS := range itemsLine {
				item, err := strconv.Atoi(strings.TrimSpace(itemS))
				if err != nil {
					log.Fatal(err)
				}
				monkeyItems = append(monkeyItems, item)
			}

			operation := lines[i+2][19:]

			testDenominator, err := strconv.Atoi(strings.TrimSpace(lines[i+3][21:]))
			targetMonkeySuccess, err := strconv.Atoi(strings.TrimSpace(lines[i+4][29:]))
			targetMonkeyFailure, err := strconv.Atoi(strings.TrimSpace(lines[i+5][30:]))

			if err != nil {
				log.Fatal(err)
			}
			monkey := Monkey{id: id, items: monkeyItems, operation: operation,
				testDenominator: testDenominator, targetMonkeySuccess: targetMonkeySuccess,
				targetMonkeyFailure: targetMonkeyFailure, inspectionCounter: 0}
			monkeys = append(monkeys, &monkey)
		}
	}

	numberOfRonds := 10000
	for i := 0; i < numberOfRonds; i++ {
		playOneRound(monkeys)
		if i%1000 == 0 {
			printMonkeys(monkeys)
		}
	}
	printMonkeys(monkeys)
}

func playOneRound(monkeys []*Monkey) {

	superModulo := 1
	for _, monkey := range monkeys {
		superModulo = superModulo * monkey.testDenominator
	}

	for _, monkey := range monkeys {
		//fmt.Printf("Monkey %d:\n", monkey.id)
		for _, item := range monkey.items {
			//fmt.Printf("  Monkey inspects an item with a worry level of: %v\n", item)
			monkey.inspectionCounter += 1
			newItem := monkeyOperation(item, monkey.operation)
			// Part 1
			// newItem = newItem / 3
			// Part2
			newItem = newItem % superModulo

			//fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %v\n", newItem)
			var targetMonkey *Monkey
			if newItem%monkey.testDenominator == 0 {
				//fmt.Printf("    Current worry level is divisible by %v\n", monkey.testDenominator)
				targetMonkey = getMonkeyById(monkeys, monkey.targetMonkeySuccess)
			} else {
				//fmt.Printf("    Current worry level is not divisible by %v\n", monkey.testDenominator)
				targetMonkey = getMonkeyById(monkeys, monkey.targetMonkeyFailure)
			}
			// throw item
			if len(monkey.items) == 1 {
				monkey.items = make([]int, 0)
			} else {
				monkey.items = monkey.items[1:]
			}
			targetMonkey.items = append(targetMonkey.items, newItem)
			//fmt.Printf("    Item with worry level %v is thrown to monkey %d\n", newItem, targetMonkey.id)
		}
	}
}

func monkeyOperation(item int, op string) int {
	var newItem int
	if strings.Contains(op, "*") {
		if strings.Contains(op, "old * old") {
			newItem = item * item
		} else {
			multiplier, err := strconv.Atoi(strings.TrimSpace(strings.Split(op, "*")[1]))
			if err != nil {
				log.Fatal(err)
			}
			newItem = item * multiplier
			//fmt.Printf("    Worry level is multiplied by %v to %v\n", multiplier, newItem)
		}
	} else if strings.Contains(op, "+") {
		add, err := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(op, "old", ""), "+", "")))
		if err != nil {
			log.Fatal(err)
		}
		newItem = item + add
		//fmt.Printf("    Worry level is increased by %v to %v\n", add, newItem)
	}
	return newItem

}

func getMonkeyById(monkeys []*Monkey, id int) *Monkey {
	for _, monkey := range monkeys {
		if monkey.id == id {
			return monkey
		}
	}
	return nil
}

func printMonkeys(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d: %v - inspections: %d\n", monkey.id, monkey.items, monkey.inspectionCounter)
	}
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
