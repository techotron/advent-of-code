package main

import (
	"log"
	"os"
	"strings"
	"bufio"
	"fmt"
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

type Card struct {
	CardNumber				string
	MatchingNumbers 		[]string
	Points					int
	NumberOfMatchingNumbers	int
	NumberOfCopies			int
	CopiedCards				[]Card
}

func puzzleOne(inputLines []string) {
	var cards []Card
	var result int

	for _, line := range inputLines {
		cardNumber := strings.Split(strings.Split(line, ":")[0], " ")[1]
		numberSets := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumberSet := numberSets[0]
		winningNumbers := strings.Split(winningNumberSet, " ")
		myNumberSet := numberSets[1]
		myNumbers := strings.Split(myNumberSet, " ")
		//fmt.Println(cardNumber)
		//fmt.Println(winningNumbers)
		//fmt.Println(myNumbers)

		card := Card{
			CardNumber: cardNumber,
		}

		for _, myNumber := range myNumbers {
			_myNumber := string(myNumber)
			//fmt.Printf("card: %s | myNumber: %s\n", cardNumber, _myNumber)
			for _, winningNumber := range winningNumbers {
				_winningNumber := string(winningNumber)
				if _myNumber == "" {
					continue
				}
				if _winningNumber == "" {
					continue
				}
				if _myNumber == _winningNumber {
					card.MatchingNumbers = append(card.MatchingNumbers, _myNumber)
				}
			}
		}

		cards = append(cards, card)
	}

	for _, card := range cards {
		numberOfMatches := len(card.MatchingNumbers)
		var points int

		for i := 1; i <= numberOfMatches; i++ {
			if i == 1 {
				points = 1
			} else {
				//fmt.Println(points)
				points = points * 2
			}
		}
		result = result + points
	}


	fmt.Printf("Result of part 1: %v\n", result)
}

func puzzleTwo(inputLines []string) {
	var cards []Card
	//var _cards []Card
	cardsWon := make(map[int][]Card)
	var result int

	for _, line := range inputLines {
		cardNumber := strings.Split(strings.Split(line, ":")[0], " ")[1]
		numberSets := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumberSet := numberSets[0]
		winningNumbers := strings.Split(winningNumberSet, " ")
		myNumberSet := numberSets[1]
		myNumbers := strings.Split(myNumberSet, " ")
		//fmt.Println(cardNumber)
		//fmt.Println(winningNumbers)
		//fmt.Println(myNumbers)

		card := Card{
			CardNumber: cardNumber,
		}

		for _, myNumber := range myNumbers {
			_myNumber := string(myNumber)
			//fmt.Printf("card: %s | myNumber: %s\n", cardNumber, _myNumber)
			for _, winningNumber := range winningNumbers {
				_winningNumber := string(winningNumber)
				if _myNumber == "" {
					continue
				}
				if _winningNumber == "" {
					continue
				}
				if _myNumber == _winningNumber {
					card.MatchingNumbers = append(card.MatchingNumbers, _myNumber)
					card.NumberOfMatchingNumbers = card.NumberOfMatchingNumbers + 1
				}
			}
		}

		cards = append(cards, card)
		//_cards = append(_cards, card)
	}

	_, _ = strconv.Atoi("0")



	// Initialize cardsWon
	for idx, card := range cards {
		card.NumberOfCopies = len(card.MatchingNumbers)
		cardsWon[idx] = append(cardsWon[idx], card)

	}

	var cardsToCopy int
	// idx 1
	for idx, _ := range cards {
		if idx == 0 {
			continue
		}

		cardsToCopy = cardsWon[idx-1][0].NumberOfCopies
		numberOfCardsToCopy := len(cardsWon[idx-1])

		fmt.Println("indx: ", idx, " cardstocopy: ", cardsToCopy)

		for j := 0; j < numberOfCardsToCopy; j++ {
			for i := 0; i < cardsToCopy; i ++ {
				fmt.Println("  copying card: ", cards[idx], " to ", idx + i)
				cardsWon[idx + i] = append(cardsWon[idx + i], cards[idx])
			}
		}


	}

	//for cardNumber, _cw := range cardsWon {
	//	numberOfCardsToCopy := len(_cw.MatchingNumbers)
	//	fmt.Println("Looking at ", cardNumber, " number of cards to copy: ", numberOfCardsToCopy)
	//	for i := 1; i < numberOfCardsToCopy + 1; i ++ {
	//		nextCardIdx := i + 1
	//		nextCardIdxStr := strconv.Itoa(nextCardIdx)
	//		fmt.Println("  Adding number of copies to cardNumber: ", cardsWon[nextCardIdxStr])
	//
	//	}
	//
	//	//for _, _cards := range cardsWon[cardNumber] {
	//	//	numberOfCardsToCopy := len(cardsWon[cardNumber])
	//	//
	//	//	for i := 0; i < numberOfCardsToCopy; i++ {
	//	//		nextCardIdx := strconv.Itoa(i + 1)
	//	//		cardsWon[nextCardIdx] = append(cardsWon[nextCardIdx], _cards)
	//	//	}
	//	//}
	//}

	//for _, card := range cards {
	//	cardNumberAsInt, _ := strconv.Atoi(card.CardNumber)
	//
	//	startIdx, _ := strconv.Atoi(card.CardNumber)
	//	endIdx := startIdx + (len(card.MatchingNumbers))
	//	cardsToCopy := cards[startIdx:endIdx]
	//	fmt.Println("card ", card.CardNumber, " cards to copy: ", cardsToCopy)
	//	for i, cardToCopy := range cardsToCopy {
	//		nextCard := cardNumberAsInt + i + 1
	//		nextCardAsStr := strconv.Itoa(nextCard)
	//		fmt.Println("  next card: ", nextCardAsStr)
	//		fmt.Println("  next card has already: ", cardsWon[nextCardAsStr])
	//		cardsWon[nextCardAsStr] = append(cardsWon[nextCardAsStr], cardToCopy)
	//		fmt.Println("  next card now has: ", cardsWon[nextCardAsStr])
	//	}
	//}

	for idx, value := range cardsWon {
		fmt.Println("card idx: ", idx, " card value: ", len(value))
		result = result + len(value)
	}



	fmt.Printf("Result of part 2: %v (should be 30)\n", result)
}