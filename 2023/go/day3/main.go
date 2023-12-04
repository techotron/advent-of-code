package main

import (
	"log"
	"os"
	"strings"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

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



func main() {
	inputLines := ReadFileIntoArray("input")

	//puzzleOne(inputLines)
	puzzleTwo(inputLines)
}

type Part struct {
	rowIdx				int
	rowLength			int
	partValue			string
	partPosition		[]int
	adjacentPosition	[]int
	isPart				bool
}

func puzzleOne(inputLines []string) {
	var grid [][]string
	var parts []Part
	var partValues int

	partRegex := regexp.MustCompile("[0-9]+")
	symbolRegex := regexp.MustCompile("[^0-9]+")

	for ridx, line := range inputLines {
		matchingIdx := partRegex.FindAllStringSubmatchIndex(line, -1)
		matchingPart := partRegex.FindAllStringSubmatch(line, -1)

		for i, _part := range matchingPart {
			var _adjacentPosition []int
			var _startIdx int
			var _endIdx int

			if matchingIdx[i][0] > 0 {
				_startIdx = matchingIdx[i][0] -1
			} else {
				_startIdx = matchingIdx[i][0]
			}

			if matchingIdx[i][1] < len(line) - 1 {
				_endIdx = matchingIdx[i][1] + 1
			} else {
				_endIdx = matchingIdx[i][1]
			}

			_adjacentPosition = []int{_startIdx, _endIdx}

			part := Part{
				rowIdx: ridx,
				partValue: _part[0],
				adjacentPosition: _adjacentPosition,
				partPosition: matchingIdx[i],
				rowLength: len(line) -1,
			}
			parts = append(parts, part)
		}

		var row []string
		for _, char := range line {
			_char := string(char)
			row = append(row, _char)
		}
		grid = append(grid, row)
	}

	lastRowIdx := len(grid) -1

	for _, part := range parts {
		var adjacentCharacters string
		var alignedCharacters string
		var characterInfront string
		var characterBehind string

		if part.rowIdx == 0 {
			// Look below only
			adjacentCharacters = inputLines[part.rowIdx + 1][part.adjacentPosition[0]:part.adjacentPosition[1]]
		} else if part.rowIdx == lastRowIdx {
			// Look above only
			adjacentCharacters = inputLines[part.rowIdx - 1][part.adjacentPosition[0]:part.adjacentPosition[1]]
		} else {
			// Look up and down
			adjacentCharacters = inputLines[part.rowIdx - 1][part.adjacentPosition[0]:part.adjacentPosition[1]]
			adjacentCharacters = adjacentCharacters + inputLines[part.rowIdx + 1][part.adjacentPosition[0]:part.adjacentPosition[1]]
		}

		if part.partPosition[0] > 0 {
			characterInfront = string(inputLines[part.rowIdx][part.adjacentPosition[0]])
		}
		if part.partPosition[1] < part.rowLength {
			characterBehind = string(inputLines[part.rowIdx][part.partPosition[1]])
		}

		alignedCharacters = characterInfront + characterBehind
		adjacentCharacters = strings.ReplaceAll(adjacentCharacters, ".", "") + strings.ReplaceAll(alignedCharacters, ".", "")

		adjacentSymbols := symbolRegex.FindAllStringSubmatch(adjacentCharacters, -1)
		fmt.Printf("%s: %s\n", part.partValue, adjacentSymbols)

		if len(adjacentSymbols) > 0 {
			part.isPart = true
			_partValue, _ := strconv.Atoi(part.partValue)
			partValues = partValues + _partValue
		} else {
			part.isPart = false
		}
	}

	fmt.Printf("Result of part 1: %v\n", partValues)
}

type Gear struct {
	rowIdx				int
	gearPosition		[]int
	adjacentPosition	[]int
	gears				[]int
}

func puzzleTwo(inputLines []string) {
	var partValues int
	var gears []Gear
	//var parts []Part

	gearRegex := regexp.MustCompile(`\*`)
	partRegex := regexp.MustCompile("[0-9]+")

	for ridx, line := range inputLines {
		matchingIdx := gearRegex.FindAllStringSubmatchIndex(line, -1)


		for i, _ := range matchingIdx {
			var _adjacentPosition []int
			var _startIdx int
			var _endIdx int

			if matchingIdx[i][0] > 0 {
				_startIdx = matchingIdx[i][0] - 1
			} else {
				_startIdx = matchingIdx[i][0]
			}

			if matchingIdx[i][1] < len(line) - 1 {
				_endIdx = matchingIdx[i][1]
			} else {
				_endIdx = matchingIdx[i][1]
			}

			_adjacentPosition = []int{_startIdx, _endIdx}

			gear := Gear{
				rowIdx: ridx,
				adjacentPosition: _adjacentPosition,
				gearPosition: matchingIdx[i],
				gears: []int{},
			}

			var matchingAdjacentParts []int
			var currentMatches [][]int
			var nextLineMatches [][]int
			var previousLineMatches [][]int

			var currentMatchesValue [][]string
			var nextLineMatchesValue [][]string
			var previousLineMatchesValue [][]string

			if ridx == 0 {
				currentMatches = partRegex.FindAllStringSubmatchIndex(inputLines[ridx], -1)
				nextLineMatches = partRegex.FindAllStringSubmatchIndex(inputLines[ridx + 1], -1)
				currentMatchesValue = partRegex.FindAllStringSubmatch(inputLines[ridx], -1)
				nextLineMatchesValue = partRegex.FindAllStringSubmatch(inputLines[ridx + 1], -1)
			} else if ridx == len(inputLines) -1 {
				previousLineMatches = partRegex.FindAllStringSubmatchIndex(inputLines[ridx - 1], -1)
				currentMatches = partRegex.FindAllStringSubmatchIndex(inputLines[ridx], -1)
				previousLineMatchesValue = partRegex.FindAllStringSubmatch(inputLines[ridx - 1], -1)
				currentMatchesValue = partRegex.FindAllStringSubmatch(inputLines[ridx], -1)
			} else {
				previousLineMatches = partRegex.FindAllStringSubmatchIndex(inputLines[ridx - 1], -1)
				currentMatches = partRegex.FindAllStringSubmatchIndex(inputLines[ridx], -1)
				nextLineMatches = partRegex.FindAllStringSubmatchIndex(inputLines[ridx + 1], -1)
				previousLineMatchesValue = partRegex.FindAllStringSubmatch(inputLines[ridx - 1], -1)
				currentMatchesValue = partRegex.FindAllStringSubmatch(inputLines[ridx], -1)
				nextLineMatchesValue = partRegex.FindAllStringSubmatch(inputLines[ridx + 1], -1)
			}

			_ = matchingAdjacentParts
			fmt.Printf("====== gear pos: %v:%v\n", gear.adjacentPosition[0], gear.adjacentPosition[1])
			//fmt.Println(previousLineMatches)
			//fmt.Println(currentMatches)
			//fmt.Println(nextLineMatches)


			//checkPrev:
			for i := gear.adjacentPosition[0]; i <= gear.adjacentPosition[1]; i++ {
				if len(previousLineMatches) > 0 {
					for idx, p := range previousLineMatches {
						if (p[0] == i) || (p[1] -1 == i) {
							fmt.Println("prev line p is: ", p)
							fmt.Println("i is: ", i)
							fmt.Println(previousLineMatchesValue[idx])
							_previousLineMatchesValue, _ := strconv.Atoi(previousLineMatchesValue[idx][0])
							if _previousLineMatchesValue != 0 {
								gear.gears = append(gear.gears, _previousLineMatchesValue)
							}

							previousLineMatchesValue[idx][0] = ""
							//break checkPrev
						}
					}
				}
			}
			//checkCurr:
			for i := gear.adjacentPosition[0]; i <= gear.adjacentPosition[1]; i++ {
				if len(currentMatches) > 0 {
					for idx, p := range currentMatches {
						if (p[0] == i) || (p[1] -1 == i) {
							fmt.Println("curr line p is: ", p)
							fmt.Println("i is: ", i)
							fmt.Println(currentMatchesValue[idx])
							_currentMatchesValue, _ := strconv.Atoi(currentMatchesValue[idx][0])
							if _currentMatchesValue != 0 {
								gear.gears = append(gear.gears, _currentMatchesValue)
							}
							currentMatchesValue[idx][0] = ""
							//break checkCurr
						}
					}
				}
			}
			//checkNext:
			for i := gear.adjacentPosition[0]; i <= gear.adjacentPosition[1]; i++ {
				if len(nextLineMatches) > 0 {
					for idx, p := range nextLineMatches {
						if (p[0] == i) || (p[1] -1 == i) {
							fmt.Println("next line p is: ", p)
							fmt.Println("i is: ", i)
							fmt.Println(nextLineMatchesValue[idx])
							_nextLineMatchesValue, _ := strconv.Atoi(nextLineMatchesValue[idx][0])
							if _nextLineMatchesValue != 0 {
								gear.gears = append(gear.gears, _nextLineMatchesValue)
							}
							nextLineMatchesValue[idx][0] = ""
							//break checkNext
						}
					}
				}
			}



			gears = append(gears, gear)
		}
	}


	for _, gear := range gears {
		var ratio int
		fmt.Println(gear.gears)
		if len(gear.gears) == 2 {
			ratio = gear.gears[0] * gear.gears[1]
		}
		partValues = partValues + ratio
	}

	fmt.Printf("Result of part 2: %v\n", partValues)
}

