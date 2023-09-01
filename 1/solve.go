package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
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
	// getting the top 3
	slices.SortFunc(total,
		func(a int, b int) int {

			return cmp.Compare(b, a)
		})
	result := 0
	for i := 0; i < 3; i++ {
		result += total[i]
	}

	fmt.Println(result)

}
