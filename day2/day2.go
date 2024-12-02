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
	//numbers := make([][]int, len(strings.Split(line, " ")))
	//fmt.Println(numbers)
}
