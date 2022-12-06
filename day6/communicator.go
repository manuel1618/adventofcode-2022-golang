package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	cwd, _ := os.Getwd()
	path := (cwd + "/data/stream.txt")

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initializing
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	// part 1
	lookahead := 4
	fmt.Println(line)
	for i := 0; i < len(line)-lookahead; i++ {
		chars := make([]string, lookahead)
		// constant lookahead might be not efficient but it my only idea rn
		markerFound := true
		for j := 0; j < lookahead; j++ {
			newChar := string(line[i+j])
			if contains(chars, newChar) {
				markerFound = false
				break
			} else {
				chars[j] = newChar
			}
		}

		if markerFound {
			fmt.Printf("Found marker at %d\n", i+lookahead)
			marker := ""
			for k := 0; k < len(chars); k++ {
				marker += string(chars[k])
			}
			fmt.Printf("Marker is %s\n", marker)
			break
		}
	}

	// part 2
	lookahead = 14
	fmt.Println(line)
	for i := 0; i < len(line)-lookahead; i++ {
		chars := make([]string, lookahead)
		// constant lookahead might be not efficient but it my only idea rn
		markerFound := true
		for j := 0; j < lookahead; j++ {
			newChar := string(line[i+j])
			if contains(chars, newChar) {
				markerFound = false
				break
			} else {
				chars[j] = newChar
			}
		}

		if markerFound {
			fmt.Printf("Found marker at %d\n", i+lookahead)
			marker := ""
			for k := 0; k < len(chars); k++ {
				marker += string(chars[k])
			}
			fmt.Printf("Marker is %s\n", marker)
			break
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
