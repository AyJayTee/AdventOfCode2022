package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("res/day4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var pairs []string
	for scanner.Scan() {
		pairs = append(pairs, scanner.Text())
	}

	var count int
	for _, p := range pairs {
		pair := strings.Split(p, ",")
		leftBounds := strings.Split(pair[0], "-")
		rightBounds := strings.Split(pair[1], "-")
		if DetectOverlap(leftBounds, rightBounds) {
			fmt.Println(pair)
			count++
		}
	}

	fmt.Println(count)
}

func DetectEncapsulated(leftBounds []string, rightBounds []string) bool {
	leftLower, _ := strconv.Atoi(leftBounds[0])
	leftUpper, _ := strconv.Atoi(leftBounds[1])
	rightLower, _ := strconv.Atoi(rightBounds[0])
	rightUpper, _ := strconv.Atoi(rightBounds[1])

	if rightLower >= leftLower && rightUpper <= leftUpper {
		return true
	}

	if leftLower >= rightLower && leftUpper <= rightUpper {
		return true
	}

	return false
}

func DetectOverlap(leftBounds []string, rightBounds []string) bool {
	leftLower, _ := strconv.Atoi(leftBounds[0])
	leftUpper, _ := strconv.Atoi(leftBounds[1])
	rightLower, _ := strconv.Atoi(rightBounds[0])
	rightUpper, _ := strconv.Atoi(rightBounds[1])

	if DetectEncapsulated(leftBounds, rightBounds) {
		return true
	}

	if rightLower <= leftUpper && leftLower <= rightLower {
		return true
	}

	if leftLower >= rightLower && leftLower <= rightUpper {
		return true
	}

	return false
}
