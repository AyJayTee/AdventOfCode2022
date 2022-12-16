package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AyJayTee/AdventOfCode2022/parser"
)

var (
	pairsAsString []string
)

func main() {
	input := parser.ReadByLines("res/day13.txt")

	for i := 0; i < len(input); i++ {
		if i%3 == 0 {
			pair := fmt.Sprintf(`"%d":{"left":%s, "right":%s}`, (i/3)+1, input[i], input[i+1])
			pairsAsString = append(pairsAsString, pair)
		}
	}

	jsonData := `{"pairs":{` + strings.Join(pairsAsString, ",") + "}}"
	var result map[string]any
	json.Unmarshal([]byte(jsonData), &result)
	pairs := result["pairs"].(map[string]any)

	var sumOfIndexes int

	for i := 1; i < len(pairsAsString); i++ {
		left := pairs[fmt.Sprint(i)].(map[string]any)["left"]
		right := pairs[fmt.Sprint(i)].(map[string]any)["right"]

		for j := 0; true; j++ {
			// Check for out of range
			if j > len(left.([]any))-1 && j <= len(right.([]any))-1 {
				// Left ran out first
				sumOfIndexes += i
				break
			}

			if j > len(right.([]any))-1 {
				// Right ran out first
				break
			}

			leftValue := left.([]any)[j]
			rightValue := right.([]any)[j]

			// If both are integers
			if BothIntegers(leftValue, rightValue) {
				if r := CompareIntegers(int(leftValue.(float64)), int(rightValue.(float64))); r == 2 {
					continue
				} else if r == 1 {
					// Correct
					sumOfIndexes += i
					break
				}
				// Incorrect
				break
			}

			// If both are lists
			if BothLists(leftValue, rightValue) {
				if r := CompareLists(leftValue.([]any), rightValue.([]any)); r == 2 {
					continue
				} else if r == 1 {
					// Correct
					sumOfIndexes += i
					break
				}
				// Incorrect
				break
			}

			// Mixed - find which is integer and convert
			if IsInteger(leftValue) {
				leftValue = []any{leftValue}
			} else {
				rightValue = []any{rightValue}
			}

			// Then compare as lists
			if r := CompareLists(leftValue.([]any), rightValue.([]any)); r == 2 {
				continue
			} else if r == 1 {
				// Correct
				sumOfIndexes += i
				break
			}
			// Incorrect
			break
		}
	}

	fmt.Println(sumOfIndexes)
}

func BothIntegers(left any, right any) bool {
	return (IsInteger(left) && IsInteger(right))
}

func BothLists(left any, right any) bool {
	return (fmt.Sprintf("%T", left) == "[]interface {}" && fmt.Sprintf("%T", right) == "[]interface {}")
}

func IsInteger(value any) bool {
	return fmt.Sprintf("%T", value) == "float64"
}

func CompareIntegers(l int, r int) int {
	if l > r {
		return 0
	}
	if r > l {
		return 1
	}
	return 2
}

func CompareLists(l []any, r []any) int {
	for i := range l {
		// Check for out of range
		if i > len(r)-1 {
			// Right list ran out first
			return 0
		}

		// Check for emtpy list

		if BothLists(l[i], r[i]) {
			return CompareLists(l[i].([]any), r[i].([]any))
		}
		if !BothIntegers(l[i], r[i]) {
			// There is a list - find which one
			if IsInteger(l[i]) {
				return CompareLists([]any{l[i].(float64)}, r[i].([]any))
			}
			return CompareLists(l[i].([]any), []any{r[i].(float64)})
		}

		if result := CompareIntegers(int(l[i].(float64)), int(r[i].(float64))); result == 2 {
			continue
		} else if result == 1 {
			return 1
		}
		return 0
	}
	if len(l) < len(r) {
		return 1
	}
	return 2
}
