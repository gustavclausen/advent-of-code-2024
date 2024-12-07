package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func readInput() (letterMatrix [][]string, error error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("unable to get the current file path")
	}

	currentDir := filepath.Dir(currentFile)

	file, err := os.Open(filepath.Join(currentDir, "input.txt"))
	if err != nil {
		return nil, fmt.Errorf("error opening input: %s", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := []rune(line)

		row := make([]string, len(parts))
		for i, part := range parts {
			row[i] = string(part)
		}

		letterMatrix = append(letterMatrix, row)
	}

	return letterMatrix, nil
}

func isOutOfBounds(letterMatrix [][]string, startX, startY int) bool {
	return startX < 0 || startY < 0 || startX > len(letterMatrix)-1 || startY > len(letterMatrix[startX])-1
}

func horizontalSearch(letterMatrix [][]string, startX, startY int, searchString string) bool {
	searchCharsRunes := []rune(searchString)
	for i, r := range searchCharsRunes {
		if isOutOfBounds(letterMatrix, startX, startY+i) || letterMatrix[startX][startY+i] != string(r) {
			return false
		}
	}

	return true
}

func verticalSearch(letterMatrix [][]string, startX, startY int, searchString string) bool {
	searchCharsRunes := []rune(searchString)
	for i, r := range searchCharsRunes {
		if isOutOfBounds(letterMatrix, startX+i, startY) || letterMatrix[startX+i][startY] != string(r) {
			return false
		}
	}

	return true
}

func diagonalSearchRight(letterMatrix [][]string, startX, startY int, searchString string) bool {
	searchCharsRunes := []rune(searchString)
	for i, r := range searchCharsRunes {
		if isOutOfBounds(letterMatrix, startX+i, startY+i) || letterMatrix[startX+i][startY+i] != string(r) {
			return false
		}
	}

	return true
}

func diagonalSearchLeft(letterMatrix [][]string, startX, startY int, searchString string) bool {
	searchCharsRunes := []rune(searchString)
	for i, r := range searchCharsRunes {
    if isOutOfBounds(letterMatrix, startX-i, startY-i) || letterMatrix[startX-i][startY-i] != string(r) {
			return false
		}
	}

	return true
}

func task1(letterMatrix [][]string) int {
	counter := 0
	for i, row := range letterMatrix {
		for j := range row {
			tests := [](func(letterMatrix [][]string, startX int, startY int, searchString string) bool){
				horizontalSearch,
				verticalSearch,
				diagonalSearchLeft,
				diagonalSearchRight,
			}

			for _, test := range tests {
				if test(letterMatrix, i, j, "XMAS") {
					counter++
				}

				if test(letterMatrix, i, j, "SAMX") {
					counter++
				}
			}
		}
	}

	return counter
}

func main() {
	letterMatrix, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Task 1 answer:", task1(letterMatrix))
}
