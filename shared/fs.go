package shared

import (
	"fmt"
	"os"
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

func LoadFileIntMatrix(path string) (rows [][]int) {
	lines := LoadFileRows(path)

	rows = [][]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, ":", "")
		parts := strings.Split(line, " ")

		row := []int{}
		for _, part := range parts {
			row = append(row, ToInt(part))
		}

		rows = append(rows, row)
	}

	return
}

func LoadFileIntRow(path string, separator string) (row []int) {
	line := LoadFileRows(path)[0]
	parts := strings.Split(line, separator)

	for _, part := range parts {
		row = append(row, ToInt(part))
	}

	return
}

func LoadFileStringMatrix(path string) (rows [][]string) {
	lines := LoadFileRows(path)

	rows = [][]string{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		row := strings.Split(line, "")
		rows = append(rows, row)
	}

	return
}

func LoadFileStringMatrixStruct(path string) (matrix Matrix[string]) {
	lines := LoadFileRows(path)

	coordinates := []Point{}
	rows := [][]string{}

	for y, line := range lines {
		if line == "" {
			continue
		}

		row := strings.Split(line, "")
		rows = append(rows, row)
		for x := range row {
			coordinates = append(coordinates, Point{X: x, Y: y})
		}
	}

	return Matrix[string]{Rows: rows, Coords: coordinates}
}

func LoadFileIntMatrixStruct(path string) (matrix Matrix[int]) {
	lines := LoadFileRows(path)

	coordinates := []Point{}
	rows := [][]int{}

	for y, line := range lines {
		if line == "" {
			continue
		}

		line := strings.Split(line, "")

		row := []int{}
		for x, part := range line {
			row = append(row, ToInt(part))
			coordinates = append(coordinates, Point{X: x, Y: y})
		}

		rows = append(rows, row)
	}

	return Matrix[int]{Rows: rows, Coords: coordinates}
}
