package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func addLocToList(rawLocVal string, locs *[]int) error {
	loc, err := strconv.Atoi(rawLocVal)
	if err != nil {
		return fmt.Errorf("failed to convert location input '%s' to integer: %s", rawLocVal, err)
	}

	*locs = append(*locs, loc)

	return nil
}

func readInput() (leftLocs []int, rightLocs []int, error error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, nil, fmt.Errorf("unable to get the current file path")
	}

	currentDir := filepath.Dir(currentFile)

	file, err := os.Open(filepath.Join(currentDir, "input.txt"))
	if err != nil {
		return nil, nil, fmt.Errorf("error opening input: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftLocs = []int{}
	rightLocs = []int{}

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		if len(columns) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		addLocToList(columns[0], &leftLocs)
		addLocToList(columns[1], &rightLocs)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading input file: %s", err)
	}

	return leftLocs, rightLocs, nil
}

func task1(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	var distance = 0
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(right[i] - left[i])))
	}

	return distance
}

func task2(left []int, right []int) int {
	rightOccurrence := map[int]int{}

	for _, location := range right {
		count, exists := rightOccurrence[location]

		if exists {
			rightOccurrence[location] = count + 1
		} else {
			rightOccurrence[location] = 1
		}
	}

	total := 0
	for _, location := range left {
		count, exists := rightOccurrence[location]

		if exists {
			total += location * count
		}
	}

	return total
}

func main() {
	left, right, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	if len(left) != len(right) {
		log.Panicf("length of left and right location list is not the same. Left length: %d. Right length: %d", len(left), len(right))
	}

	fmt.Printf("Task 1 answer: %d\n", task1(left, right))
	fmt.Printf("Task 2 answer: %d\n", task2(left, right))
}
