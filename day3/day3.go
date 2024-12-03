package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	file, err := os.ReadFile("./day3/input.txt")
	if err != nil {
		panic(err)
	}

	//part 1
	regex := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	matchMul := regex.FindAllString(string(file), -1)
	fmt.Println("Part 1:", trimAndCalculate(matchMul))

	//part 2
	regex = regexp.MustCompile(`(do\(\)|don't\(\)|mul\(\d+,\d+\))`)
	matchMul = regex.FindAllString(string(file), -1)

	var result []string
	isEnabled := true
	for _, match := range matchMul {
		if match == "do()" {
			isEnabled = true
		} else if match == "don't()" {
			isEnabled = false
		} else if isEnabled && match[:4] == "mul(" {
			result = append(result, match)
		}
	}

	fmt.Println("Part 2:", trimAndCalculate(result))
}

func trimAndCalculate(data []string) int {
	var sum int = 0
	for i := range data {
		trim := strings.TrimSuffix(strings.TrimPrefix(data[i], "mul("), ")")
		parts := strings.Split(trim, ",")
		num0, _ := strconv.Atoi(parts[0])
		num1, _ := strconv.Atoi(parts[1])
		sum = num0*num1 + sum
	}
	return sum
}
