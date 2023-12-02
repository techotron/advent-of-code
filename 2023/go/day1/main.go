package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	"regexp"
	"strconv"
)

func main() {
	inputLines := ReadFileIntoArray("input")

	puzzleOne(inputLines)
	puzzleTwo(inputLines)
}

func puzzleTwo(inputLines []string) {
	re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[0-9]")

	var allNumbers []int
	for _, line := range inputLines {
	    // Well, it was today when I found out that regexp doesn't support over lapping matches so this hack is to workaround that in the time I had to solve this puzzle!
		_line := strings.ReplaceAll(line, "oneight", "oneeight")
		_line = strings.ReplaceAll(_line, "threeight", "threeeight")
		_line = strings.ReplaceAll(_line, "fiveight", "fiveeight")
		_line = strings.ReplaceAll(_line, "nineight", "nineeight")
		_line = strings.ReplaceAll(_line, "twone", "twoone")
		_line = strings.ReplaceAll(_line, "sevenine", "sevennine")
		_line = strings.ReplaceAll(_line, "eightwo", "eighttwo")

		matches := re.FindAllStringSubmatch(_line, -1)
		first := matches[0][0]
		second := matches[len(matches)-1][0]

		var calibrationNumber string

		if IsNumber(first) {
			calibrationNumber = calibrationNumber + first
		} else {
			stringNumberValue := GetNumberValue(first)
			calibrationNumber = calibrationNumber + stringNumberValue
		}

		if IsNumber(second) {
			calibrationNumber = calibrationNumber + second
		} else {
			stringNumberValue := GetNumberValue(second)
			calibrationNumber = calibrationNumber + stringNumberValue
		}

		_cn, _ := strconv.Atoi(calibrationNumber)
		allNumbers = append(allNumbers, _cn)
	}
	fmt.Printf("Answer for part 2: %d\n", Sum(allNumbers))
}

func puzzleOne(inputLines []string) {
	var allNumbers []int
	for _, line := range inputLines {
		var lineNumbers []string
		for _, character := range line {
			_character := string(character)
			if IsNumber(_character) {
				lineNumbers = append(lineNumbers, _character)
			}
		}
		var calibrationNumber string
		calibrationNumber = fmt.Sprintf("%s%s", lineNumbers[0], lineNumbers[len(lineNumbers)-1])
		_cn, _ := strconv.Atoi(calibrationNumber)
		allNumbers = append(allNumbers, _cn)
	}
	fmt.Printf("Answer for part 1: %d\n", Sum(allNumbers))
}

func GetNumberValue(s string) (val string) {
	switch s {
	case "one":
		val = "1"
	case "two":
		val = "2"
	case "three":
		val = "3"
	case "four":
		val = "4"
	case "five":
		val = "5"
	case "six":
		val = "6"
	case "seven":
		val = "7"
	case "eight":
		val = "8"
	case "nine":
		val = "9"
	case "oneight":
		val = "8"
	case "threeight":
		val = "8"
	case "fiveight":
		val = "8"
	case "nineight":
		val = "8"
	case "twone":
		val = "1"
	case "sevenine":
		val = "9"
	case "eightwo":
		val = "2"
	}
	return val
}

func IsNumber(s string) bool {
	return regexp.MustCompile(`\d`).MatchString(s)
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