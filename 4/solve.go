package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) []int {
	pairs := strings.Split(line, ",")
	result := []int{}
	for _, pair := range pairs {
		strNums := strings.Split(pair, "-")
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				log.Fatal(err)
				break
			}
			result = append(result, num)
		}
	}
	return result
}

func inRange(num, a, b int) bool {
	return num >= a && num <= b
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
		nums := parseLine(scanner.Text())
		a1, a2, b1, b2 := nums[0], nums[1], nums[2], nums[3]
		if inRange(a1, b1, b2) || inRange(a2, b1, b2) {
			result++
		} else if inRange(b1, a1, a2) || inRange(b2, a1, a2) {
			result++
		}

	}
	fmt.Println(result)
}
