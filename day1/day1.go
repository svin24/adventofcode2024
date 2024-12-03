package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	// read file
	file, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	//make and fill arrays
	arr1, arr2 := make([]int, len(lines)), make([]int, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		arr1[i], _ = strconv.Atoi(fields[0])
		arr2[i], _ = strconv.Atoi(fields[1])
	}

	//array sorting
	sort.Ints(arr1)
	sort.Ints(arr2)

	sum := 0
	for i := range lines {
		sum += abs(arr1[i] - arr2[i])
	}
	fmt.Println("Part 1:", sum)
	// part 2
	sum = 0
	for i := range lines {
		sum = (arr1[i] * countTimes(arr1[i], arr2)) + sum
	}

	fmt.Println("Part 2:", sum)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countTimes(target int, arr []int) int {
	count := 0
	for i := range arr {
		if target == arr[i] {
			count++
		}
	}
	return count
}
