package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	startingItems  []int
	operation      []string
	test           int
	trueMonkey     int
	falseMonkey    int
	timesInspected int
}

var monkeys []*Monkey

func main() {
	file, err := os.Open("res/day11.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var monkeyAttributes []string
	for scanner.Scan() {
		monkeyAttributes = append(monkeyAttributes, scanner.Text())
	}

	var monkey *Monkey
	for _, a := range monkeyAttributes {
		if strings.HasPrefix(a, "Monkey") {
			monkey = &Monkey{}
			continue
		}
		if a == "" {
			monkeys = append(monkeys, monkey)
			continue
		}
		if strings.HasPrefix(a, "  Starting items:") {
			for _, i := range strings.Split(a[18:], ", ") {
				item, _ := strconv.Atoi(i)
				monkey.startingItems = append(monkey.startingItems, item)
			}
			continue
		}
		if strings.HasPrefix(a, "  Operation:") {
			monkey.operation = strings.Split(a[13:], " ")[3:]
			continue
		}
		if strings.HasPrefix(a, "  Test:") {
			test := strings.Split(a[8:], " ")
			monkey.test, _ = strconv.Atoi(test[len(test)-1])
			continue
		}
		if strings.HasPrefix(a, "    If true:") {
			monkey.trueMonkey, _ = strconv.Atoi(string(a[len(a)-1]))
			continue
		}
		if strings.HasPrefix(a, "    If false:") {
			monkey.falseMonkey, _ = strconv.Atoi(string(a[len(a)-1]))
		}
	}
	monkeys = append(monkeys, monkey)

	for i := 0; i < 1000; i++ {
		for _, m := range monkeys {
			m.TakeTurn()
		}
		for _, m := range monkeys {
			fmt.Println(m)
		}
	}

	var timesInspected []int
	for _, m := range monkeys {
		timesInspected = append(timesInspected, m.timesInspected)
	}

	sort.Ints(timesInspected)
	fmt.Println(timesInspected[2] * timesInspected[3])
}

func (m *Monkey) TakeTurn() {
	for _, value := range m.startingItems {
		value = m.InspectItem(value)
		if value%m.test == 0 {
			monkeys[m.trueMonkey].startingItems = append(monkeys[m.trueMonkey].startingItems, value)
		} else {
			monkeys[m.falseMonkey].startingItems = append(monkeys[m.falseMonkey].startingItems, value)
		}
	}
	m.startingItems = []int{}
}

func (m *Monkey) InspectItem(v int) int {
	m.timesInspected++
	rvalue := v
	if m.operation[1] != "old" {
		rvalue, _ = strconv.Atoi(m.operation[1])
	}
	switch m.operation[0] {
	case "+":
		return v + rvalue
	case "*":
		return v * rvalue
	}
	return 0
}
