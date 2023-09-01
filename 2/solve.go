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
var OUTCOME_SCORES = [3]int{0, 3, 6}

func turnFromOutcome(outcome, opp int) int {

	switch outcome {
	case 0:
		// walk RPS backwards
		if opp == 0 {
			return 2
		}
		return opp - 1

	case 1:
		// stay
		return opp

	case 2:
		// walk RPS forward
		return (opp + 1) % 3
	}
	return 0
}

func calc(opp, outcome rune) int {
	oppNum, outcomeNum :=
		strings.IndexRune(FIRST_CHIFRE, opp),
		strings.IndexRune(SECOND_CHIFRE, outcome)
	turn := turnFromOutcome(outcomeNum, oppNum)
	outcomeScore := OUTCOME_SCORES[outcomeNum]
	choiceScore := CHOICE_TO_SCORE[turn]
	result := choiceScore + outcomeScore
	return result
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
		opponentChoice, outcomeVal := letters[0], letters[2]
		r := calc(opponentChoice, outcomeVal)

		result += r

	}
	fmt.Println(result)

}
