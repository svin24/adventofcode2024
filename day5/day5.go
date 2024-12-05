package day5

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Rules struct {
	Before, After int
}

func Day5() {
	arr1, arr2, err := readAdventFile("./day5/smallinput.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// text parsing stuff
	rules := getRules(arr1)
	updates := getUpdates(arr2)

	var totalMiddle int
	for _, update := range updates {
		if isUpdateValid(update, rules) {
			middle := findMiddlePage(update)
			fmt.Printf("Valid update: %v, Middle page: %d\n", update, middle)
			totalMiddle += middle
		} else {
			fmt.Printf("Invalid update: %v\n", update)
		}
	}
	fmt.Println(totalMiddle)
}

func findMiddlePage(update []int) int {
	mid := len(update) / 2
	return update[mid]
}

func isUpdateValid(update []int, rules []Rules) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}
	for _, rule := range rules {
		if posBefore, okBefore := position[rule.Before]; okBefore {
			if posAfter, okAfter := position[rule.After]; okAfter {
				if posBefore > posAfter {
					return false
				}
			}
		}
	}
	return true
}

func getRules(arr []string) []Rules {
	var rules []Rules
	for _, line := range arr {
		var before, after int
		fmt.Sscanf(line, "%d|%d", &before, &after)
		rules = append(rules, Rules{Before: before, After: after})
	}
	return rules
}

func getUpdates(arr []string) [][]int {
	var updates [][]int
	for _, line := range arr {
		var update []int
		for _, page := range strings.Split(line, ",") {
			var pageNum int
			fmt.Sscanf(page, "%d", &pageNum)
			update = append(update, pageNum)
		}
		updates = append(updates, update)
	}
	return updates
}

func readAdventFile(input string) ([]string, []string, error) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer file.Close()

	var arr1, arr2 []string
	var foundEmptyLine bool = false

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		if strings.TrimSpace(line) == "" && !foundEmptyLine {
			foundEmptyLine = true
			continue
		}

		if !foundEmptyLine {
			arr1 = append(arr1, line)
		} else {
			arr2 = append(arr2, line)
		}
	}
	return arr1, arr2, err
}
