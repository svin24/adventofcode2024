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

	// part 1
	var silver int
	var gold int
	var invalidUpdates [][]int

	for _, update := range updates {
		if isUpdateValid(update, rules) {
			middle := findMiddlePage(update)
			silver += middle
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	// Process invalid updates by reordering them
	for _, invalidUpdate := range invalidUpdates {
		reorderedUpdate := reorderUpdate(invalidUpdate, rules)
		middle := findMiddlePage(reorderedUpdate)
		gold += middle
	}

	fmt.Println("Part 1:", silver)
	fmt.Println("Part 2:", gold)
}

func reorderUpdate(update []int, rules []Rules) []int {
	// Build a graph of dependencies
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	// Initialize graph nodes
	for _, page := range update {
		graph[page] = []int{}
		inDegree[page] = 0
	}

	// Populate graph edges based on rules
	for _, rule := range rules {
		if contains(update, rule.Before) && contains(update, rule.After) {
			graph[rule.Before] = append(graph[rule.Before], rule.After)
			inDegree[rule.After]++
		}
	}

	// Perform topological sort
	var sorted []int
	var queue []int

	// Find nodes with no incoming edges
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

// Go apparently can't do this on it's own
func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
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
