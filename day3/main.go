package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

func readInput() (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("unable to get the current file path")
	}

	currentDir := filepath.Dir(currentFile)

	file, err := os.ReadFile(filepath.Join(currentDir, "input.txt"))
	if err != nil {
		return "", fmt.Errorf("error opening input: %s", err)
	}

	return string(file), nil
}

func sumMulInstructs(program string) int {
	pattern := `mul\((\d+),(\d+)\)`

	reg := regexp.MustCompile(pattern)
	matches := reg.FindAllStringSubmatch(program, -1)

	sum := 0

	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])

		sum += first * second
	}

	return sum
}

func main() {
	program, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task 1 result: %d\n", sumMulInstructs(program))
}
