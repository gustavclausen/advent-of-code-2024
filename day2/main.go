package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func readInput() (levels [][]int, error error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("unable to get the current file path")
	}

	currentDir := filepath.Dir(currentFile)

	file, err := os.Open(filepath.Join(currentDir, "input.txt"))
	if err != nil {
		return nil, fmt.Errorf("error opening input: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	levels = [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		levelsRow := []int{}

		for _, col := range columns {
			level, err := strconv.Atoi(col)
			if err != nil {
				return nil, fmt.Errorf("failed to convert level input '%s' to integer: %s", col, err)
			}

			levelsRow = append(levelsRow, level)
		}

		levels = append(levels, levelsRow)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input file: %s", err)
	}

	return levels, nil
}

func isAscending(level1 int, level2 int) bool {
	return level2 > level1
}

func checkSafeLevels(levels []int) bool {
	asc := isAscending(levels[0], levels[1])

	for i := range levels {
		if i == len(levels)-1 {
			return true
		}

		if asc && !isAscending(levels[i], levels[i+1]) {
			return false
		}

		if asc {
			if levels[i+1]-levels[i] > 3 || levels[i+1]-levels[i] <= 0 {
				return false
			}
		} else {
			if levels[i]-levels[i+1] > 3 || levels[i]-levels[i+1] <= 0 {
				return false
			}
		}
	}

	return false
}

func task1(levelReports [][]int) int {
	safeLevelsCounter := 0

	for _, levels := range levelReports {
		safeLevels := checkSafeLevels(levels)

		if safeLevels {
			safeLevelsCounter++
		}
	}

	return safeLevelsCounter
}

func task2(levelReports [][]int) int {
	safeLevelsCounter := 0

	for _, levels := range levelReports {
		safeLevels := checkSafeLevels(levels)

		if safeLevels {
			safeLevelsCounter++
		} else {
			for i := range levels {
				newLevels := append([]int{}, levels[:i]...)
				newLevels = append(newLevels, levels[i+1:]...)

				safeLevels = checkSafeLevels(newLevels)

				if safeLevels {
					safeLevelsCounter++
					break
				}
			}
		}
	}

	return safeLevelsCounter
}

func main() {
	levelReports, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task 1 answer: %d\n", task1(levelReports))
	fmt.Printf("Task 2 answer: %d\n", task2(levelReports))
}
