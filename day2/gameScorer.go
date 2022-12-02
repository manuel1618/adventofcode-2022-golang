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
	path := (cwd +"/data/game.txt")
	lines := readLines(path)
	
	// Part 1
	pointsForItemOpponent := [3]string{"A", "B", "C"}
	pointsForItemSelf := [3]string{"X", "Y", "Z"}

	points := 0
    for _,line := range lines {
        itemOpponent := strings.Split(line, " ")[0]
        itemSelf := strings.Split(line, " ")[1]

		pointsItem := indexOf(pointsForItemSelf[:],itemSelf)+1

		// we win if our index is one bigger than the opponent's
		// [Rock, Paper, Scissors] vs [Rock, Paper, Scissors]
		pointsWin := 0 // loose is default
		indexOpponentItem :=  indexOf(pointsForItemOpponent[:],itemOpponent)
		indexOpponentSelf :=  indexOf(pointsForItemSelf[:],itemSelf)
		if indexOpponentSelf==indexOpponentItem+1 {
			pointsWin = 6 // win
		} else if (indexOpponentItem==indexOpponentSelf){
			pointsWin = 3 // draw
		}
		// corner case case 
		if indexOpponentItem == 2 && indexOpponentSelf== 0 {
			pointsWin = 6 // also a win
		}

		// Error Checking
		if pointsItem == -1 || pointsWin == -1 {
			fmt.Println("Something went wrong.")
		}
		points += pointsItem+pointsWin
    }
	fmt.Printf("Total Score (1st part): %d\n",points)

	// Part 2
	points = 0
	for _,line := range lines{
		itemOpponent := strings.Split(line, " ")[0]
        itemSelf := strings.Split(line, " ")[1]

		// [loose, draw, win] -> [0, 3, 6]
		pointsWinBasedOnStrategy := indexOf(pointsForItemSelf[:],itemSelf)*3
 
		// shift -1 for loose, 0 for draw and +1 for win
		// [Rock, Paper, Scissors] vs [Rock, Paper, Scissors]
		shift := indexOf(pointsForItemSelf[:],itemSelf)-1
		indexOpponentItem :=  indexOf(pointsForItemOpponent[:],itemOpponent)
		indexOfItemSelf :=  indexOpponentItem+shift
		// bodaries
		if indexOfItemSelf == -1 {
			indexOfItemSelf = 2
		} else if indexOfItemSelf == 3{
			indexOfItemSelf = 0
		}
		poinstBasedOnItemSelf := indexOfItemSelf+1

		if pointsWinBasedOnStrategy == -1 || poinstBasedOnItemSelf == -1 {
			fmt.Println("Something went wrong.")
		}

		points += poinstBasedOnItemSelf+pointsWinBasedOnStrategy
	}
	fmt.Printf("Total Score (2nd part): %d",points)
	
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

// Get the Winning Points like a beginner - DEPRICATED, but kept for reference
func getWinningPoints(itemOpponent string, itemSelf string) int {
	if itemOpponent == "A" {
		if itemSelf == "X" {
			return 3
		} else if itemSelf == "Y" {
			return 6
		} else if itemSelf == "Z" {
			return 0
		}
	} else if itemOpponent == "B" {
		if itemSelf == "X" {
			return 0
		} else if itemSelf == "Y" {
			return 3
		} else if itemSelf == "Z" {
			return 6
			}

	} else if itemOpponent == "C" {
		if itemSelf == "X" {
			return 6
		} else if itemSelf == "Y" {
			return 0
		} else if itemSelf == "Z" {
			return 3
			}
	}
	return -1
}



func indexOf(arr []string, val string) int {
    for pos, v := range arr {
        if v == val {
            return pos
        }
    }
    return -1
}