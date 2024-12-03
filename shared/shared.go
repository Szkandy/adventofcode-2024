package shared

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadFile(path string) (content string) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	return string(file)
}

func LoadFileRows(path string) (lines []string) {
	lines = strings.Split(LoadFile(path), "\n")
	return
}

func LoadFileIntRows(path string) (rows [][]int) {
	lines := LoadFileRows(path)

	rows = [][]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")

		row := []int{}
		for _, part := range parts {
			row = append(row, ToInt(part))
		}

		rows = append(rows, row)
	}

	return
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func WithoutIndex(s []int, index int) []int {
	var dst []int
	dst = append(dst, s[:index]...)
	return append(dst, s[index+1:]...)
}
