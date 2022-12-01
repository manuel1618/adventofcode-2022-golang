package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

    cwd, err := os.Getwd()
    file, err := os.Open(cwd+ "/data/elvesWithCalories.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Initializing
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    counter := int64(0)
    max_calories := int64(0)
    elf_counter := int64(1)
    var elf_calories []int64

    // Scon the file
    for scanner.Scan() {
        line:= scanner.Text()
        if line == ""  {
            elf_calories=append(elf_calories,counter)
            if counter>max_calories {
                max_calories = counter
            }
            elf_counter += int64(1)
            counter = 0
        } else {
            calories, _ := strconv.ParseInt(line, 10, 0)
            counter += calories
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Output 1: Max Calories of one elf
    fmt.Printf("The elf: %d has the most calories %d \n", elf_counter, max_calories)

    // Output 2: Get the top three elves' calories
    sort.Slice(elf_calories, func(i, j int) bool { return elf_calories[i] > elf_calories[j]})
    max_calories_top3 := int64(0)
    for i := 0; i < 3; i++ {
        max_calories_top3 += elf_calories[i]
    }
    fmt.Printf("The top three elves have a total amount of calories %d", max_calories_top3)

}