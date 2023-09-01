package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total []int
	current := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			total = append(total, current)
			current = 0
		} else {
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			current += val
		}
	}
	// hacky EOF handling
	if current != 0 {
		total = append(total, current)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	highest := 0
	for _, v := range total {
		if v > highest {
			highest = v
		}
	}
	fmt.Println(highest)

}
