package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cycle = 1
var registerX = 1
var instructionIndex = 0
var adding = false

func main() {
	file, err := os.Open("res/day10.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	var screen []string
	var row int
	for {
		if cycle == 41 || cycle == 81 || cycle == 121 || cycle == 161 || cycle == 201 {
			row++
		}

		// Check if sprite is overlapping CRT cursor
		if cycle-(row*40) == registerX || cycle-(row*40) == registerX+1 || cycle-(row*40) == registerX+2 {
			screen = append(screen, "#")
		} else {
			screen = append(screen, ".")
		}

		if instructionIndex >= len(instructions) {
			break
		}
		instruction := strings.Split(instructions[instructionIndex], " ")

		NextCycle(instruction)
	}

	fmt.Println(strings.Join(screen[:40], ""))
	fmt.Println(strings.Join(screen[40:80], ""))
	fmt.Println(strings.Join(screen[80:120], ""))
	fmt.Println(strings.Join(screen[120:160], ""))
	fmt.Println(strings.Join(screen[160:200], ""))
	fmt.Println(strings.Join(screen[200:240], ""))
}

func NextCycle(i []string) {
	cycle++
	if i[0] == "noop" {
		instructionIndex++
		return
	}
	if i[0] == "addx" && !adding {
		adding = true
		return
	}
	if adding {
		amount, _ := strconv.Atoi(i[1])
		registerX += amount
		instructionIndex++
		adding = false
		return
	}
}
