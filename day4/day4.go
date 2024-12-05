package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Day4() {
	matrix, err := readXmasFile("./day4/smallinput.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	word := "XMAS"

	directions := [][2]int{
		{0, 1},   // Right
		{1, 0},   // Down
		{1, 1},   // Diagonal Down-Right
		{-1, 1},  // Diagonal Up-Right
		{0, -1},  // Left
		{-1, 0},  // Up
		{-1, -1}, // Diagonal Up-Left
		{1, -1},  // Diagonal Down-Left
	}
	diagonals := [][2]int{
		{1, 1},   // Diagonal Down-Right
		{-1, 1},  // Diagonal Up-Right
		{-1, -1}, // Diagonal Up-Left
		{1, -1},  // Diagonal Down-Left
	}
	var silver int = 0
	var gold int = 0
	for r := range matrix {
		for c := range matrix[0] {
			// part 1
			for _, dir := range directions {
				if findXmasPart1(matrix, r, c, word, dir) {
					silver++
				}
			}
			// part 1
			if matrix[r][c] == "A" {
				if findXmasPart2(matrix, r, c, diagonals) {
					gold++
				}
			}
		}
	}
	fmt.Println("Part 1:", silver)
	fmt.Println("Part 2:", gold)
}

func findXmasPart1(matrix [][]string, row int, col int, word string, dir [2]int) bool {
	for i := range word {
		newRow := row + i*dir[0]
		newCol := col + i*dir[1]
		//check the limits
		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) {
			return false
		}
		if matrix[newRow][newCol] != string(word[i]) {
			return false
		}
	}
	return true
}

func findXmasPart2(matrix [][]string, row int, col int, diag [][2]int) bool {
	// each MAS should have an A in the middle
	mCount := 0
	for _, d := range diag {
		newRow := row + d[0]
		newCol := col + d[1]
		opRow := row + d[0]*-1
		opCol := col + d[1]*-1
		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) || opRow < 0 || opRow >= len(matrix) || opCol < 0 || opCol >= len(matrix[0]) {
			return false
		}
		if matrix[newRow][newCol] == "M" && matrix[opRow][opCol] == "S" {
			mCount++
			if mCount > 2 {
				return false
			}
		}
	}

	return mCount == 2
}

func readXmasFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var arr [][]string

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lineCut := strings.Split(line, "")
		arr = append(arr, lineCut)
	}
	return arr, err
}
