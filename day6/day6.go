package day6

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Character struct {
	charPos  [2]int
	rotation int
	// 0 is up
	// 1 is right
	// 2 is down
	// 3 is left
}

func Day6() {
	guardMap, err := readAdventFile("./day6/smallinput.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	guard := Character{charPos: [2]int{0, 0}, rotation: 0}
	guard.charPos, err = findGuard(guardMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	// part 1
	part1(guardMap, guard)
	//mapPrint(guardMap)
	silver := mapXCount(guardMap)
	fmt.Println("Part 1:", silver)
}

func mapXCount(guardMap [][]string) int {
	xCount := 0
	for x := range guardMap {
		for y := range guardMap[0] {
			if guardMap[x][y] == "X" {
				xCount++
			}
		}
	}
	return xCount
}

func part1(guardMap [][]string, guard Character) {
	for i := 0; i < len(guardMap); i++ {
		for j := 0; j < len(guardMap[0]); j++ {

			if guard.charPos[0] <= 0 || guard.charPos[0] >= len(guardMap)-1 ||
				guard.charPos[1] <= 0 || guard.charPos[1] >= len(guardMap[0])-1 {
				guardMap[guard.charPos[0]][guard.charPos[1]] = "X"
				return
			}
			// wallcheck
			if guard.rotation == 0 && guardMap[guard.charPos[0]-1][guard.charPos[1]] == "#" {
				guard.rotation = 1
			} else if guard.rotation == 1 && guardMap[guard.charPos[0]][guard.charPos[1]+1] == "#" {
				guard.rotation = 2
			} else if guard.rotation == 2 && guardMap[guard.charPos[0]+1][guard.charPos[1]] == "#" {
				guard.rotation = 3
			} else if guard.rotation == 3 && guardMap[guard.charPos[0]][guard.charPos[1]-1] == "#" { // LEFT
				guard.rotation = 0
			}

			guardMap[guard.charPos[0]][guard.charPos[1]] = "X"
			switch guard.rotation {
			case 0: // UP
				guard.charPos[0]--
				guardMap[guard.charPos[0]][guard.charPos[1]] = "^"
			case 1: // RIGHT
				guard.charPos[1]++
				guardMap[guard.charPos[0]][guard.charPos[1]] = ">"
			case 2: // DOWN
				guard.charPos[0]++
				guardMap[guard.charPos[0]][guard.charPos[1]] = "v"
			case 3: // LEFT
				guard.charPos[1]--
				guardMap[guard.charPos[0]][guard.charPos[1]] = "<"
			}
		}
	}

}

func findGuard(guardMap [][]string) ([2]int, error) {
	for x := range guardMap {
		for y := range guardMap[0] {
			if guardMap[x][y] == "^" {
				return [2]int{x, y}, nil
			}
		}
	}
	return [2]int{0, 0}, errors.New("no guard in map")
}

func mapPrint(arr [][]string) {
	for i := range arr {
		for j := range arr[0] {
			fmt.Printf("%s", arr[i][j])
		}
	}
}

func readAdventFile(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
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
