package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type Colors struct {
	blue 	int
	red 	int
	green	int
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

func Sum(s []int) (result int) {
	for _, i := range s {
		result = result + i
	}
	return result
}

func main() {
	inputLines := ReadFileIntoArray("input")

	puzzleOne(inputLines)
	puzzleTwo(inputLines)
}

func puzzleOne(inputLines []string) {
	var gameIds int
	gameLoop:
	for _, game := range inputLines {
		gameId := strings.ReplaceAll(strings.Split(game, ":")[0], "Game ", "")
		_ = gameId

		gamesPlayed := strings.Split(strings.Split(game, ":")[1], ";")

		for _, gamePlayed := range gamesPlayed {
			cubesPicked := strings.Split(gamePlayed, ",")
			for _, colours := range cubesPicked {

				_colours := strings.TrimSpace(colours)
				number := strings.Split(_colours, " ")[0]
				_number, _ := strconv.Atoi(number)
				colour := strings.Split(_colours, " ")[1]

				switch colour {
				case "red":
					if _number > 12 {
						continue gameLoop
					}
				case "green":
					if _number > 13 {
						continue gameLoop
					}
				case "blue":
					if _number > 14 {
						continue gameLoop
					}
				}
			}
		}
		_gameId, _ := strconv.Atoi(gameId)
		gameIds = gameIds + _gameId
	}
	fmt.Printf("Result of part 1: %v\n", gameIds)
}


func puzzleTwo(inputLines []string) {
	var powerTotal int
	allGameData := make(map[string]Colors)

	for _, game := range inputLines {
		gameId := strings.ReplaceAll(strings.Split(game, ":")[0], "Game ", "")
		_ = gameId

		gamesPlayed := strings.Split(strings.Split(game, ":")[1], ";")

		var blue int
		var red int
		var green int

		for _, gamePlayed := range gamesPlayed {

			cubesPicked := strings.Split(gamePlayed, ",")
			for _, colours := range cubesPicked {
				_colours := strings.TrimSpace(colours)
				number := strings.Split(_colours, " ")[0]
				_number, _ := strconv.Atoi(number)
				colour := strings.Split(_colours, " ")[1]
				switch colour {
				case "blue":
					if _number > blue {
						blue = _number
					}
				case "red":
					if _number > red {
						red = _number
					}
				case "green":
					if _number > green {
						green = _number
					}
				}
			}
		}

		allGameData[gameId] = Colors{
			blue: blue,
			red: red,
			green: green,
		}

	}

	for _, result := range allGameData {
		setPower := result.blue * result.green * result.red
		powerTotal = powerTotal + setPower
	}

	fmt.Printf("Result of part 2: %v\n", powerTotal)
}
