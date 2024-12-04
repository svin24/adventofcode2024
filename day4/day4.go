package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Day4() {
	matrix, err := readXmasFile("./day4/input.txt")
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

	var count int = 0
	for r := range matrix {
		for c := range matrix[0] {
			for _, dir := range directions {
				var check bool = true
				for i := range word {
					newRow := r + i*dir[0]
					newCol := c + i*dir[1]
					if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) {
						check = false
						continue
					}
					if matrix[newRow][newCol] != string(word[i]) {
						check = false
						continue
					}
				}
				if check {
					count++
				}
			}
		}
	}
	fmt.Println(count)
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
