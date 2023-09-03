package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	alphabetLowerCase = "abcdefghijklmnopqrstuvwxyz"
	alphabetUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func runeToValue(letter rune) int {
	if strings.ContainsRune(alphabetLowerCase, letter) {
		return strings.IndexRune(alphabetLowerCase, letter) + 1
	}
	return strings.IndexRune(alphabetUpperCase, letter) + 27
}

func calc(a, b string) int {
	if !strings.ContainsAny(a, b) {
		return 0
	}
	sharedLetter := a[strings.IndexAny(a, b)]
	return runeToValue(rune(sharedLetter))
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
		contents := scanner.Text()
		firstHalf, secondHalf := contents[:len(contents)/2], contents[len(contents)/2:]
		result += calc(firstHalf, secondHalf)
	}

	fmt.Println(result)

}
