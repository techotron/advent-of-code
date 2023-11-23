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

type ElfCalorie struct {
	Name  string
	Total int
}

func Sum(s []int) (result int) {
	for _, i := range s {
		result = result + i
	}
	return result
}

func ReadFileIntoArray(path string) []string {
	var fileLines []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fileLines = append(fileLines, strings.TrimSpace(scanner.Text()))
	}

	return fileLines
}

func ItoA(i int) string {
	return strconv.Itoa(i)
}

func AtoI(s string) int {
	_s, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return _s
}

func SortElves(elves []ElfCalorie) []ElfCalorie {
	sort.SliceStable(elves, func(i, j int) bool { return elves[i].Total < elves[j].Total })
	return elves
}

func main() {
	counter := "0"
	var elfCalorieCounter []int
	lines := ReadFileIntoArray("input")

	var elves []ElfCalorie

	for _, line := range lines {
		if line == "" {
			elf := ElfCalorie{Name: counter, Total: Sum(elfCalorieCounter)}
			elves = append(elves, elf)

			counter = ItoA(AtoI(counter) + 1)
			elfCalorieCounter = []int{}
		} else {
			elfCalorieCounter = append(elfCalorieCounter, AtoI(line))
		}
	}

	SortElves(elves)

	topCalories := elves[len(elves)-1].Total
	topThreeCalories := elves[len(elves)-1].Total + elves[len(elves)-2].Total + elves[len(elves)-3].Total

	fmt.Printf("Highest number of calories: %v\n", topCalories)
	fmt.Printf("Sum of top three: %v\n", topThreeCalories)
}
