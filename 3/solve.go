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

func runeToValue(letter string) int {
	if strings.Contains(alphabetLowerCase, letter) {
		return strings.Index(alphabetLowerCase, letter) + 1
	}
	return strings.Index(alphabetUpperCase, letter) + 27
}

func calc(a, b, c string) int {
	if !strings.ContainsAny(a, b) {
		return 0
	}
	occurrence := strings.IndexAny(a, b)
	sharedLetter := string(a[occurrence])

	if !strings.Contains(c, sharedLetter) {
		a = strings.ReplaceAll(a, sharedLetter, "")
		b = strings.ReplaceAll(b, sharedLetter, "")
		return calc(a, b, c)
	} else {
		return runeToValue(sharedLetter)
	}

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0

	tmp := []string{}

	for scanner.Scan() {
		if len(tmp) < 2 {
			tmp = append(tmp, scanner.Text())
			continue
		} else {
			result += calc(tmp[0], tmp[1], scanner.Text())
			tmp = []string{}
		}
	}

	fmt.Println(result)

}
