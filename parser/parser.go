package parser

import (
	"bufio"
	"os"
)

func ReadByLines(filePath string) []string {
	var out []string

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := newScanner("lines", file)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out
}

func newScanner(scannerType string, file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	switch scannerType {
	case "lines":
		scanner.Split(bufio.ScanLines)
	case "runes":
		scanner.Split(bufio.ScanRunes)
	default:
		panic("need a valid scanner type")
	}
	return scanner
}
