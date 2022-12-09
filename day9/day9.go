package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type RopeSystem struct {
	knots []Knot
}

type Knot struct {
	x int
	y int
}

var tailPositions []Knot

func main() {
	file, err := os.Open("res/day9.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	bridge := RopeSystem{}
	for i := 0; i < 10; i++ {
		bridge.knots = append(bridge.knots, Knot{})
	}
	tailPositions = append(tailPositions, bridge.knots[len(bridge.knots)-1])
	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		iterations, _ := strconv.Atoi(move[1])
		for i := 0; i < iterations; i++ {
			bridge.NextStep(move[0])
		}
	}
	fmt.Println(len(tailPositions))
}

func (s *RopeSystem) NextStep(dir string) {
	s.knots[0].Move(GetDirection(dir))
	for i := 0; i < len(s.knots)-1; i++ {
		xDiff := s.knots[i].x - s.knots[i+1].x
		yDiff := s.knots[i].y - s.knots[i+1].y
		if s.knots[i].x != s.knots[i+1].x && s.knots[i].y != s.knots[i+1].y {
			if math.Abs(float64(xDiff)) > 1 || math.Abs(float64(yDiff)) > 1 {
				s.knots[i+1].Move(xDiff/int(math.Abs(float64(xDiff))), yDiff/int(math.Abs(float64(yDiff))))
				tailPositions = AddIfNotExists(s.knots[len(s.knots)-1], tailPositions)
			}
		} else if math.Abs(float64(xDiff)) > 1 && yDiff == 0 {
			s.knots[i+1].Move(xDiff/int(math.Abs(float64(xDiff))), 0)
			tailPositions = AddIfNotExists(s.knots[len(s.knots)-1], tailPositions)
		} else if math.Abs(float64(yDiff)) > 1 && xDiff == 0 {
			s.knots[i+1].Move(0, yDiff/int(math.Abs(float64(yDiff))))
			tailPositions = AddIfNotExists(s.knots[len(s.knots)-1], tailPositions)
		}
	}
}

func (e *Knot) Move(x int, y int) {
	e.x += x
	e.y += y
}

func GetDirection(dir string) (int, int) {
	switch dir {
	case "U":
		return 0, 1
	case "D":
		return 0, -1
	case "L":
		return -1, 0
	case "R":
		return 1, 0
	}
	return 0, 0
}

func AddIfNotExists(position Knot, positions []Knot) []Knot {
	exists := false
	for _, p := range positions {
		if position == p {
			exists = true
			break
		}
	}
	if !exists {
		positions = append(positions, position)
	}
	return positions
}
