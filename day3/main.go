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

func calMulInstructs(mulInstructs [][]string) int {
	sum := 0

	for _, instruct := range mulInstructs {
		first, _ := strconv.Atoi(instruct[1])
		second, _ := strconv.Atoi(instruct[2])

		sum += first * second
	}

	return sum
}

func findAllMatches(program string, pattern string) [][]string {
	reg := regexp.MustCompile(pattern)

	return reg.FindAllStringSubmatch(program, -1)
}

func findFirstMatchIndex(program string, pattern string) []int {
	reg := regexp.MustCompile(pattern)

	return reg.FindStringIndex(program)

}

func task1(program string) int {
	mulInstructs := findAllMatches(program, `mul\((\d+),(\d+)\)`)

	return calMulInstructs(mulInstructs)
}

func task2(program string) int {
	enable := true

	mulInstructs := [][]string{}

	for {
		if enable {
			disableMatch := findFirstMatchIndex(program, `don't\(\)`)
			subprogram := program

			if len(disableMatch) != 0 {
				subprogram = program[:disableMatch[0]]
				program = program[disableMatch[1]:]

				enable = false
			}

			mulInstructs = append(mulInstructs, findAllMatches(subprogram, `mul\((\d+),(\d+)\)`)...)

			if len(disableMatch) == 0 {
				break
			}
		} else {
			enableMatch := findFirstMatchIndex(program, `do\(\)`)
			program = program[enableMatch[1]:]

			enable = true
		}
	}

	return calMulInstructs(mulInstructs)
}

func main() {
	program, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task 1 result: %d\n", task1(program))
	fmt.Printf("Task 2 result: %d\n", task2(program))
}
