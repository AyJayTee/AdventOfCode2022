package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Supplies struct {
	Stacks []Stack
}

type Stack struct {
	Crates []string
}

func main() {
	file, err := os.Open("res/day5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var positions []string
	var orders []string
	isOrders := false
	for scanner.Scan() {
		if scanner.Text() != "" && !isOrders {
			positions = append(positions, scanner.Text())
		}
		if isOrders {
			orders = append(orders, scanner.Text())
		}
		if scanner.Text() == "" {
			isOrders = true
		}
	}

	columns := len(strings.Split(strings.TrimSpace(positions[len(positions)-1]), "   "))
	supplies := Supplies{}
	for i := 0; i < columns; i++ {
		stack := Stack{}
		supplies.Stacks = append(supplies.Stacks, stack)
	}

	for i := len(positions) - 2; i > -1; i-- {
		row := positions[i]
		for j := 0; j < columns; j++ {
			var letter string
			if j == 0 {
				letter = string(row[1])
			} else {
				letter = string(row[1+(j*4)])
			}

			if letter != " " {
				supplies.Stacks[j].Crates = append(supplies.Stacks[j].Crates, letter)
			}
		}
	}

	for _, o := range orders {
		order := strings.Split(o, " ")
		amount, _ := strconv.Atoi(order[1])
		src, _ := strconv.Atoi(order[3])
		dest, _ := strconv.Atoi(order[5])
		supplies.MoveCrates(amount, src, dest)
	}

	for _, s := range supplies.Stacks {
		fmt.Println(s.Crates[len(s.Crates)-1])
	}
}

func (s *Supplies) MoveCrate(src int, dest int) {
	// Copy last crate to end of dest stack
	srcCrateIndex := len(s.Stacks[src-1].Crates) - 1
	s.Stacks[dest-1].Crates = append(s.Stacks[dest-1].Crates, s.Stacks[src-1].Crates[srcCrateIndex])

	// Remove last crate from src stack
	s.Stacks[src-1].Crates = s.Stacks[src-1].Crates[:srcCrateIndex]
}

func (s *Supplies) MoveCrates(amount int, src int, dest int) {
	// Copy the crates to end of dest stack
	cratesIndex := len(s.Stacks[src-1].Crates) - amount
	crates := s.Stacks[src-1].Crates[cratesIndex:]
	s.Stacks[dest-1].Crates = append(s.Stacks[dest-1].Crates, crates...)

	// Remove the crates from the src stack
	s.Stacks[src-1].Crates = s.Stacks[src-1].Crates[:cratesIndex]
}
