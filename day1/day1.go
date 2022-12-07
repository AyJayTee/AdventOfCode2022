package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("res/day1.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var totals []int
	var total int

	for scanner.Scan() {
		if scanner.Text() == "" {
			totals = append(totals, total)
			total = 0
		} else {
			amount, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			total += amount
		}
	}
	totals = append(totals, total)

	fmt.Println(totals)
	highest := ReturnThreeHighest(totals)
	fmt.Println(highest)

}

func ReturnHighestIndex(totals []int) int {
	highest := 0
	index := 0

	for i, t := range totals {
		if t > highest {
			highest = t
			index = i
		}
	}

	return index
}

func ReturnThreeHighest(totals []int) int {
	total := 0

	for i := 0; i < 3; i++ {
		i := ReturnHighestIndex(totals)
		total += totals[i]
		totals = append(totals[:i], totals[i+1:]...)
	}

	return total
}
