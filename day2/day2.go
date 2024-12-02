package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	file, err := os.ReadFile("./day2/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	numbers := cutLines(lines)
	// part 1
	totalSum := 0
	for _, line := range numbers {
		if isSafe(line) {
			totalSum++
		}
	}
	fmt.Println(totalSum)
}

func cutLines(lines []string) [][]int {
	var numbers [][]int
	for _, line := range lines {
		// empty line for some reason here
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		stringNums := strings.Fields(line)

		intNums := make([]int, len(stringNums))
		for i, numStr := range stringNums {
			intNums[i], _ = strconv.Atoi(numStr)
		}

		numbers = append(numbers, intNums)
	}
	return numbers
}

func isSafe(line []int) bool {
	// there is the potential for a bug here that I won't check
	var isIncreasing bool = true
	var isDecreasing bool = true
	for i := 0; i < len(line)-1; i++ {
		diff := line[i+1] - line[i]

		// Check if the difference is between 1 and 3
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		// Check if the sequence is not consistently increasing or decreasing
		if diff > 0 {
			isDecreasing = false
		} else if diff < 0 {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
