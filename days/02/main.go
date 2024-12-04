package main

import (
	"fmt"
	"szkandy/adventofcode-2024/shared"
)

const MAX_DISTANCE = 3
const MIN_DISTANCE = 1

func isValid(current int, next int, isIncreasing bool) bool {
	if (isIncreasing && current > next) ||
		(!isIncreasing && current < next) {
		return false
	}

	d := shared.Abs(next - current)
	if d < MIN_DISTANCE || d > MAX_DISTANCE {
		return false
	}

	return true
}

func isRowValid(row []int, allowError bool) bool {
	isIncreasing := row[0] < row[1]

	for i := 0; i < len(row)-1; i++ {
		if isValid(row[i], row[i+1], isIncreasing) {
			continue
		}

		if !allowError {
			return false
		}

		if i > 0 {
			withoutPrevious := shared.WithoutIndex(row, i-1)

			if isRowValid(withoutPrevious, false) {
				break
			}
		}

		withoutCurrent := shared.WithoutIndex(row, i)
		withoutNext := shared.WithoutIndex(row, i+1)

		if !isRowValid(withoutCurrent, false) && !isRowValid(withoutNext, false) {
			return false
		}
	}

	return true
}

func pt1(rows [][]int) {

	validCount := 0

	for _, row := range rows {
		if isRowValid(row, false) {
			validCount++
		}
	}

	fmt.Println("Part 1 Valid: ", validCount)
}

func pt2(rows [][]int) {

	validCount := 0

	for _, row := range rows {
		if isRowValid(row, true) {
			validCount++
		}
	}

	fmt.Println("Part 2 Valid: ", validCount)
}

func main() {
	reports := shared.LoadFileIntMatrix("./input.txt")

	pt1(reports)
	pt2(reports)
}
