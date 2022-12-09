package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var grid [][]int

func main() {
	file, err := os.Open("res/day8.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var row []int
		for _, c := range scanner.Text() {
			digit, _ := strconv.Atoi(string(c))
			row = append(row, digit)
		}
		grid = append(grid, row)
	}

	var highestScore int
	for y, yv := range grid {
		for x := range yv {
			if x == 0 || y == 0 || x == len(yv)-1 || y == len(grid)-1 {
				continue
			} else {
				if GetScenicScore(x, y) > highestScore {
					highestScore = GetScenicScore(x, y)
				}
			}
		}
	}
	fmt.Println(highestScore)
}

func CheckVisible(x int, y int, xdir int, ydir int) (bool, int, int) {
	for j := 1; j < len(grid[0])-1; j++ {
		for i := 1; i < len(grid[j])-1; i++ {
			xCoord := x + (i * xdir)
			yCoord := y + (j * ydir)
			if xCoord < 0 || xCoord > len(grid[j])-1 {
				continue
			}
			if yCoord < 0 || yCoord > len(grid[j])-1 {
				continue
			}
			if grid[y][x] <= grid[yCoord][xCoord] {
				return false, i, j
			}
		}
	}
	return true, 0, 0
}

func GetScenicScore(x int, y int) int {
	var upscore int
	var downscore int
	var leftscore int
	var rightscore int
	if visible, _, j := CheckVisible(x, y, 0, -1); visible {
		upscore = y
	} else {
		upscore = j
	}
	if visible, _, j := CheckVisible(x, y, 0, 1); visible {
		downscore = len(grid[0]) - 1 - y
	} else {
		downscore = j
	}
	if visible, i, _ := CheckVisible(x, y, -1, 0); visible {
		leftscore = x
	} else {
		leftscore = i
	}
	if visible, i, _ := CheckVisible(x, y, 1, 0); visible {
		rightscore = len(grid[0]) - 1 - x
	} else {
		rightscore = i
	}
	return upscore * downscore * leftscore * rightscore
}
