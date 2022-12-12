package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Coordinate struct {
	x     int
	y     int
	count int
}

var (
	grid                  [][]int
	start                 Coordinate
	end                   Coordinate
	usedSquares           []Coordinate
	newSquares            []Coordinate
	numberOfRows          int
	elevationACoordinates []Coordinate
)

func main() {
	file, err := os.Open("res/day12.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
		numberOfRows++
	}

	for y, row := range rows {
		var convertedRow []int
		for x, character := range row {
			if string(character) == "S" {
				convertedRow = append(convertedRow, 1)
				start = Coordinate{x: x, y: y}
				continue
			}
			if string(character) == "E" {
				convertedRow = append(convertedRow, 26)
				end = Coordinate{x: x, y: y}
				continue
			}
			convertedRow = append(convertedRow, int(character)-96)
		}
		grid = append(grid, convertedRow)
	}

	for y, row := range grid {
		for x, square := range row {
			if square == 1 {
				elevationACoordinates = append(elevationACoordinates, Coordinate{x: x, y: y})
			}
		}
	}

	var steps []int
	for _, startingPoint := range elevationACoordinates {
		usedSquares = []Coordinate{startingPoint}
		reachedEnd := false
		index := 0
		newSquares = []Coordinate{}

		for !reachedEnd {
			for _, s := range usedSquares[index:] {
				newSquares = append(newSquares, FindSquares(s)...)
			}
			if len(newSquares) == 0 {
				reachedEnd = true
				fmt.Printf("Completed for: %o \n", startingPoint)
			}
			for _, s := range newSquares {
				if s.x == end.x && s.y == end.y {
					reachedEnd = true
					steps = append(steps, s.count)
					fmt.Printf("Completed for: %o \n", startingPoint)
				}
			}
			index = len(usedSquares)
			usedSquares = append(usedSquares, newSquares...)
			newSquares = []Coordinate{}
		}
	}

	sort.Ints(steps)
	fmt.Println(steps[0])
}

func FindSquares(p Coordinate) []Coordinate {
	var squares []Coordinate
	if p.y-1 >= 0 {
		if grid[p.y][p.x]+1 >= grid[p.y-1][p.x] {
			square := Coordinate{x: p.x, y: p.y - 1, count: p.count + 1}
			if !SquareUsed(square, usedSquares) && !SquareUsed(square, newSquares) {
				squares = append(squares, square)
			}
		}
	}
	if p.y+1 < numberOfRows {
		if grid[p.y][p.x]+1 >= grid[p.y+1][p.x] {
			square := Coordinate{x: p.x, y: p.y + 1, count: p.count + 1}
			if !SquareUsed(square, usedSquares) && !SquareUsed(square, newSquares) {
				squares = append(squares, square)
			}
		}
	}
	if p.x-1 >= 0 {
		if grid[p.y][p.x]+1 >= grid[p.y][p.x-1] {
			square := Coordinate{x: p.x - 1, y: p.y, count: p.count + 1}
			if !SquareUsed(square, usedSquares) && !SquareUsed(square, newSquares) {
				squares = append(squares, square)
			}
		}
	}
	if p.x+1 < len(grid[0]) {
		if grid[p.y][p.x]+1 >= grid[p.y][p.x+1] {
			square := Coordinate{x: p.x + 1, y: p.y, count: p.count + 1}
			if !SquareUsed(square, usedSquares) && !SquareUsed(square, newSquares) {
				squares = append(squares, square)
			}
		}
	}
	return squares
}

func SquareUsed(c Coordinate, squaresToCheck []Coordinate) bool {
	for _, square := range squaresToCheck {
		if c.x == square.x && c.y == square.y {
			return true
		}
	}
	return false
}
