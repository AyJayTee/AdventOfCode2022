package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("res/day3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rucksacks []string
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}

	var total int
	for i := 0; i < len(rucksacks); i += 3 {
		three := make([]string, 0)
		three = append(three, rucksacks[i])
		three = append(three, rucksacks[i+1])
		three = append(three, rucksacks[i+2])

		common := DetectCommonCharsThree(three)
		for _, c := range common {
			total += GetCharPriority(c)
		}
	}

	fmt.Println(total)
}

func GetCharPriority(c rune) int {
	v := int(c)
	if v > 97 {
		return v - 96
	}
	return v - 38
}

func DetectCommonChars(l string, r string) string {
	var chars string

	for _, i := range l {
		for _, j := range r {
			if i == j {
				if !ExistsInString(i, chars) {
					chars += string(i)
				}
			}
		}
	}

	return chars
}

func DetectCommonCharsThree(three []string) string {
	commonOneTwo := DetectCommonChars(three[0], three[1])
	return DetectCommonChars(commonOneTwo, three[2])
}

func ExistsInString(c rune, arr string) bool {
	for _, r := range arr {
		if c == r {
			return true
		}
	}
	return false
}
