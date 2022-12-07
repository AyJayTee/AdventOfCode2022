package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("res/day6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var sequence []string
	var count int
	for scanner.Scan() {
		if len(sequence) > 14 {
			sequence = sequence[1:]
		}
		if len(sequence) == 14 {
			if IsMarker(sequence) {
				break
			}
		}
		sequence = append(sequence, scanner.Text())
		count++
	}
	fmt.Printf("Marker found: %s \n", strings.Join(sequence, ""))
	fmt.Printf("Character count: %d \n", count)
}

func IsMarker(s []string) bool {
	for i := len(s) - 1; i > 0; i-- {
		seq := s[:i]
		for _, c := range seq {
			if c == s[i] {
				return false
			}
		}
	}
	return s[0] != s[1]
}
