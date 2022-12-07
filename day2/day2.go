package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("res/day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rows []string

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	var total int

	for _, v := range rows {
		row := strings.Split(v, " ")
		rowValue, err := CalculateRowValue(row)
		if err != nil {
			panic(err)
		}
		total += rowValue
	}

	fmt.Println(total)
}

func CalculateRowValue(r []string) (int, error) {
	// Get win state
	outcomeValue, err := GetOutcomeValue(r)
	if err != nil {
		return 0, err
	}

	// Get player value

	playerChoice, err := CalculatePlayerChoice(r)
	if err != nil {
		return 0, err
	}
	playerValue, err := GetPlayerValue(playerChoice)
	if err != nil {
		return 0, err
	}

	return outcomeValue + playerValue, nil
}

func GetOutcomeValue(r []string) (int, error) {
	switch r[1] {
	case "X":
		return 0, nil
	case "Y":
		return 3, nil
	case "Z":
		return 6, nil
	default:
		return 0, errors.New("must be XYZ")
	}
}

func GetPlayerValue(s string) (int, error) {
	switch s {
	case "X":
		return 1, nil
	case "Y":
		return 2, nil
	case "Z":
		return 3, nil
	default:
		return 0, errors.New("must be X, Y or Z")
	}
}

func CalculatePlayerChoice(r []string) (string, error) {
	switch r[1] {
	case "X":
		return GetLoseChoice(r[0])
	case "Y":
		return GetDrawChoice(r[0])
	case "Z":
		return GetWinChoice(r[0])
	default:
		return "", errors.New("must be XYZ")
	}
}

func GetLoseChoice(opponent string) (string, error) {
	switch opponent {
	case "A":
		return "Z", nil
	case "B":
		return "X", nil
	case "C":
		return "Y", nil
	default:
		return "", errors.New("must be ABC")
	}
}

func GetDrawChoice(opponent string) (string, error) {
	switch opponent {
	case "A":
		return "X", nil
	case "B":
		return "Y", nil
	case "C":
		return "Z", nil
	default:
		return "", errors.New("must be ABC")
	}
}

func GetWinChoice(opponent string) (string, error) {
	switch opponent {
	case "A":
		return "Y", nil
	case "B":
		return "Z", nil
	case "C":
		return "X", nil
	default:
		return "", errors.New("must be ABC")
	}
}

/*
func GetOutcomeValue(r []string) (int, error) {
	switch r[1] {
	case "X":
		return GetRockState(r[0])
	case "Y":
		return GetPaperState(r[0])
	case "Z":
		return GetScissorsState(r[0])
	default:
		return 0, errors.New("must be X, Y or Z")
	}
}

func GetRockState(opponent string) (int, error) {
	switch opponent {
	case "A":
		return 3, nil
	case "B":
		return 0, nil
	case "C":
		return 6, nil
	default:
		return 0, errors.New("must be A, B or C")
	}
}

func GetPaperState(opponent string) (int, error) {
	switch opponent {
	case "A":
		return 6, nil
	case "B":
		return 3, nil
	case "C":
		return 0, nil
	default:
		return 0, errors.New("must be A, B or C")
	}
}

func GetScissorsState(opponent string) (int, error) {
	switch opponent {
	case "A":
		return 0, nil
	case "B":
		return 6, nil
	case "C":
		return 3, nil
	default:
		return 0, errors.New("must be A, B or C")
	}
}
*/
