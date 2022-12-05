package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Print() {
	for i := 0; i < len(*s); i++ {
		fmt.Printf("%s ", (*s)[i])
	}
	fmt.Println()
}

func (s *Stack) Reverse() Stack {
	var newStack Stack
	for i := len(*s) - 1; i >= 0; i-- {
		newStack.Push((*s)[i])
	}
	return newStack
}

func (s *Stack) Move(qnt int, to *Stack) {
	for i := 0; i < qnt; i++ {
		value, _ := s.Pop()
		to.Push(value)
	}
}

func (s *Stack) MoveInOrder(qnt int, to *Stack) {
	values := make([]string, qnt)
	for i := 0; i < qnt; i++ {
		value, _ := s.Pop()
		values[i] = value
	}
	for i := len(values) - 1; i >= 0; i-- {
		to.Push(values[i])
	}
}

func main() {
	cwd, _ := os.Getwd()
	path := (cwd + "/data/shipStacks.txt")
	lines := readLines(path)

	var stacks [9]Stack
	for i := 0; i < 9; i++ {
		stacks[i] = Stack{}
	}

	// Part 1
	msg := ""
	// Reset Stacks
	for i := 0; i < 9; i++ {
		stacks[i] = Stack{}
	}
	for _, line := range lines {

		if strings.Contains(line, "[") {
			for i := 0; i < 9; i++ {
				value := line[i*4+1 : i*4+2]
				if value != " " {
					stacks[i].Push(value)
				}
			}
		} else if line == " 1   2   3   4   5   6   7   8   9 " {
			// inverse stacks because I read them in the wrong order
			for i := 0; i < 9; i++ {
				stacks[i] = stacks[i].Reverse()
			}
		} else if strings.Contains(line, "move") {
			qnt, err := strconv.Atoi(strings.Split(line, " ")[1])
			from, err := strconv.Atoi(strings.Split(line, " ")[3])
			to, err := strconv.Atoi(strings.Split(line, " ")[5])
			if err != nil {
				log.Fatal(err)
			}
			stacks[from-1].Move(qnt, &stacks[to-1])
		}
	}

	for _, stack := range stacks {
		stack.Print()
		letter, _ := stack.Pop()
		msg += letter
	}
	fmt.Printf("The message at the end is: %s\n", msg)

	// Part 2 - move it differently
	msg = ""
	// Reset Stacks
	for i := 0; i < 9; i++ {
		stacks[i] = Stack{}
	}
	for _, line := range lines {

		if strings.Contains(line, "[") {
			for i := 0; i < 9; i++ {
				value := line[i*4+1 : i*4+2]
				if value != " " {
					stacks[i].Push(value)
				}
			}
		} else if line == " 1   2   3   4   5   6   7   8   9 " {
			// inverse stacks because I read them in the wrong order
			for i := 0; i < 9; i++ {
				stacks[i] = stacks[i].Reverse()
			}
		} else if strings.Contains(line, "move") {
			qnt, err := strconv.Atoi(strings.Split(line, " ")[1])
			from, err := strconv.Atoi(strings.Split(line, " ")[3])
			to, err := strconv.Atoi(strings.Split(line, " ")[5])
			if err != nil {
				log.Fatal(err)
			}
			stacks[from-1].MoveInOrder(qnt, &stacks[to-1])
		}

	}
	for _, stack := range stacks {
		stack.Print()
		letter, _ := stack.Pop()
		msg += letter
	}

	fmt.Printf("The message at the end is: %s\n", msg)
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
