package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	FIRST_CHIFRE  = "ABC"
	SECOND_CHIFRE = "XYZ"
)

var CHOICE_TO_SCORE = [3]int{1, 2, 3}

func getOutcomeScore(a, b int) int {
	// https://stackoverflow.com/a/2795421
	result := (3 + a - b) % 3
	switch result {
	case 0:
		// draw
		return 3

	case 2:
		// win
		return 6

	case 1:
		// lose
		return 0
	}
	return 0
}

func calc(opp, player rune) int {
	oppNum, playerNum :=
		strings.IndexRune(FIRST_CHIFRE, opp),
		strings.IndexRune(SECOND_CHIFRE, player)

	choiceScore := CHOICE_TO_SCORE[playerNum]
	outcomeScore := getOutcomeScore(oppNum, playerNum)

	return choiceScore + outcomeScore
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		letters := []rune(scanner.Text())
		// account for space
		opponentChoice, yourChoice := letters[0], letters[2]

		result += calc(opponentChoice, yourChoice)
	}
	fmt.Println(result)
}
