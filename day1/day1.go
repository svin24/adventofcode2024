package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1() {
	// read file
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// actual program
	var arr1, arr2 []int

	arr1, arr2 = cutData(lines)
	selectionSort(arr1)
	selectionSort(arr2)

	sum := 0
	for i := range arr1 {
		sum = diff(arr1[i], arr2[i]) + sum
	}
	fmt.Println(sum)
	// part 2
	sum = 0
	for i := range arr1 {
		sum = (arr1[i] * countTimes(arr1[i], arr2)) + sum
	}

	fmt.Println(sum)
}

// takes the data and cuts it into two arrays
func cutData(lines []string) ([]int, []int) {
	var arr1, arr2 []int
	for _, line := range lines {
		fields := strings.Fields(line)

		// not gonna check for errors every line, screw that
		var1, _ := strconv.Atoi(fields[0])
		var2, _ := strconv.Atoi(fields[1])

		arr1 = append(arr1, var1)
		arr2 = append(arr2, var2)
	}
	return arr1, arr2
}

// high school memories
func selectionSort(arr []int) []int {
	for i := range arr {
		var minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func diff(var1 int, var2 int) int {
	if var1 < var2 {
		return var2 - var1
	}
	return var1 - var2
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
