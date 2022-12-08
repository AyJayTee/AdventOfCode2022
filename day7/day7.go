package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	currentDir  []string
	directories []string
	fileSizes   map[string]int
	totalSpace  = 70000000
	spaceNeeded = 30000000
	unusedSpace int
)

func main() {
	file, err := os.Open("res/day7.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	fileSizes = make(map[string]int)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "$" {
			if line[1] == "cd" {
				if line[2] != ".." {
					currentDir = append(currentDir, line[2])
					if !DirectoryExists(strings.Join(currentDir, "_")) {
						directories = append(directories, strings.Join(currentDir, "_"))
					}
				} else {
					currentDir = currentDir[:len(currentDir)-1]
				}
			}
		} else if line[0] != "dir" {
			file := strings.Join(currentDir, "_") + "_" + line[1]
			fileSize, _ := strconv.Atoi(line[0])
			fileSizes[file] = fileSize
		}
	}

	unusedSpace = totalSpace - GetDirectorySize("/")
	spaceToFree := spaceNeeded - unusedSpace

	var directorySizes []int
	for _, d := range directories {
		directoryTotal := GetDirectorySize(d)
		if directoryTotal >= spaceToFree {
			directorySizes = append(directorySizes, directoryTotal)
		}
	}

	smallest := directorySizes[0]
	for _, v := range directorySizes[1:] {
		if v < smallest {
			smallest = v
		}
	}
	fmt.Println(smallest)
}

func GetDirectorySize(dir string) int {
	var total int
	for k, v := range fileSizes {
		if strings.HasPrefix(k, dir) {
			total += v
		}
	}
	return total
}

func DirectoryExists(current string) bool {
	for _, d := range directories {
		if current == d {
			return true
		}
	}
	return false
}
